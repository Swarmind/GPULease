// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gpulease

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

// GPULeaseFrozenFundsInfo is an auto generated low-level Go binding around an user-defined struct.
type GPULeaseFrozenFundsInfo struct {
	LeaseId *big.Int
	Amount  *big.Int
}

// GpuleaseMetaData contains all meta data concerning the Gpulease contract.
var GpuleaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"credit_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"name\":\"LeaseCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"name\":\"LeasePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"name\":\"LeaseResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LeaseStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leaseId\",\"type\":\"uint256\"}],\"name\":\"completeLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"frozenFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFrozenFunds\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structGPULease.FrozenFundsInfo[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storagePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"computePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pausedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pausedDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leaseId\",\"type\":\"uint256\"}],\"name\":\"pauseLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leaseId\",\"type\":\"uint256\"}],\"name\":\"resumeLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_storagePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_computePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"startLease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userActiveLeases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GpuleaseABI is the input ABI used to generate the binding from.
// Deprecated: Use GpuleaseMetaData.ABI instead.
var GpuleaseABI = GpuleaseMetaData.ABI

// Gpulease is an auto generated Go binding around an Ethereum contract.
type Gpulease struct {
	GpuleaseCaller     // Read-only binding to the contract
	GpuleaseTransactor // Write-only binding to the contract
	GpuleaseFilterer   // Log filterer for contract events
}

// GpuleaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type GpuleaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GpuleaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GpuleaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GpuleaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GpuleaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GpuleaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GpuleaseSession struct {
	Contract     *Gpulease         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GpuleaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GpuleaseCallerSession struct {
	Contract *GpuleaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GpuleaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GpuleaseTransactorSession struct {
	Contract     *GpuleaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GpuleaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type GpuleaseRaw struct {
	Contract *Gpulease // Generic contract binding to access the raw methods on
}

// GpuleaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GpuleaseCallerRaw struct {
	Contract *GpuleaseCaller // Generic read-only contract binding to access the raw methods on
}

// GpuleaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GpuleaseTransactorRaw struct {
	Contract *GpuleaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGpulease creates a new instance of Gpulease, bound to a specific deployed contract.
func NewGpulease(address common.Address, backend bind.ContractBackend) (*Gpulease, error) {
	contract, err := bindGpulease(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gpulease{GpuleaseCaller: GpuleaseCaller{contract: contract}, GpuleaseTransactor: GpuleaseTransactor{contract: contract}, GpuleaseFilterer: GpuleaseFilterer{contract: contract}}, nil
}

// NewGpuleaseCaller creates a new read-only instance of Gpulease, bound to a specific deployed contract.
func NewGpuleaseCaller(address common.Address, caller bind.ContractCaller) (*GpuleaseCaller, error) {
	contract, err := bindGpulease(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GpuleaseCaller{contract: contract}, nil
}

// NewGpuleaseTransactor creates a new write-only instance of Gpulease, bound to a specific deployed contract.
func NewGpuleaseTransactor(address common.Address, transactor bind.ContractTransactor) (*GpuleaseTransactor, error) {
	contract, err := bindGpulease(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GpuleaseTransactor{contract: contract}, nil
}

// NewGpuleaseFilterer creates a new log filterer instance of Gpulease, bound to a specific deployed contract.
func NewGpuleaseFilterer(address common.Address, filterer bind.ContractFilterer) (*GpuleaseFilterer, error) {
	contract, err := bindGpulease(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GpuleaseFilterer{contract: contract}, nil
}

// bindGpulease binds a generic wrapper to an already deployed contract.
func bindGpulease(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GpuleaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gpulease *GpuleaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gpulease.Contract.GpuleaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gpulease *GpuleaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gpulease.Contract.GpuleaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gpulease *GpuleaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gpulease.Contract.GpuleaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gpulease *GpuleaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gpulease.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gpulease *GpuleaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gpulease.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gpulease *GpuleaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gpulease.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Gpulease *GpuleaseCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Gpulease *GpuleaseSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Gpulease.Contract.Balances(&_Gpulease.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Gpulease *GpuleaseCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Gpulease.Contract.Balances(&_Gpulease.CallOpts, arg0)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_Gpulease *GpuleaseCaller) Credit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "credit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_Gpulease *GpuleaseSession) Credit() (common.Address, error) {
	return _Gpulease.Contract.Credit(&_Gpulease.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_Gpulease *GpuleaseCallerSession) Credit() (common.Address, error) {
	return _Gpulease.Contract.Credit(&_Gpulease.CallOpts)
}

// FrozenFunds is a free data retrieval call binding the contract method 0xa641992a.
//
// Solidity: function frozenFunds(uint256 ) view returns(uint256)
func (_Gpulease *GpuleaseCaller) FrozenFunds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "frozenFunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenFunds is a free data retrieval call binding the contract method 0xa641992a.
//
// Solidity: function frozenFunds(uint256 ) view returns(uint256)
func (_Gpulease *GpuleaseSession) FrozenFunds(arg0 *big.Int) (*big.Int, error) {
	return _Gpulease.Contract.FrozenFunds(&_Gpulease.CallOpts, arg0)
}

// FrozenFunds is a free data retrieval call binding the contract method 0xa641992a.
//
// Solidity: function frozenFunds(uint256 ) view returns(uint256)
func (_Gpulease *GpuleaseCallerSession) FrozenFunds(arg0 *big.Int) (*big.Int, error) {
	return _Gpulease.Contract.FrozenFunds(&_Gpulease.CallOpts, arg0)
}

// GetUserFrozenFunds is a free data retrieval call binding the contract method 0x97c4d284.
//
// Solidity: function getUserFrozenFunds(address user) view returns((uint256,uint256)[] result)
func (_Gpulease *GpuleaseCaller) GetUserFrozenFunds(opts *bind.CallOpts, user common.Address) ([]GPULeaseFrozenFundsInfo, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "getUserFrozenFunds", user)

	if err != nil {
		return *new([]GPULeaseFrozenFundsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]GPULeaseFrozenFundsInfo)).(*[]GPULeaseFrozenFundsInfo)

	return out0, err

}

// GetUserFrozenFunds is a free data retrieval call binding the contract method 0x97c4d284.
//
// Solidity: function getUserFrozenFunds(address user) view returns((uint256,uint256)[] result)
func (_Gpulease *GpuleaseSession) GetUserFrozenFunds(user common.Address) ([]GPULeaseFrozenFundsInfo, error) {
	return _Gpulease.Contract.GetUserFrozenFunds(&_Gpulease.CallOpts, user)
}

// GetUserFrozenFunds is a free data retrieval call binding the contract method 0x97c4d284.
//
// Solidity: function getUserFrozenFunds(address user) view returns((uint256,uint256)[] result)
func (_Gpulease *GpuleaseCallerSession) GetUserFrozenFunds(user common.Address) ([]GPULeaseFrozenFundsInfo, error) {
	return _Gpulease.Contract.GetUserFrozenFunds(&_Gpulease.CallOpts, user)
}

// LeaseCount is a free data retrieval call binding the contract method 0xb4c0498b.
//
// Solidity: function leaseCount() view returns(uint256)
func (_Gpulease *GpuleaseCaller) LeaseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "leaseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeaseCount is a free data retrieval call binding the contract method 0xb4c0498b.
//
// Solidity: function leaseCount() view returns(uint256)
func (_Gpulease *GpuleaseSession) LeaseCount() (*big.Int, error) {
	return _Gpulease.Contract.LeaseCount(&_Gpulease.CallOpts)
}

// LeaseCount is a free data retrieval call binding the contract method 0xb4c0498b.
//
// Solidity: function leaseCount() view returns(uint256)
func (_Gpulease *GpuleaseCallerSession) LeaseCount() (*big.Int, error) {
	return _Gpulease.Contract.LeaseCount(&_Gpulease.CallOpts)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) view returns(address user, address provider, uint256 startTime, uint256 storagePricePerSecond, uint256 computePricePerSecond, bool active, bool completed, bool paused, uint256 pausedAt, uint256 pausedDuration)
func (_Gpulease *GpuleaseCaller) Leases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User                  common.Address
	Provider              common.Address
	StartTime             *big.Int
	StoragePricePerSecond *big.Int
	ComputePricePerSecond *big.Int
	Active                bool
	Completed             bool
	Paused                bool
	PausedAt              *big.Int
	PausedDuration        *big.Int
}, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "leases", arg0)

	outstruct := new(struct {
		User                  common.Address
		Provider              common.Address
		StartTime             *big.Int
		StoragePricePerSecond *big.Int
		ComputePricePerSecond *big.Int
		Active                bool
		Completed             bool
		Paused                bool
		PausedAt              *big.Int
		PausedDuration        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StoragePricePerSecond = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ComputePricePerSecond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Completed = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.PausedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.PausedDuration = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) view returns(address user, address provider, uint256 startTime, uint256 storagePricePerSecond, uint256 computePricePerSecond, bool active, bool completed, bool paused, uint256 pausedAt, uint256 pausedDuration)
func (_Gpulease *GpuleaseSession) Leases(arg0 *big.Int) (struct {
	User                  common.Address
	Provider              common.Address
	StartTime             *big.Int
	StoragePricePerSecond *big.Int
	ComputePricePerSecond *big.Int
	Active                bool
	Completed             bool
	Paused                bool
	PausedAt              *big.Int
	PausedDuration        *big.Int
}, error) {
	return _Gpulease.Contract.Leases(&_Gpulease.CallOpts, arg0)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) view returns(address user, address provider, uint256 startTime, uint256 storagePricePerSecond, uint256 computePricePerSecond, bool active, bool completed, bool paused, uint256 pausedAt, uint256 pausedDuration)
func (_Gpulease *GpuleaseCallerSession) Leases(arg0 *big.Int) (struct {
	User                  common.Address
	Provider              common.Address
	StartTime             *big.Int
	StoragePricePerSecond *big.Int
	ComputePricePerSecond *big.Int
	Active                bool
	Completed             bool
	Paused                bool
	PausedAt              *big.Int
	PausedDuration        *big.Int
}, error) {
	return _Gpulease.Contract.Leases(&_Gpulease.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gpulease *GpuleaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gpulease *GpuleaseSession) Owner() (common.Address, error) {
	return _Gpulease.Contract.Owner(&_Gpulease.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gpulease *GpuleaseCallerSession) Owner() (common.Address, error) {
	return _Gpulease.Contract.Owner(&_Gpulease.CallOpts)
}

// PlatformFeePercentage is a free data retrieval call binding the contract method 0xcdd78cfc.
//
// Solidity: function platformFeePercentage() view returns(uint256)
func (_Gpulease *GpuleaseCaller) PlatformFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "platformFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeePercentage is a free data retrieval call binding the contract method 0xcdd78cfc.
//
// Solidity: function platformFeePercentage() view returns(uint256)
func (_Gpulease *GpuleaseSession) PlatformFeePercentage() (*big.Int, error) {
	return _Gpulease.Contract.PlatformFeePercentage(&_Gpulease.CallOpts)
}

// PlatformFeePercentage is a free data retrieval call binding the contract method 0xcdd78cfc.
//
// Solidity: function platformFeePercentage() view returns(uint256)
func (_Gpulease *GpuleaseCallerSession) PlatformFeePercentage() (*big.Int, error) {
	return _Gpulease.Contract.PlatformFeePercentage(&_Gpulease.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Gpulease *GpuleaseCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Gpulease *GpuleaseSession) Treasury() (common.Address, error) {
	return _Gpulease.Contract.Treasury(&_Gpulease.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Gpulease *GpuleaseCallerSession) Treasury() (common.Address, error) {
	return _Gpulease.Contract.Treasury(&_Gpulease.CallOpts)
}

// UserActiveLeases is a free data retrieval call binding the contract method 0xe869a060.
//
// Solidity: function userActiveLeases(address , uint256 ) view returns(uint256)
func (_Gpulease *GpuleaseCaller) UserActiveLeases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "userActiveLeases", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserActiveLeases is a free data retrieval call binding the contract method 0xe869a060.
//
// Solidity: function userActiveLeases(address , uint256 ) view returns(uint256)
func (_Gpulease *GpuleaseSession) UserActiveLeases(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Gpulease.Contract.UserActiveLeases(&_Gpulease.CallOpts, arg0, arg1)
}

// UserActiveLeases is a free data retrieval call binding the contract method 0xe869a060.
//
// Solidity: function userActiveLeases(address , uint256 ) view returns(uint256)
func (_Gpulease *GpuleaseCallerSession) UserActiveLeases(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Gpulease.Contract.UserActiveLeases(&_Gpulease.CallOpts, arg0, arg1)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_Gpulease *GpuleaseCaller) UserBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gpulease.contract.Call(opts, &out, "userBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_Gpulease *GpuleaseSession) UserBalance(user common.Address) (*big.Int, error) {
	return _Gpulease.Contract.UserBalance(&_Gpulease.CallOpts, user)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_Gpulease *GpuleaseCallerSession) UserBalance(user common.Address) (*big.Int, error) {
	return _Gpulease.Contract.UserBalance(&_Gpulease.CallOpts, user)
}

// CompleteLease is a paid mutator transaction binding the contract method 0x95e6e242.
//
// Solidity: function completeLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseTransactor) CompleteLease(opts *bind.TransactOpts, _leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "completeLease", _leaseId)
}

// CompleteLease is a paid mutator transaction binding the contract method 0x95e6e242.
//
// Solidity: function completeLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseSession) CompleteLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.CompleteLease(&_Gpulease.TransactOpts, _leaseId)
}

// CompleteLease is a paid mutator transaction binding the contract method 0x95e6e242.
//
// Solidity: function completeLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseTransactorSession) CompleteLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.CompleteLease(&_Gpulease.TransactOpts, _leaseId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Gpulease *GpuleaseTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Gpulease *GpuleaseSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.Deposit(&_Gpulease.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Gpulease *GpuleaseTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.Deposit(&_Gpulease.TransactOpts, amount)
}

// PauseLease is a paid mutator transaction binding the contract method 0xb8eeaa78.
//
// Solidity: function pauseLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseTransactor) PauseLease(opts *bind.TransactOpts, _leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "pauseLease", _leaseId)
}

// PauseLease is a paid mutator transaction binding the contract method 0xb8eeaa78.
//
// Solidity: function pauseLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseSession) PauseLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.PauseLease(&_Gpulease.TransactOpts, _leaseId)
}

// PauseLease is a paid mutator transaction binding the contract method 0xb8eeaa78.
//
// Solidity: function pauseLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseTransactorSession) PauseLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.PauseLease(&_Gpulease.TransactOpts, _leaseId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gpulease *GpuleaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gpulease *GpuleaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gpulease.Contract.RenounceOwnership(&_Gpulease.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gpulease *GpuleaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gpulease.Contract.RenounceOwnership(&_Gpulease.TransactOpts)
}

// ResumeLease is a paid mutator transaction binding the contract method 0x18ba11c9.
//
// Solidity: function resumeLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseTransactor) ResumeLease(opts *bind.TransactOpts, _leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "resumeLease", _leaseId)
}

// ResumeLease is a paid mutator transaction binding the contract method 0x18ba11c9.
//
// Solidity: function resumeLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseSession) ResumeLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.ResumeLease(&_Gpulease.TransactOpts, _leaseId)
}

// ResumeLease is a paid mutator transaction binding the contract method 0x18ba11c9.
//
// Solidity: function resumeLease(uint256 _leaseId) returns()
func (_Gpulease *GpuleaseTransactorSession) ResumeLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.ResumeLease(&_Gpulease.TransactOpts, _leaseId)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feePercentage) returns()
func (_Gpulease *GpuleaseTransactor) SetPlatformFee(opts *bind.TransactOpts, _feePercentage *big.Int) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "setPlatformFee", _feePercentage)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feePercentage) returns()
func (_Gpulease *GpuleaseSession) SetPlatformFee(_feePercentage *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.SetPlatformFee(&_Gpulease.TransactOpts, _feePercentage)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feePercentage) returns()
func (_Gpulease *GpuleaseTransactorSession) SetPlatformFee(_feePercentage *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.SetPlatformFee(&_Gpulease.TransactOpts, _feePercentage)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_Gpulease *GpuleaseTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_Gpulease *GpuleaseSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _Gpulease.Contract.SetTreasury(&_Gpulease.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_Gpulease *GpuleaseTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _Gpulease.Contract.SetTreasury(&_Gpulease.TransactOpts, newTreasury)
}

// StartLease is a paid mutator transaction binding the contract method 0xf648da80.
//
// Solidity: function startLease(uint256 _duration, uint256 _storagePricePerSecond, uint256 _computePricePerSecond, address _provider, address _user) returns(uint256 leaseId)
func (_Gpulease *GpuleaseTransactor) StartLease(opts *bind.TransactOpts, _duration *big.Int, _storagePricePerSecond *big.Int, _computePricePerSecond *big.Int, _provider common.Address, _user common.Address) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "startLease", _duration, _storagePricePerSecond, _computePricePerSecond, _provider, _user)
}

// StartLease is a paid mutator transaction binding the contract method 0xf648da80.
//
// Solidity: function startLease(uint256 _duration, uint256 _storagePricePerSecond, uint256 _computePricePerSecond, address _provider, address _user) returns(uint256 leaseId)
func (_Gpulease *GpuleaseSession) StartLease(_duration *big.Int, _storagePricePerSecond *big.Int, _computePricePerSecond *big.Int, _provider common.Address, _user common.Address) (*types.Transaction, error) {
	return _Gpulease.Contract.StartLease(&_Gpulease.TransactOpts, _duration, _storagePricePerSecond, _computePricePerSecond, _provider, _user)
}

// StartLease is a paid mutator transaction binding the contract method 0xf648da80.
//
// Solidity: function startLease(uint256 _duration, uint256 _storagePricePerSecond, uint256 _computePricePerSecond, address _provider, address _user) returns(uint256 leaseId)
func (_Gpulease *GpuleaseTransactorSession) StartLease(_duration *big.Int, _storagePricePerSecond *big.Int, _computePricePerSecond *big.Int, _provider common.Address, _user common.Address) (*types.Transaction, error) {
	return _Gpulease.Contract.StartLease(&_Gpulease.TransactOpts, _duration, _storagePricePerSecond, _computePricePerSecond, _provider, _user)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gpulease *GpuleaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gpulease *GpuleaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gpulease.Contract.TransferOwnership(&_Gpulease.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gpulease *GpuleaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gpulease.Contract.TransferOwnership(&_Gpulease.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gpulease *GpuleaseTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gpulease.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gpulease *GpuleaseSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.Withdraw(&_Gpulease.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gpulease *GpuleaseTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Gpulease.Contract.Withdraw(&_Gpulease.TransactOpts, amount)
}

// GpuleaseDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Gpulease contract.
type GpuleaseDepositIterator struct {
	Event *GpuleaseDeposit // Event containing the contract specifics and raw log

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
func (it *GpuleaseDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseDeposit)
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
		it.Event = new(GpuleaseDeposit)
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
func (it *GpuleaseDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseDeposit represents a Deposit event raised by the Gpulease contract.
type GpuleaseDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Gpulease *GpuleaseFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*GpuleaseDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &GpuleaseDepositIterator{contract: _Gpulease.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Gpulease *GpuleaseFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GpuleaseDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseDeposit)
				if err := _Gpulease.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Gpulease *GpuleaseFilterer) ParseDeposit(log types.Log) (*GpuleaseDeposit, error) {
	event := new(GpuleaseDeposit)
	if err := _Gpulease.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuleaseLeaseCompletedIterator is returned from FilterLeaseCompleted and is used to iterate over the raw logs and unpacked data for LeaseCompleted events raised by the Gpulease contract.
type GpuleaseLeaseCompletedIterator struct {
	Event *GpuleaseLeaseCompleted // Event containing the contract specifics and raw log

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
func (it *GpuleaseLeaseCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseLeaseCompleted)
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
		it.Event = new(GpuleaseLeaseCompleted)
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
func (it *GpuleaseLeaseCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseLeaseCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseLeaseCompleted represents a LeaseCompleted event raised by the Gpulease contract.
type GpuleaseLeaseCompleted struct {
	LeaseId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLeaseCompleted is a free log retrieval operation binding the contract event 0x5a1241b1f059ce7cc7acee21bf522f44360638a58f17c58bd390a0f0af7e1937.
//
// Solidity: event LeaseCompleted(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) FilterLeaseCompleted(opts *bind.FilterOpts) (*GpuleaseLeaseCompletedIterator, error) {

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "LeaseCompleted")
	if err != nil {
		return nil, err
	}
	return &GpuleaseLeaseCompletedIterator{contract: _Gpulease.contract, event: "LeaseCompleted", logs: logs, sub: sub}, nil
}

// WatchLeaseCompleted is a free log subscription operation binding the contract event 0x5a1241b1f059ce7cc7acee21bf522f44360638a58f17c58bd390a0f0af7e1937.
//
// Solidity: event LeaseCompleted(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) WatchLeaseCompleted(opts *bind.WatchOpts, sink chan<- *GpuleaseLeaseCompleted) (event.Subscription, error) {

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "LeaseCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseLeaseCompleted)
				if err := _Gpulease.contract.UnpackLog(event, "LeaseCompleted", log); err != nil {
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

// ParseLeaseCompleted is a log parse operation binding the contract event 0x5a1241b1f059ce7cc7acee21bf522f44360638a58f17c58bd390a0f0af7e1937.
//
// Solidity: event LeaseCompleted(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) ParseLeaseCompleted(log types.Log) (*GpuleaseLeaseCompleted, error) {
	event := new(GpuleaseLeaseCompleted)
	if err := _Gpulease.contract.UnpackLog(event, "LeaseCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuleaseLeasePausedIterator is returned from FilterLeasePaused and is used to iterate over the raw logs and unpacked data for LeasePaused events raised by the Gpulease contract.
type GpuleaseLeasePausedIterator struct {
	Event *GpuleaseLeasePaused // Event containing the contract specifics and raw log

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
func (it *GpuleaseLeasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseLeasePaused)
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
		it.Event = new(GpuleaseLeasePaused)
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
func (it *GpuleaseLeasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseLeasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseLeasePaused represents a LeasePaused event raised by the Gpulease contract.
type GpuleaseLeasePaused struct {
	LeaseId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLeasePaused is a free log retrieval operation binding the contract event 0xb04140f2785e8aa46842b88fdbee148f5bcd51a8e9c4b367b7f431e80a6ed5a1.
//
// Solidity: event LeasePaused(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) FilterLeasePaused(opts *bind.FilterOpts) (*GpuleaseLeasePausedIterator, error) {

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "LeasePaused")
	if err != nil {
		return nil, err
	}
	return &GpuleaseLeasePausedIterator{contract: _Gpulease.contract, event: "LeasePaused", logs: logs, sub: sub}, nil
}

// WatchLeasePaused is a free log subscription operation binding the contract event 0xb04140f2785e8aa46842b88fdbee148f5bcd51a8e9c4b367b7f431e80a6ed5a1.
//
// Solidity: event LeasePaused(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) WatchLeasePaused(opts *bind.WatchOpts, sink chan<- *GpuleaseLeasePaused) (event.Subscription, error) {

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "LeasePaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseLeasePaused)
				if err := _Gpulease.contract.UnpackLog(event, "LeasePaused", log); err != nil {
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

// ParseLeasePaused is a log parse operation binding the contract event 0xb04140f2785e8aa46842b88fdbee148f5bcd51a8e9c4b367b7f431e80a6ed5a1.
//
// Solidity: event LeasePaused(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) ParseLeasePaused(log types.Log) (*GpuleaseLeasePaused, error) {
	event := new(GpuleaseLeasePaused)
	if err := _Gpulease.contract.UnpackLog(event, "LeasePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuleaseLeaseResumedIterator is returned from FilterLeaseResumed and is used to iterate over the raw logs and unpacked data for LeaseResumed events raised by the Gpulease contract.
type GpuleaseLeaseResumedIterator struct {
	Event *GpuleaseLeaseResumed // Event containing the contract specifics and raw log

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
func (it *GpuleaseLeaseResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseLeaseResumed)
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
		it.Event = new(GpuleaseLeaseResumed)
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
func (it *GpuleaseLeaseResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseLeaseResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseLeaseResumed represents a LeaseResumed event raised by the Gpulease contract.
type GpuleaseLeaseResumed struct {
	LeaseId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLeaseResumed is a free log retrieval operation binding the contract event 0x0ec1e889672cca90b2287dde9161b83295c31ef8ed2a8bd943a7f04b183f9167.
//
// Solidity: event LeaseResumed(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) FilterLeaseResumed(opts *bind.FilterOpts) (*GpuleaseLeaseResumedIterator, error) {

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "LeaseResumed")
	if err != nil {
		return nil, err
	}
	return &GpuleaseLeaseResumedIterator{contract: _Gpulease.contract, event: "LeaseResumed", logs: logs, sub: sub}, nil
}

// WatchLeaseResumed is a free log subscription operation binding the contract event 0x0ec1e889672cca90b2287dde9161b83295c31ef8ed2a8bd943a7f04b183f9167.
//
// Solidity: event LeaseResumed(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) WatchLeaseResumed(opts *bind.WatchOpts, sink chan<- *GpuleaseLeaseResumed) (event.Subscription, error) {

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "LeaseResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseLeaseResumed)
				if err := _Gpulease.contract.UnpackLog(event, "LeaseResumed", log); err != nil {
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

// ParseLeaseResumed is a log parse operation binding the contract event 0x0ec1e889672cca90b2287dde9161b83295c31ef8ed2a8bd943a7f04b183f9167.
//
// Solidity: event LeaseResumed(uint256 leaseId)
func (_Gpulease *GpuleaseFilterer) ParseLeaseResumed(log types.Log) (*GpuleaseLeaseResumed, error) {
	event := new(GpuleaseLeaseResumed)
	if err := _Gpulease.contract.UnpackLog(event, "LeaseResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuleaseLeaseStartedIterator is returned from FilterLeaseStarted and is used to iterate over the raw logs and unpacked data for LeaseStarted events raised by the Gpulease contract.
type GpuleaseLeaseStartedIterator struct {
	Event *GpuleaseLeaseStarted // Event containing the contract specifics and raw log

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
func (it *GpuleaseLeaseStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseLeaseStarted)
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
		it.Event = new(GpuleaseLeaseStarted)
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
func (it *GpuleaseLeaseStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseLeaseStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseLeaseStarted represents a LeaseStarted event raised by the Gpulease contract.
type GpuleaseLeaseStarted struct {
	LeaseId  *big.Int
	User     common.Address
	Provider common.Address
	Duration *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLeaseStarted is a free log retrieval operation binding the contract event 0x4501059f1dbad8e151132c2515ee7ea78d4b2540b9e56e941a03fdb9888efe28.
//
// Solidity: event LeaseStarted(uint256 leaseId, address user, address provider, uint256 duration, uint256 amount)
func (_Gpulease *GpuleaseFilterer) FilterLeaseStarted(opts *bind.FilterOpts) (*GpuleaseLeaseStartedIterator, error) {

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "LeaseStarted")
	if err != nil {
		return nil, err
	}
	return &GpuleaseLeaseStartedIterator{contract: _Gpulease.contract, event: "LeaseStarted", logs: logs, sub: sub}, nil
}

// WatchLeaseStarted is a free log subscription operation binding the contract event 0x4501059f1dbad8e151132c2515ee7ea78d4b2540b9e56e941a03fdb9888efe28.
//
// Solidity: event LeaseStarted(uint256 leaseId, address user, address provider, uint256 duration, uint256 amount)
func (_Gpulease *GpuleaseFilterer) WatchLeaseStarted(opts *bind.WatchOpts, sink chan<- *GpuleaseLeaseStarted) (event.Subscription, error) {

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "LeaseStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseLeaseStarted)
				if err := _Gpulease.contract.UnpackLog(event, "LeaseStarted", log); err != nil {
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

// ParseLeaseStarted is a log parse operation binding the contract event 0x4501059f1dbad8e151132c2515ee7ea78d4b2540b9e56e941a03fdb9888efe28.
//
// Solidity: event LeaseStarted(uint256 leaseId, address user, address provider, uint256 duration, uint256 amount)
func (_Gpulease *GpuleaseFilterer) ParseLeaseStarted(log types.Log) (*GpuleaseLeaseStarted, error) {
	event := new(GpuleaseLeaseStarted)
	if err := _Gpulease.contract.UnpackLog(event, "LeaseStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuleaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gpulease contract.
type GpuleaseOwnershipTransferredIterator struct {
	Event *GpuleaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GpuleaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseOwnershipTransferred)
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
		it.Event = new(GpuleaseOwnershipTransferred)
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
func (it *GpuleaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseOwnershipTransferred represents a OwnershipTransferred event raised by the Gpulease contract.
type GpuleaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gpulease *GpuleaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GpuleaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GpuleaseOwnershipTransferredIterator{contract: _Gpulease.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gpulease *GpuleaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GpuleaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseOwnershipTransferred)
				if err := _Gpulease.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gpulease *GpuleaseFilterer) ParseOwnershipTransferred(log types.Log) (*GpuleaseOwnershipTransferred, error) {
	event := new(GpuleaseOwnershipTransferred)
	if err := _Gpulease.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuleaseWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Gpulease contract.
type GpuleaseWithdrawIterator struct {
	Event *GpuleaseWithdraw // Event containing the contract specifics and raw log

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
func (it *GpuleaseWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuleaseWithdraw)
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
		it.Event = new(GpuleaseWithdraw)
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
func (it *GpuleaseWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuleaseWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuleaseWithdraw represents a Withdraw event raised by the Gpulease contract.
type GpuleaseWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Gpulease *GpuleaseFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*GpuleaseWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gpulease.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &GpuleaseWithdrawIterator{contract: _Gpulease.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Gpulease *GpuleaseFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GpuleaseWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gpulease.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuleaseWithdraw)
				if err := _Gpulease.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Gpulease *GpuleaseFilterer) ParseWithdraw(log types.Log) (*GpuleaseWithdraw, error) {
	event := new(GpuleaseWithdraw)
	if err := _Gpulease.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
