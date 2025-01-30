// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_straxDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposalAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ActiveValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableDepositZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableRewardZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActivePubkeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalExecFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNumberOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"TooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownClaimType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"ExitValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableDeposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIStaking.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"}],\"name\":\"NodeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeRewardsFileCid\",\"type\":\"string\"}],\"name\":\"SetMerkleRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"UnstakeRemainingBalance\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"getActivePubkeysOfNode\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"getActivePubkeysOfNodeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"nodeList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"getPubkeysOfNode\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestMerkleRootEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeAvailableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalExitDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumIStaking.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"nodeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeyInfoOf\",\"outputs\":[{\"internalType\":\"enumIStaking.PubkeyStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pubkeysOfNode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsFileCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_rewardsFileCid\",\"type\":\"string\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setStakingEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUpdateRewardsEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"straxDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalClaimedDepositOfNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalClaimedRewardOfNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeRemainingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewardsEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// GetActivePubkeysOfNode is a free data retrieval call binding the contract method 0xfea29364.
//
// Solidity: function getActivePubkeysOfNode(address _node) view returns(bytes[] pubkeys)
func (_Staking *StakingCaller) GetActivePubkeysOfNode(opts *bind.CallOpts, _node common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getActivePubkeysOfNode", _node)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetActivePubkeysOfNode is a free data retrieval call binding the contract method 0xfea29364.
//
// Solidity: function getActivePubkeysOfNode(address _node) view returns(bytes[] pubkeys)
func (_Staking *StakingSession) GetActivePubkeysOfNode(_node common.Address) ([][]byte, error) {
	return _Staking.Contract.GetActivePubkeysOfNode(&_Staking.CallOpts, _node)
}

// GetActivePubkeysOfNode is a free data retrieval call binding the contract method 0xfea29364.
//
// Solidity: function getActivePubkeysOfNode(address _node) view returns(bytes[] pubkeys)
func (_Staking *StakingCallerSession) GetActivePubkeysOfNode(_node common.Address) ([][]byte, error) {
	return _Staking.Contract.GetActivePubkeysOfNode(&_Staking.CallOpts, _node)
}

// GetActivePubkeysOfNodeLength is a free data retrieval call binding the contract method 0xae9bc2c0.
//
// Solidity: function getActivePubkeysOfNodeLength(address _node) view returns(uint256)
func (_Staking *StakingCaller) GetActivePubkeysOfNodeLength(opts *bind.CallOpts, _node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getActivePubkeysOfNodeLength", _node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActivePubkeysOfNodeLength is a free data retrieval call binding the contract method 0xae9bc2c0.
//
// Solidity: function getActivePubkeysOfNodeLength(address _node) view returns(uint256)
func (_Staking *StakingSession) GetActivePubkeysOfNodeLength(_node common.Address) (*big.Int, error) {
	return _Staking.Contract.GetActivePubkeysOfNodeLength(&_Staking.CallOpts, _node)
}

// GetActivePubkeysOfNodeLength is a free data retrieval call binding the contract method 0xae9bc2c0.
//
// Solidity: function getActivePubkeysOfNodeLength(address _node) view returns(uint256)
func (_Staking *StakingCallerSession) GetActivePubkeysOfNodeLength(_node common.Address) (*big.Int, error) {
	return _Staking.Contract.GetActivePubkeysOfNodeLength(&_Staking.CallOpts, _node)
}

// GetNodes is a free data retrieval call binding the contract method 0x038d67e8.
//
// Solidity: function getNodes(uint256 _start, uint256 _end) view returns(address[] nodeList)
func (_Staking *StakingCaller) GetNodes(opts *bind.CallOpts, _start *big.Int, _end *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getNodes", _start, _end)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNodes is a free data retrieval call binding the contract method 0x038d67e8.
//
// Solidity: function getNodes(uint256 _start, uint256 _end) view returns(address[] nodeList)
func (_Staking *StakingSession) GetNodes(_start *big.Int, _end *big.Int) ([]common.Address, error) {
	return _Staking.Contract.GetNodes(&_Staking.CallOpts, _start, _end)
}

// GetNodes is a free data retrieval call binding the contract method 0x038d67e8.
//
// Solidity: function getNodes(uint256 _start, uint256 _end) view returns(address[] nodeList)
func (_Staking *StakingCallerSession) GetNodes(_start *big.Int, _end *big.Int) ([]common.Address, error) {
	return _Staking.Contract.GetNodes(&_Staking.CallOpts, _start, _end)
}

// GetNodesLength is a free data retrieval call binding the contract method 0x1821aa62.
//
// Solidity: function getNodesLength() view returns(uint256)
func (_Staking *StakingCaller) GetNodesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getNodesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodesLength is a free data retrieval call binding the contract method 0x1821aa62.
//
// Solidity: function getNodesLength() view returns(uint256)
func (_Staking *StakingSession) GetNodesLength() (*big.Int, error) {
	return _Staking.Contract.GetNodesLength(&_Staking.CallOpts)
}

// GetNodesLength is a free data retrieval call binding the contract method 0x1821aa62.
//
// Solidity: function getNodesLength() view returns(uint256)
func (_Staking *StakingCallerSession) GetNodesLength() (*big.Int, error) {
	return _Staking.Contract.GetNodesLength(&_Staking.CallOpts)
}

// GetPubkeysOfNode is a free data retrieval call binding the contract method 0xac1464a3.
//
// Solidity: function getPubkeysOfNode(address _node) view returns(bytes[])
func (_Staking *StakingCaller) GetPubkeysOfNode(opts *bind.CallOpts, _node common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getPubkeysOfNode", _node)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetPubkeysOfNode is a free data retrieval call binding the contract method 0xac1464a3.
//
// Solidity: function getPubkeysOfNode(address _node) view returns(bytes[])
func (_Staking *StakingSession) GetPubkeysOfNode(_node common.Address) ([][]byte, error) {
	return _Staking.Contract.GetPubkeysOfNode(&_Staking.CallOpts, _node)
}

// GetPubkeysOfNode is a free data retrieval call binding the contract method 0xac1464a3.
//
// Solidity: function getPubkeysOfNode(address _node) view returns(bytes[])
func (_Staking *StakingCallerSession) GetPubkeysOfNode(_node common.Address) ([][]byte, error) {
	return _Staking.Contract.GetPubkeysOfNode(&_Staking.CallOpts, _node)
}

// LatestMerkleRootEpoch is a free data retrieval call binding the contract method 0xb5ca7410.
//
// Solidity: function latestMerkleRootEpoch() view returns(uint256)
func (_Staking *StakingCaller) LatestMerkleRootEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "latestMerkleRootEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestMerkleRootEpoch is a free data retrieval call binding the contract method 0xb5ca7410.
//
// Solidity: function latestMerkleRootEpoch() view returns(uint256)
func (_Staking *StakingSession) LatestMerkleRootEpoch() (*big.Int, error) {
	return _Staking.Contract.LatestMerkleRootEpoch(&_Staking.CallOpts)
}

// LatestMerkleRootEpoch is a free data retrieval call binding the contract method 0xb5ca7410.
//
// Solidity: function latestMerkleRootEpoch() view returns(uint256)
func (_Staking *StakingCallerSession) LatestMerkleRootEpoch() (*big.Int, error) {
	return _Staking.Contract.LatestMerkleRootEpoch(&_Staking.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Staking *StakingCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Staking *StakingSession) MerkleRoot() ([32]byte, error) {
	return _Staking.Contract.MerkleRoot(&_Staking.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Staking *StakingCallerSession) MerkleRoot() ([32]byte, error) {
	return _Staking.Contract.MerkleRoot(&_Staking.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_Staking *StakingCaller) NetworkProposalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "networkProposalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_Staking *StakingSession) NetworkProposalAddress() (common.Address, error) {
	return _Staking.Contract.NetworkProposalAddress(&_Staking.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_Staking *StakingCallerSession) NetworkProposalAddress() (common.Address, error) {
	return _Staking.Contract.NetworkProposalAddress(&_Staking.CallOpts)
}

// NodeAvailableBalance is a free data retrieval call binding the contract method 0xfa9b6817.
//
// Solidity: function nodeAvailableBalance(address ) view returns(uint256)
func (_Staking *StakingCaller) NodeAvailableBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nodeAvailableBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeAvailableBalance is a free data retrieval call binding the contract method 0xfa9b6817.
//
// Solidity: function nodeAvailableBalance(address ) view returns(uint256)
func (_Staking *StakingSession) NodeAvailableBalance(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.NodeAvailableBalance(&_Staking.CallOpts, arg0)
}

// NodeAvailableBalance is a free data retrieval call binding the contract method 0xfa9b6817.
//
// Solidity: function nodeAvailableBalance(address ) view returns(uint256)
func (_Staking *StakingCallerSession) NodeAvailableBalance(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.NodeAvailableBalance(&_Staking.CallOpts, arg0)
}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 status, address owner, uint256 depositBlock)
func (_Staking *StakingCaller) PubkeyInfoOf(opts *bind.CallOpts, arg0 []byte) (struct {
	Status       uint8
	Owner        common.Address
	DepositBlock *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pubkeyInfoOf", arg0)

	outstruct := new(struct {
		Status       uint8
		Owner        common.Address
		DepositBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DepositBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 status, address owner, uint256 depositBlock)
func (_Staking *StakingSession) PubkeyInfoOf(arg0 []byte) (struct {
	Status       uint8
	Owner        common.Address
	DepositBlock *big.Int
}, error) {
	return _Staking.Contract.PubkeyInfoOf(&_Staking.CallOpts, arg0)
}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 status, address owner, uint256 depositBlock)
func (_Staking *StakingCallerSession) PubkeyInfoOf(arg0 []byte) (struct {
	Status       uint8
	Owner        common.Address
	DepositBlock *big.Int
}, error) {
	return _Staking.Contract.PubkeyInfoOf(&_Staking.CallOpts, arg0)
}

// PubkeysOfNode is a free data retrieval call binding the contract method 0xa6b9ba17.
//
// Solidity: function pubkeysOfNode(address , uint256 ) view returns(bytes)
func (_Staking *StakingCaller) PubkeysOfNode(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pubkeysOfNode", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PubkeysOfNode is a free data retrieval call binding the contract method 0xa6b9ba17.
//
// Solidity: function pubkeysOfNode(address , uint256 ) view returns(bytes)
func (_Staking *StakingSession) PubkeysOfNode(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Staking.Contract.PubkeysOfNode(&_Staking.CallOpts, arg0, arg1)
}

// PubkeysOfNode is a free data retrieval call binding the contract method 0xa6b9ba17.
//
// Solidity: function pubkeysOfNode(address , uint256 ) view returns(bytes)
func (_Staking *StakingCallerSession) PubkeysOfNode(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Staking.Contract.PubkeysOfNode(&_Staking.CallOpts, arg0, arg1)
}

// RewardsFileCid is a free data retrieval call binding the contract method 0x326ad50b.
//
// Solidity: function rewardsFileCid() view returns(string)
func (_Staking *StakingCaller) RewardsFileCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardsFileCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RewardsFileCid is a free data retrieval call binding the contract method 0x326ad50b.
//
// Solidity: function rewardsFileCid() view returns(string)
func (_Staking *StakingSession) RewardsFileCid() (string, error) {
	return _Staking.Contract.RewardsFileCid(&_Staking.CallOpts)
}

// RewardsFileCid is a free data retrieval call binding the contract method 0x326ad50b.
//
// Solidity: function rewardsFileCid() view returns(string)
func (_Staking *StakingCallerSession) RewardsFileCid() (string, error) {
	return _Staking.Contract.RewardsFileCid(&_Staking.CallOpts)
}

// StakingEnabled is a free data retrieval call binding the contract method 0x1cfff51b.
//
// Solidity: function stakingEnabled() view returns(bool)
func (_Staking *StakingCaller) StakingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakingEnabled is a free data retrieval call binding the contract method 0x1cfff51b.
//
// Solidity: function stakingEnabled() view returns(bool)
func (_Staking *StakingSession) StakingEnabled() (bool, error) {
	return _Staking.Contract.StakingEnabled(&_Staking.CallOpts)
}

// StakingEnabled is a free data retrieval call binding the contract method 0x1cfff51b.
//
// Solidity: function stakingEnabled() view returns(bool)
func (_Staking *StakingCallerSession) StakingEnabled() (bool, error) {
	return _Staking.Contract.StakingEnabled(&_Staking.CallOpts)
}

// StakingTokenAddress is a free data retrieval call binding the contract method 0x5298b869.
//
// Solidity: function stakingTokenAddress() view returns(address)
func (_Staking *StakingCaller) StakingTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakingTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingTokenAddress is a free data retrieval call binding the contract method 0x5298b869.
//
// Solidity: function stakingTokenAddress() view returns(address)
func (_Staking *StakingSession) StakingTokenAddress() (common.Address, error) {
	return _Staking.Contract.StakingTokenAddress(&_Staking.CallOpts)
}

// StakingTokenAddress is a free data retrieval call binding the contract method 0x5298b869.
//
// Solidity: function stakingTokenAddress() view returns(address)
func (_Staking *StakingCallerSession) StakingTokenAddress() (common.Address, error) {
	return _Staking.Contract.StakingTokenAddress(&_Staking.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Staking *StakingCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Staking *StakingSession) StartBlock() (*big.Int, error) {
	return _Staking.Contract.StartBlock(&_Staking.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Staking *StakingCallerSession) StartBlock() (*big.Int, error) {
	return _Staking.Contract.StartBlock(&_Staking.CallOpts)
}

// StraxDepositAddress is a free data retrieval call binding the contract method 0x575c0e4c.
//
// Solidity: function straxDepositAddress() view returns(address)
func (_Staking *StakingCaller) StraxDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "straxDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StraxDepositAddress is a free data retrieval call binding the contract method 0x575c0e4c.
//
// Solidity: function straxDepositAddress() view returns(address)
func (_Staking *StakingSession) StraxDepositAddress() (common.Address, error) {
	return _Staking.Contract.StraxDepositAddress(&_Staking.CallOpts)
}

// StraxDepositAddress is a free data retrieval call binding the contract method 0x575c0e4c.
//
// Solidity: function straxDepositAddress() view returns(address)
func (_Staking *StakingCallerSession) StraxDepositAddress() (common.Address, error) {
	return _Staking.Contract.StraxDepositAddress(&_Staking.CallOpts)
}

// TotalClaimedDepositOfNode is a free data retrieval call binding the contract method 0x6c570dc1.
//
// Solidity: function totalClaimedDepositOfNode(address ) view returns(uint256)
func (_Staking *StakingCaller) TotalClaimedDepositOfNode(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalClaimedDepositOfNode", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedDepositOfNode is a free data retrieval call binding the contract method 0x6c570dc1.
//
// Solidity: function totalClaimedDepositOfNode(address ) view returns(uint256)
func (_Staking *StakingSession) TotalClaimedDepositOfNode(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.TotalClaimedDepositOfNode(&_Staking.CallOpts, arg0)
}

// TotalClaimedDepositOfNode is a free data retrieval call binding the contract method 0x6c570dc1.
//
// Solidity: function totalClaimedDepositOfNode(address ) view returns(uint256)
func (_Staking *StakingCallerSession) TotalClaimedDepositOfNode(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.TotalClaimedDepositOfNode(&_Staking.CallOpts, arg0)
}

// TotalClaimedRewardOfNode is a free data retrieval call binding the contract method 0xbb2d840c.
//
// Solidity: function totalClaimedRewardOfNode(address ) view returns(uint256)
func (_Staking *StakingCaller) TotalClaimedRewardOfNode(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalClaimedRewardOfNode", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedRewardOfNode is a free data retrieval call binding the contract method 0xbb2d840c.
//
// Solidity: function totalClaimedRewardOfNode(address ) view returns(uint256)
func (_Staking *StakingSession) TotalClaimedRewardOfNode(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.TotalClaimedRewardOfNode(&_Staking.CallOpts, arg0)
}

// TotalClaimedRewardOfNode is a free data retrieval call binding the contract method 0xbb2d840c.
//
// Solidity: function totalClaimedRewardOfNode(address ) view returns(uint256)
func (_Staking *StakingCallerSession) TotalClaimedRewardOfNode(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.TotalClaimedRewardOfNode(&_Staking.CallOpts, arg0)
}

// UpdateRewardsEpochs is a free data retrieval call binding the contract method 0xc9f2836d.
//
// Solidity: function updateRewardsEpochs() view returns(uint256)
func (_Staking *StakingCaller) UpdateRewardsEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "updateRewardsEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateRewardsEpochs is a free data retrieval call binding the contract method 0xc9f2836d.
//
// Solidity: function updateRewardsEpochs() view returns(uint256)
func (_Staking *StakingSession) UpdateRewardsEpochs() (*big.Int, error) {
	return _Staking.Contract.UpdateRewardsEpochs(&_Staking.CallOpts)
}

// UpdateRewardsEpochs is a free data retrieval call binding the contract method 0xc9f2836d.
//
// Solidity: function updateRewardsEpochs() view returns(uint256)
func (_Staking *StakingCallerSession) UpdateRewardsEpochs() (*big.Int, error) {
	return _Staking.Contract.UpdateRewardsEpochs(&_Staking.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_Staking *StakingCaller) WithdrawCredentials(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawCredentials")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_Staking *StakingSession) WithdrawCredentials() ([]byte, error) {
	return _Staking.Contract.WithdrawCredentials(&_Staking.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_Staking *StakingCallerSession) WithdrawCredentials() ([]byte, error) {
	return _Staking.Contract.WithdrawCredentials(&_Staking.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_Staking *StakingTransactor) Deposit(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "deposit", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_Staking *StakingSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_Staking *StakingTransactorSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// ExitValidators is a paid mutator transaction binding the contract method 0xfabed6b6.
//
// Solidity: function exitValidators() returns()
func (_Staking *StakingTransactor) ExitValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "exitValidators")
}

// ExitValidators is a paid mutator transaction binding the contract method 0xfabed6b6.
//
// Solidity: function exitValidators() returns()
func (_Staking *StakingSession) ExitValidators() (*types.Transaction, error) {
	return _Staking.Contract.ExitValidators(&_Staking.TransactOpts)
}

// ExitValidators is a paid mutator transaction binding the contract method 0xfabed6b6.
//
// Solidity: function exitValidators() returns()
func (_Staking *StakingTransactorSession) ExitValidators() (*types.Transaction, error) {
	return _Staking.Contract.ExitValidators(&_Staking.TransactOpts)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xfdf435e9.
//
// Solidity: function nodeClaim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_Staking *StakingTransactor) NodeClaim(opts *bind.TransactOpts, _index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "nodeClaim", _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xfdf435e9.
//
// Solidity: function nodeClaim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_Staking *StakingSession) NodeClaim(_index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _Staking.Contract.NodeClaim(&_Staking.TransactOpts, _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xfdf435e9.
//
// Solidity: function nodeClaim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_Staking *StakingTransactorSession) NodeClaim(_index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _Staking.Contract.NodeClaim(&_Staking.TransactOpts, _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x12b81931.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot, string _rewardsFileCid) returns()
func (_Staking *StakingTransactor) SetMerkleRoot(opts *bind.TransactOpts, _dealedEpoch *big.Int, _merkleRoot [32]byte, _rewardsFileCid string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMerkleRoot", _dealedEpoch, _merkleRoot, _rewardsFileCid)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x12b81931.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot, string _rewardsFileCid) returns()
func (_Staking *StakingSession) SetMerkleRoot(_dealedEpoch *big.Int, _merkleRoot [32]byte, _rewardsFileCid string) (*types.Transaction, error) {
	return _Staking.Contract.SetMerkleRoot(&_Staking.TransactOpts, _dealedEpoch, _merkleRoot, _rewardsFileCid)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x12b81931.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot, string _rewardsFileCid) returns()
func (_Staking *StakingTransactorSession) SetMerkleRoot(_dealedEpoch *big.Int, _merkleRoot [32]byte, _rewardsFileCid string) (*types.Transaction, error) {
	return _Staking.Contract.SetMerkleRoot(&_Staking.TransactOpts, _dealedEpoch, _merkleRoot, _rewardsFileCid)
}

// SetStakingEnabled is a paid mutator transaction binding the contract method 0xbce8567e.
//
// Solidity: function setStakingEnabled(bool _enabled) returns()
func (_Staking *StakingTransactor) SetStakingEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setStakingEnabled", _enabled)
}

// SetStakingEnabled is a paid mutator transaction binding the contract method 0xbce8567e.
//
// Solidity: function setStakingEnabled(bool _enabled) returns()
func (_Staking *StakingSession) SetStakingEnabled(_enabled bool) (*types.Transaction, error) {
	return _Staking.Contract.SetStakingEnabled(&_Staking.TransactOpts, _enabled)
}

// SetStakingEnabled is a paid mutator transaction binding the contract method 0xbce8567e.
//
// Solidity: function setStakingEnabled(bool _enabled) returns()
func (_Staking *StakingTransactorSession) SetStakingEnabled(_enabled bool) (*types.Transaction, error) {
	return _Staking.Contract.SetStakingEnabled(&_Staking.TransactOpts, _enabled)
}

// SetUpdateRewardsEpochs is a paid mutator transaction binding the contract method 0x0af5743b.
//
// Solidity: function setUpdateRewardsEpochs(uint256 _value) returns()
func (_Staking *StakingTransactor) SetUpdateRewardsEpochs(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setUpdateRewardsEpochs", _value)
}

// SetUpdateRewardsEpochs is a paid mutator transaction binding the contract method 0x0af5743b.
//
// Solidity: function setUpdateRewardsEpochs(uint256 _value) returns()
func (_Staking *StakingSession) SetUpdateRewardsEpochs(_value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetUpdateRewardsEpochs(&_Staking.TransactOpts, _value)
}

// SetUpdateRewardsEpochs is a paid mutator transaction binding the contract method 0x0af5743b.
//
// Solidity: function setUpdateRewardsEpochs(uint256 _value) returns()
func (_Staking *StakingTransactorSession) SetUpdateRewardsEpochs(_value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetUpdateRewardsEpochs(&_Staking.TransactOpts, _value)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactorSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// UnstakeRemainingBalance is a paid mutator transaction binding the contract method 0x5f1da688.
//
// Solidity: function unstakeRemainingBalance() returns()
func (_Staking *StakingTransactor) UnstakeRemainingBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstakeRemainingBalance")
}

// UnstakeRemainingBalance is a paid mutator transaction binding the contract method 0x5f1da688.
//
// Solidity: function unstakeRemainingBalance() returns()
func (_Staking *StakingSession) UnstakeRemainingBalance() (*types.Transaction, error) {
	return _Staking.Contract.UnstakeRemainingBalance(&_Staking.TransactOpts)
}

// UnstakeRemainingBalance is a paid mutator transaction binding the contract method 0x5f1da688.
//
// Solidity: function unstakeRemainingBalance() returns()
func (_Staking *StakingTransactorSession) UnstakeRemainingBalance() (*types.Transaction, error) {
	return _Staking.Contract.UnstakeRemainingBalance(&_Staking.TransactOpts)
}

// StakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Staking contract.
type StakingDepositIterator struct {
	Event *StakingDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDeposit represents a Deposit event raised by the Staking contract.
type StakingDeposit struct {
	Node               common.Address
	Pubkey             []byte
	ValidatorSignature []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x8efb20bd0b30f1ff45f78dc73475151d028fb0811bda4df1065c922f1243c1d5.
//
// Solidity: event Deposit(address node, bytes pubkey, bytes validatorSignature)
func (_Staking *StakingFilterer) FilterDeposit(opts *bind.FilterOpts) (*StakingDepositIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StakingDepositIterator{contract: _Staking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x8efb20bd0b30f1ff45f78dc73475151d028fb0811bda4df1065c922f1243c1d5.
//
// Solidity: event Deposit(address node, bytes pubkey, bytes validatorSignature)
func (_Staking *StakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StakingDeposit) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDeposit)
				if err := _Staking.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x8efb20bd0b30f1ff45f78dc73475151d028fb0811bda4df1065c922f1243c1d5.
//
// Solidity: event Deposit(address node, bytes pubkey, bytes validatorSignature)
func (_Staking *StakingFilterer) ParseDeposit(log types.Log) (*StakingDeposit, error) {
	event := new(StakingDeposit)
	if err := _Staking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingExitValidatorsIterator is returned from FilterExitValidators and is used to iterate over the raw logs and unpacked data for ExitValidators events raised by the Staking contract.
type StakingExitValidatorsIterator struct {
	Event *StakingExitValidators // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingExitValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingExitValidators)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingExitValidators)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingExitValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingExitValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingExitValidators represents a ExitValidators event raised by the Staking contract.
type StakingExitValidators struct {
	Node    common.Address
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitValidators is a free log retrieval operation binding the contract event 0x9eda5c3255c7ecdfb783150c53aea53178e0db99ed3031ffaea3bfb870104546.
//
// Solidity: event ExitValidators(address node, bytes[] pubkeys)
func (_Staking *StakingFilterer) FilterExitValidators(opts *bind.FilterOpts) (*StakingExitValidatorsIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ExitValidators")
	if err != nil {
		return nil, err
	}
	return &StakingExitValidatorsIterator{contract: _Staking.contract, event: "ExitValidators", logs: logs, sub: sub}, nil
}

// WatchExitValidators is a free log subscription operation binding the contract event 0x9eda5c3255c7ecdfb783150c53aea53178e0db99ed3031ffaea3bfb870104546.
//
// Solidity: event ExitValidators(address node, bytes[] pubkeys)
func (_Staking *StakingFilterer) WatchExitValidators(opts *bind.WatchOpts, sink chan<- *StakingExitValidators) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ExitValidators")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingExitValidators)
				if err := _Staking.contract.UnpackLog(event, "ExitValidators", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExitValidators is a log parse operation binding the contract event 0x9eda5c3255c7ecdfb783150c53aea53178e0db99ed3031ffaea3bfb870104546.
//
// Solidity: event ExitValidators(address node, bytes[] pubkeys)
func (_Staking *StakingFilterer) ParseExitValidators(log types.Log) (*StakingExitValidators, error) {
	event := new(StakingExitValidators)
	if err := _Staking.contract.UnpackLog(event, "ExitValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNodeClaimedIterator is returned from FilterNodeClaimed and is used to iterate over the raw logs and unpacked data for NodeClaimed events raised by the Staking contract.
type StakingNodeClaimedIterator struct {
	Event *StakingNodeClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingNodeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNodeClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingNodeClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingNodeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNodeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNodeClaimed represents a NodeClaimed event raised by the Staking contract.
type StakingNodeClaimed struct {
	Index            *big.Int
	Account          common.Address
	ClaimableReward  *big.Int
	ClaimableDeposit *big.Int
	ClaimType        uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNodeClaimed is a free log retrieval operation binding the contract event 0x659e842f0209726671f562e8d7ee3a25d2cef92c2b85dcd268af93ef5d13582c.
//
// Solidity: event NodeClaimed(uint256 index, address account, uint256 claimableReward, uint256 claimableDeposit, uint8 claimType)
func (_Staking *StakingFilterer) FilterNodeClaimed(opts *bind.FilterOpts) (*StakingNodeClaimedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "NodeClaimed")
	if err != nil {
		return nil, err
	}
	return &StakingNodeClaimedIterator{contract: _Staking.contract, event: "NodeClaimed", logs: logs, sub: sub}, nil
}

// WatchNodeClaimed is a free log subscription operation binding the contract event 0x659e842f0209726671f562e8d7ee3a25d2cef92c2b85dcd268af93ef5d13582c.
//
// Solidity: event NodeClaimed(uint256 index, address account, uint256 claimableReward, uint256 claimableDeposit, uint8 claimType)
func (_Staking *StakingFilterer) WatchNodeClaimed(opts *bind.WatchOpts, sink chan<- *StakingNodeClaimed) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "NodeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNodeClaimed)
				if err := _Staking.contract.UnpackLog(event, "NodeClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeClaimed is a log parse operation binding the contract event 0x659e842f0209726671f562e8d7ee3a25d2cef92c2b85dcd268af93ef5d13582c.
//
// Solidity: event NodeClaimed(uint256 index, address account, uint256 claimableReward, uint256 claimableDeposit, uint8 claimType)
func (_Staking *StakingFilterer) ParseNodeClaimed(log types.Log) (*StakingNodeClaimed, error) {
	event := new(StakingNodeClaimed)
	if err := _Staking.contract.UnpackLog(event, "NodeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSetMerkleRootIterator is returned from FilterSetMerkleRoot and is used to iterate over the raw logs and unpacked data for SetMerkleRoot events raised by the Staking contract.
type StakingSetMerkleRootIterator struct {
	Event *StakingSetMerkleRoot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSetMerkleRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSetMerkleRoot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSetMerkleRoot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSetMerkleRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSetMerkleRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSetMerkleRoot represents a SetMerkleRoot event raised by the Staking contract.
type StakingSetMerkleRoot struct {
	DealedEpoch        *big.Int
	MerkleRoot         [32]byte
	NodeRewardsFileCid string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetMerkleRoot is a free log retrieval operation binding the contract event 0xec43b2424d0504da473794ad49016df3e06fb0d772bb403d724c9e5d53d8afb8.
//
// Solidity: event SetMerkleRoot(uint256 indexed dealedEpoch, bytes32 merkleRoot, string nodeRewardsFileCid)
func (_Staking *StakingFilterer) FilterSetMerkleRoot(opts *bind.FilterOpts, dealedEpoch []*big.Int) (*StakingSetMerkleRootIterator, error) {

	var dealedEpochRule []interface{}
	for _, dealedEpochItem := range dealedEpoch {
		dealedEpochRule = append(dealedEpochRule, dealedEpochItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SetMerkleRoot", dealedEpochRule)
	if err != nil {
		return nil, err
	}
	return &StakingSetMerkleRootIterator{contract: _Staking.contract, event: "SetMerkleRoot", logs: logs, sub: sub}, nil
}

// WatchSetMerkleRoot is a free log subscription operation binding the contract event 0xec43b2424d0504da473794ad49016df3e06fb0d772bb403d724c9e5d53d8afb8.
//
// Solidity: event SetMerkleRoot(uint256 indexed dealedEpoch, bytes32 merkleRoot, string nodeRewardsFileCid)
func (_Staking *StakingFilterer) WatchSetMerkleRoot(opts *bind.WatchOpts, sink chan<- *StakingSetMerkleRoot, dealedEpoch []*big.Int) (event.Subscription, error) {

	var dealedEpochRule []interface{}
	for _, dealedEpochItem := range dealedEpoch {
		dealedEpochRule = append(dealedEpochRule, dealedEpochItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SetMerkleRoot", dealedEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSetMerkleRoot)
				if err := _Staking.contract.UnpackLog(event, "SetMerkleRoot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMerkleRoot is a log parse operation binding the contract event 0xec43b2424d0504da473794ad49016df3e06fb0d772bb403d724c9e5d53d8afb8.
//
// Solidity: event SetMerkleRoot(uint256 indexed dealedEpoch, bytes32 merkleRoot, string nodeRewardsFileCid)
func (_Staking *StakingFilterer) ParseSetMerkleRoot(log types.Log) (*StakingSetMerkleRoot, error) {
	event := new(StakingSetMerkleRoot)
	if err := _Staking.contract.UnpackLog(event, "SetMerkleRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Staking contract.
type StakingStakeIterator struct {
	Event *StakingStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStake represents a Stake event raised by the Staking contract.
type StakingStake struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x8c000936ed6cc8b0c4218956860faac1c834b4ec96b46fa995aa58fc097fea16.
//
// Solidity: event Stake(address node)
func (_Staking *StakingFilterer) FilterStake(opts *bind.FilterOpts) (*StakingStakeIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return &StakingStakeIterator{contract: _Staking.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x8c000936ed6cc8b0c4218956860faac1c834b4ec96b46fa995aa58fc097fea16.
//
// Solidity: event Stake(address node)
func (_Staking *StakingFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakingStake) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStake)
				if err := _Staking.contract.UnpackLog(event, "Stake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStake is a log parse operation binding the contract event 0x8c000936ed6cc8b0c4218956860faac1c834b4ec96b46fa995aa58fc097fea16.
//
// Solidity: event Stake(address node)
func (_Staking *StakingFilterer) ParseStake(log types.Log) (*StakingStake, error) {
	event := new(StakingStake)
	if err := _Staking.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakeRemainingBalanceIterator is returned from FilterUnstakeRemainingBalance and is used to iterate over the raw logs and unpacked data for UnstakeRemainingBalance events raised by the Staking contract.
type StakingUnstakeRemainingBalanceIterator struct {
	Event *StakingUnstakeRemainingBalance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingUnstakeRemainingBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstakeRemainingBalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingUnstakeRemainingBalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingUnstakeRemainingBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakeRemainingBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstakeRemainingBalance represents a UnstakeRemainingBalance event raised by the Staking contract.
type StakingUnstakeRemainingBalance struct {
	Node    common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRemainingBalance is a free log retrieval operation binding the contract event 0xccc4c892a16f289cdc7be946444f106d2413972ea065ccddc7dd5b33a0dbe4d6.
//
// Solidity: event UnstakeRemainingBalance(address indexed node, uint256 balance)
func (_Staking *StakingFilterer) FilterUnstakeRemainingBalance(opts *bind.FilterOpts, node []common.Address) (*StakingUnstakeRemainingBalanceIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UnstakeRemainingBalance", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakeRemainingBalanceIterator{contract: _Staking.contract, event: "UnstakeRemainingBalance", logs: logs, sub: sub}, nil
}

// WatchUnstakeRemainingBalance is a free log subscription operation binding the contract event 0xccc4c892a16f289cdc7be946444f106d2413972ea065ccddc7dd5b33a0dbe4d6.
//
// Solidity: event UnstakeRemainingBalance(address indexed node, uint256 balance)
func (_Staking *StakingFilterer) WatchUnstakeRemainingBalance(opts *bind.WatchOpts, sink chan<- *StakingUnstakeRemainingBalance, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UnstakeRemainingBalance", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstakeRemainingBalance)
				if err := _Staking.contract.UnpackLog(event, "UnstakeRemainingBalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakeRemainingBalance is a log parse operation binding the contract event 0xccc4c892a16f289cdc7be946444f106d2413972ea065ccddc7dd5b33a0dbe4d6.
//
// Solidity: event UnstakeRemainingBalance(address indexed node, uint256 balance)
func (_Staking *StakingFilterer) ParseUnstakeRemainingBalance(log types.Log) (*StakingUnstakeRemainingBalance, error) {
	event := new(StakingUnstakeRemainingBalance)
	if err := _Staking.contract.UnpackLog(event, "UnstakeRemainingBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
