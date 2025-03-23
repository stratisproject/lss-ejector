package cmd

import (
	"context"
	"fmt"
	"math"
	"os"

	"github.com/SergeevDmitry/eth2-balance-service/shared"
	"github.com/SergeevDmitry/eth2-balance-service/shared/beacon"
	"github.com/SergeevDmitry/eth2-balance-service/shared/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stratisproject/prysm-stratis/io/prompt"

	"lss_ejector/task"
)

const EnvPassword = "KEYSTORE_PASSWORD"

const (
	flagAllAccountsFile      = "all_accounts_file"
	flagKeysDir              = "keys_dir"
	flagKeystorePasswordFile = "keystore_password_file"
	flagExecutionEndpoint    = "execution_endpoint"
	flagConsensusEndpoint    = "consensus_endpoint"
	flagLogLevel             = "log_level"
	flagStakingAddress       = "staking_address"
)

// Keystore json file representation as a Go struct.
type Keystore struct {
	Crypto  map[string]interface{} `json:"crypto"`
	ID      string                 `json:"uuid"`
	Pubkey  string                 `json:"pubkey"`
	Version uint                   `json:"version"`
	Name    string                 `json:"name"`
	Path    string                 `json:"path"`
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		Args:    cobra.ExactArgs(0),
		Short:   "Start ejector process",
		RunE: func(cmd *cobra.Command, args []string) error {
			keysDir, err := cmd.Flags().GetString(flagKeysDir)
			if err != nil {
				return err
			}
			if len(keysDir) == 0 {
				return fmt.Errorf("%s empty", flagKeysDir)
			}
			executionEndpoint, err := cmd.Flags().GetString(flagExecutionEndpoint)
			if err != nil {
				return err
			}
			if len(executionEndpoint) == 0 {
				return fmt.Errorf("%s empty", flagExecutionEndpoint)
			}
			consensusEndpoint, err := cmd.Flags().GetString(flagConsensusEndpoint)
			if err != nil {
				return err
			}
			if len(consensusEndpoint) == 0 {
				return fmt.Errorf("%s empty", flagConsensusEndpoint)
			}
			stakingAddress, err := cmd.Flags().GetString(flagStakingAddress)
			if err != nil {
				return err
			}
			if !common.IsHexAddress(stakingAddress) {
				return fmt.Errorf("parse staking address failed, address: %s", stakingAddress)
			}
			logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
			if err != nil {
				return err
			}
			logLevel, err := logrus.ParseLevel(logLevelStr)
			if err != nil {
				return err
			}
			logrus.SetLevel(logLevel)

			logrus.Infof("%s: %s", flagKeysDir, keysDir)
			logrus.Infof("%s: %s", flagExecutionEndpoint, executionEndpoint)
			logrus.Infof("%s: %s", flagConsensusEndpoint, consensusEndpoint)
			// interrupt signal
			ctx := ShutdownListener()

			keystorePasswordFile, err := cmd.Flags().GetString(flagKeystorePasswordFile)
			if err != nil {
				return err
			}

			var accountsPassword string
			if keystorePasswordFile != "" {
				pass, err := os.ReadFile(keystorePasswordFile)
				if err != nil {
					return errors.Wrap(err, "read keystorePasswordFile failed")
				}
				accountsPassword = string(pass)
			} else if pswdStr := os.Getenv(EnvPassword); pswdStr != "" {
				accountsPassword = pswdStr
			} else {
				accountsPassword, err = prompt.PasswordPrompt(
					"Enter the password for your imported accounts", prompt.NotEmpty,
				)
				if err != nil {
					return fmt.Errorf("could not read account password: %w", err)
				}
			}

			allAccountsFile, err := cmd.Flags().GetBool(flagAllAccountsFile)
			if err != nil {
				return err
			}

			keys := make(map[string]string)
			validators := make(map[string]*task.Validator)
			notExitValidators := make(map[string]*task.Validator, 0)
			checkPubkeys := make([]types.ValidatorPubkey, 0)

			if allAccountsFile {
				keystore, err := parseAllAccountsKeystoreFile(keysDir)
				if err != nil {
					return errors.Wrap(err, "parseAllAccountsKeystoreFile failed")
				}

				privKeysBytes, pubKeysBytes, _, err := attemptDecryptAllAccountsKeystore(keystore, accountsPassword)
				if err != nil {
					return err
				}

				for idx, privKeyBytes := range privKeysBytes {
					pubKeyBytes := pubKeysBytes[idx]

					pubkey := types.BytesToValidatorPubkey(pubKeyBytes)

					// if key exists prior to being added then output log that duplicate key was found
					if _, ok := keys[pubkey.String()]; ok {
						logrus.Warnf("Duplicate key in import will be ignored: %#x", pubKeyBytes)
						continue
					}
					keys[pubkey.String()] = pubkey.String()
					checkPubkeys = append(checkPubkeys, pubkey)

					validator := task.Validator{
						PrivateKey: privKeyBytes,
						Publickey:  pubKeyBytes,
					}
					validators[pubkey.String()] = &validator
					logrus.Debug("pubkeyHexString: ", pubkey.String())
				}
			} else {
				keystores, err := parseKeysDir(keysDir)
				if err != nil {
					return errors.Wrap(err, "parseKeysDir failed")
				}
				if len(keystores) == 0 {
					return fmt.Errorf("no keystore found in directory %s", keysDir)
				}

				for _, keystore := range keystores {
					privKeyBytes, pubKeyBytes, _, err := attemptDecryptKeystore(keystore, accountsPassword)
					if err != nil {
						return err
					}
					pubkey := types.BytesToValidatorPubkey(pubKeyBytes)

					// if key exists prior to being added then output log that duplicate key was found
					if _, ok := keys[pubkey.String()]; ok {
						logrus.Warnf("Duplicate key in import will be ignored: %#x", pubKeyBytes)
						continue
					}
					keys[pubkey.String()] = pubkey.String()
					checkPubkeys = append(checkPubkeys, pubkey)

					validator := task.Validator{
						PrivateKey: privKeyBytes,
						Publickey:  pubKeyBytes,
					}
					validators[pubkey.String()] = &validator
					logrus.Debug("pubkeyHexString: ", pubkey.String())
				}
			}

			connection, err := shared.NewConnection(executionEndpoint, consensusEndpoint, nil, nil, nil)
			if err != nil {
				return errors.Wrap(err, "Start connection with endpoint failed")
			}

			// check rpc on start
			{
				block, err := connection.Eth1Client().BlockNumber(context.Background())
				if err != nil {
					return fmt.Errorf("check executionEndpoint: %s , request failed: %s", executionEndpoint, err.Error())
				}
				if block == 0 {
					return fmt.Errorf("check executionEndpoint: %s , request failed: latest block is zero", executionEndpoint)
				}

				beaconHead, err := connection.Eth2Client().GetBeaconHead()
				if err != nil {
					return fmt.Errorf("check consensusEndpoint: %s , request failed: %s", consensusEndpoint, err.Error())
				}

				if beaconHead.FinalizedEpoch == 0 {
					return fmt.Errorf("check consensusEndpoint: %s , request failed: finalized epoch is zero", consensusEndpoint)
				}
			}

			beaconHead, err := connection.Eth2Client().GetBeaconHead()
			if err != nil {
				return err
			}
			statuses, err := connection.Eth2Client().GetValidatorStatuses(checkPubkeys, &beacon.ValidatorStatusOptions{Epoch: &beaconHead.FinalizedEpoch})
			if err != nil {
				return err
			}
			for _, pubkey := range checkPubkeys {
				status := statuses[pubkey]
				if !status.Exists {
					logrus.Infof("validator %s not exist, will check later", pubkey.String())
					notExitValidators[pubkey.String()] = validators[pubkey.String()]
					delete(validators, pubkey.String())
					continue
				}
				// will skip if already sign exit
				if status.ExitEpoch != math.MaxUint64 {
					logrus.Infof("validator %s already exit, will skip", pubkey.String())
					delete(validators, pubkey.String())
					continue
				}

				validators[pubkey.String()].ValidatorIndex = status.Index
			}

			logrus.Infof("find %d active validators", len(validators))

			t := task.NewTask(common.HexToAddress(stakingAddress), validators, notExitValidators, connection)
			err = t.Start()
			if err != nil {
				return err
			}
			defer func() {
				logrus.Infof("shutting down task ...")
				t.Stop()
			}()

			<-ctx.Done()

			return nil
		},
	}

	cmd.Flags().Bool(flagAllAccountsFile, false, "Whether keys_dir option is a path to all-accounts.keystore.json file")
	cmd.Flags().String(flagKeysDir, "", "Path to a directory where keystores to be imported are stored or a single prysm all-accounts.keystore.json file path (must provide)")
	cmd.Flags().String(flagExecutionEndpoint, "", "Execution node RPC provider endpoint (must provide)")
	cmd.Flags().String(flagConsensusEndpoint, "", "Consensus node RPC provider endpoint (must provide)")
	cmd.Flags().String(flagStakingAddress, "", "Network staking address")
	cmd.Flags().String(flagKeystorePasswordFile, "", "Path to a plain-text, .txt file containing a password for a keystore files")
	cmd.Flags().String(flagLogLevel, logrus.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")
	return cmd
}
