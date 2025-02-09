package task

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/SergeevDmitry/eth2-balance-service/shared"
	"github.com/SergeevDmitry/eth2-balance-service/shared/beacon"
	sharedTypes "github.com/SergeevDmitry/eth2-balance-service/shared/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/signing"
	types "github.com/stratisproject/prysm-stratis/consensus-types/primitives"
	"github.com/stratisproject/prysm-stratis/crypto/bls"
	"github.com/stratisproject/prysm-stratis/encoding/bytesutil"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"

	staking "lss_ejector/bindings/Staking"
)

var (
	domainVoluntaryExit  = bytesutil.Uint32ToBytes4(0x04000000)
	shardCommitteePeriod = types.Epoch(256) // ShardCommitteePeriod is the minimum amount of epochs a validator must participate before exiting.
)

type Task struct {
	stop              chan struct{}
	validators        map[string]*Validator
	notExitValidators map[string]*Validator
	connection        *shared.Connection
	stakingContract   *staking.Staking
	stakingAddress    common.Address
	eth2Config        *beacon.Eth2Config
	rescanBlocks      map[uint64]struct{}
}

type Validator struct {
	ValidatorIndex uint64
	Publickey      []byte
	PrivateKey     []byte
}

func NewTask(stakingAddress common.Address, validators map[string]*Validator, notExitValidators map[string]*Validator, connection *shared.Connection) *Task {
	s := &Task{
		stop:              make(chan struct{}),
		stakingAddress:    stakingAddress,
		validators:        validators,
		notExitValidators: notExitValidators,
		connection:        connection,
		rescanBlocks:      make(map[uint64]struct{}),
	}
	return s
}

func (task *Task) Start() error {
	stakingContract, err := staking.NewStaking(task.stakingAddress, task.connection.Eth1Client())
	if err != nil {
		return err
	}
	task.stakingContract = stakingContract

	ethConfig, err := task.connection.Eth2Client().GetEth2Config()
	if err != nil {
		return err
	}
	task.eth2Config = &ethConfig

	SafeGoWithRestart(task.monitorHandler)
	SafeGoWithRestart(task.rescanBlocksHandler)
	return nil
}

func (task *Task) Stop() {
	close(task.stop)
}

func (task *Task) monitorHandler() {
	logrus.Info("start monitor")

	stakingStartBlock, err := task.stakingContract.StartBlock(nil)
	if err != nil {
		panic(err)
	}
	startBlock := stakingStartBlock.Uint64()
	endBlock, err := task.connection.Eth1LatestBlock()
	if err != nil {
		panic(err)
	}

	if _, err := task.handleExitLogs(startBlock, endBlock); err != nil {
		panic(err)
	}

	startBlock = endBlock + 1

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		case <-ticker.C:
			endBlock, err = task.connection.Eth1LatestBlock()
			if err != nil {
				logrus.Errorf("error getting latest block: %v", err)
				continue
			}

			if rescanBlock, err := task.handleExitLogs(startBlock, endBlock); err != nil {
				task.rescanBlocks[rescanBlock] = struct{}{}
				logrus.Warnf("handleExitLogs err: %v", err)
				continue
			}

			startBlock = endBlock + 1
		}
	}
}

func (task *Task) rescanBlocksHandler() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-task.stop:
			return
		case <-ticker.C:
			if len(task.rescanBlocks) == 0 {
				continue
			}

			for block := range task.rescanBlocks {
				logrus.Debugf("Rescanning block %d", block)
				delete(task.rescanBlocks, block)

				if _, err := task.handleExitLogs(block, block); err != nil {
					task.rescanBlocks[block] = struct{}{}
					logrus.Warnf("rescanBlocksHandler handleExitLogs err: %v", err)
					continue
				}
			}
		}
	}
}

func (task *Task) handleExitLogs(startBlock uint64, endBlock uint64) (uint64, error) {
	logrus.Debugf("Handling ExitValidators events from: %d to: %d", startBlock, endBlock)

	iter, err := task.stakingContract.FilterExitValidators(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: context.Background(),
	})
	if err != nil {
		return startBlock, fmt.Errorf("get ExitValidators logs err: %v", err)
	}

	for iter.Next() {
		if err := task.checkPubkeys(iter.Event.Pubkeys, iter.Event.Raw.BlockNumber); err != nil {
			return iter.Event.Raw.BlockNumber, fmt.Errorf("monitor check pubkeys block err: %v", err)
		}
	}

	if err := task.checkNotExitPubkey(); err != nil {
		logrus.Warnf("checkNotExitPubkey block err: %v", err)
		return 0, nil
	}

	return 0, nil
}

func (task *Task) checkPubkeys(pubkeys [][]byte, blockNumber uint64) error {
	logrus.Debugf("checkPubkeys %d", len(pubkeys))

	for _, pubKeyBytes := range pubkeys {
		pubkey := sharedTypes.BytesToValidatorPubkey(pubKeyBytes)

		validator, exist := task.validators[pubkey.String()]
		if !exist {
			continue
		}

		logrus.Infof("validator %d exit at block %d", validator.ValidatorIndex, blockNumber)
		// check beacon sync status
		syncStatus, err := task.connection.Eth2Client().GetSyncStatus()
		if err != nil {
			return err
		}
		if syncStatus.Syncing {
			return errors.New("could not perform exit: beacon node is syncing.")
		}
		beaconHead, err := task.connection.Eth2Client().GetBeaconHead()
		if err != nil {
			return err
		}
		// check exited before
		status, err := task.connection.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{Epoch: &beaconHead.Epoch})
		if err != nil {
			return err
		}
		// will skip if already sign exit
		if status.ExitEpoch != math.MaxUint64 {
			logrus.Infof("validator %d will exit at epoch %d", validator.ValidatorIndex, status.ExitEpoch)
			delete(task.validators, pubkey.String())
			continue
		}

		currentEpoch := types.Epoch(beaconHead.Epoch)

		// not active
		if uint64(currentEpoch) < status.ActivationEpoch {
			logrus.Warnf("validator %d is not active and can't exit, will skip, active epoch: %d current epoch: %d", validator.ValidatorIndex, status.ActivationEpoch, currentEpoch)
			// remember block to rescan later
			task.rescanBlocks[blockNumber] = struct{}{}
			continue
		}
		if currentEpoch < types.Epoch(status.ActivationEpoch)+shardCommitteePeriod {
			logrus.Warnf("validator %d is not active long enough and can't exit, will skip, active epoch: %d current epoch: %d", validator.ValidatorIndex, status.ActivationEpoch, currentEpoch)
			// remember block to rescan later
			task.rescanBlocks[blockNumber] = struct{}{}
			continue
		}

		// will sign and broadcast exit msg
		exit := &ethpb.VoluntaryExit{Epoch: currentEpoch, ValidatorIndex: types.ValidatorIndex(validator.ValidatorIndex)}

		// Use Cappella fork as here - https://github.com/stratisproject/prysm-stratis/blob/develop/beacon-chain/core/blocks/exit.go#L122
		fork := &ethpb.Fork{
			PreviousVersion: task.eth2Config.CapellaForkVersion,
			CurrentVersion:  task.eth2Config.CapellaForkVersion,
			Epoch:           types.Epoch(task.eth2Config.CapellaForkEpoch),
		}
		domain, err := signing.Domain(fork, exit.Epoch, domainVoluntaryExit, task.eth2Config.GenesisValidatorsRoot)
		if err != nil {
			return errors.Wrap(err, "Get domainData failed")
		}
		exitRoot, err := signing.ComputeSigningRoot(exit, domain)
		if err != nil {
			return errors.Wrap(err, "ComputeSigningRoot failed")
		}

		secretKey, err := bls.SecretKeyFromBytes(validator.PrivateKey)
		if err != nil {
			return errors.Wrap(err, "failed to initialize keys caches from account keystore")
		}
		sig := secretKey.Sign(exitRoot[:])

		err = task.connection.Eth2Client().ExitValidator(validator.ValidatorIndex, uint64(currentEpoch), sharedTypes.BytesToValidatorSignature(sig.Marshal()))
		if err != nil {
			return err
		}

		logrus.Infof("validator %d broadcast voluntary exit ok", validator.ValidatorIndex)
		delete(task.validators, pubkey.String())
	}
	return nil
}

func (task *Task) checkNotExitPubkey() error {
	checkPubkeys := make([]sharedTypes.ValidatorPubkey, 0)
	for key := range task.notExitValidators {
		pubkey, err := sharedTypes.HexToValidatorPubkey(key)
		if err != nil {
			return err
		}

		checkPubkeys = append(checkPubkeys, pubkey)
	}

	beaconHead, err := task.connection.Eth2Client().GetBeaconHead()
	if err != nil {
		return err
	}
	statuses, err := task.connection.GetValidatorStatuses(checkPubkeys, &beacon.ValidatorStatusOptions{Epoch: &beaconHead.Epoch})
	if err != nil {
		return err
	}
	for _, pubkey := range checkPubkeys {
		status := statuses[pubkey]
		if status.Exists && status.ExitEpoch == math.MaxUint64 {
			task.validators[pubkey.String()] = task.notExitValidators[pubkey.String()]
			task.validators[pubkey.String()].ValidatorIndex = status.Index

			delete(task.notExitValidators, pubkey.String())
		}
	}

	return nil
}
