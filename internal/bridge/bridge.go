// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"LogChangeMPC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"SetAdminPermission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"SetTransmitterStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMPC\",\"type\":\"address\"}],\"name\":\"changeMPC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mpc\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTransmitter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newMPC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newMPCEffectiveTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldMPC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"}],\"name\":\"receiveRequestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveRequestV2Signed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_permission\",\"type\":\"bool\"}],\"name\":\"setAdminPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setTransmitterStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Try options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Try options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// CurrentChainId is a free data retrieval call binding the contract method 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (_Bridge *BridgeCaller) CurrentChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "currentChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentChainId is a free data retrieval call binding the contract method 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (_Bridge *BridgeSession) CurrentChainId() (*big.Int, error) {
	return _Bridge.Contract.CurrentChainId(&_Bridge.CallOpts)
}

// CurrentChainId is a free data retrieval call binding the contract method 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (_Bridge *BridgeCallerSession) CurrentChainId() (*big.Int, error) {
	return _Bridge.Contract.CurrentChainId(&_Bridge.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Bridge *BridgeCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Bridge *BridgeSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.IsAdmin(&_Bridge.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.IsAdmin(&_Bridge.CallOpts, arg0)
}

// IsTransmitter is a free data retrieval call binding the contract method 0x6fac3007.
//
// Solidity: function isTransmitter(address ) view returns(bool)
func (_Bridge *BridgeCaller) IsTransmitter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isTransmitter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransmitter is a free data retrieval call binding the contract method 0x6fac3007.
//
// Solidity: function isTransmitter(address ) view returns(bool)
func (_Bridge *BridgeSession) IsTransmitter(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.IsTransmitter(&_Bridge.CallOpts, arg0)
}

// IsTransmitter is a free data retrieval call binding the contract method 0x6fac3007.
//
// Solidity: function isTransmitter(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) IsTransmitter(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.IsTransmitter(&_Bridge.CallOpts, arg0)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Bridge *BridgeCaller) Mpc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "mpc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Bridge *BridgeSession) Mpc() (common.Address, error) {
	return _Bridge.Contract.Mpc(&_Bridge.CallOpts)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Bridge *BridgeCallerSession) Mpc() (common.Address, error) {
	return _Bridge.Contract.Mpc(&_Bridge.CallOpts)
}

// NewMPC is a free data retrieval call binding the contract method 0x474a245a.
//
// Solidity: function newMPC() view returns(address)
func (_Bridge *BridgeCaller) NewMPC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "newMPC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewMPC is a free data retrieval call binding the contract method 0x474a245a.
//
// Solidity: function newMPC() view returns(address)
func (_Bridge *BridgeSession) NewMPC() (common.Address, error) {
	return _Bridge.Contract.NewMPC(&_Bridge.CallOpts)
}

// NewMPC is a free data retrieval call binding the contract method 0x474a245a.
//
// Solidity: function newMPC() view returns(address)
func (_Bridge *BridgeCallerSession) NewMPC() (common.Address, error) {
	return _Bridge.Contract.NewMPC(&_Bridge.CallOpts)
}

// NewMPCEffectiveTime is a free data retrieval call binding the contract method 0x405fb4f7.
//
// Solidity: function newMPCEffectiveTime() view returns(uint256)
func (_Bridge *BridgeCaller) NewMPCEffectiveTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "newMPCEffectiveTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewMPCEffectiveTime is a free data retrieval call binding the contract method 0x405fb4f7.
//
// Solidity: function newMPCEffectiveTime() view returns(uint256)
func (_Bridge *BridgeSession) NewMPCEffectiveTime() (*big.Int, error) {
	return _Bridge.Contract.NewMPCEffectiveTime(&_Bridge.CallOpts)
}

// NewMPCEffectiveTime is a free data retrieval call binding the contract method 0x405fb4f7.
//
// Solidity: function newMPCEffectiveTime() view returns(uint256)
func (_Bridge *BridgeCallerSession) NewMPCEffectiveTime() (*big.Int, error) {
	return _Bridge.Contract.NewMPCEffectiveTime(&_Bridge.CallOpts)
}

// OldMPC is a free data retrieval call binding the contract method 0xc00f8a3d.
//
// Solidity: function oldMPC() view returns(address)
func (_Bridge *BridgeCaller) OldMPC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "oldMPC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldMPC is a free data retrieval call binding the contract method 0xc00f8a3d.
//
// Solidity: function oldMPC() view returns(address)
func (_Bridge *BridgeSession) OldMPC() (common.Address, error) {
	return _Bridge.Contract.OldMPC(&_Bridge.CallOpts)
}

// OldMPC is a free data retrieval call binding the contract method 0xc00f8a3d.
//
// Solidity: function oldMPC() view returns(address)
func (_Bridge *BridgeCallerSession) OldMPC() (common.Address, error) {
	return _Bridge.Contract.OldMPC(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address _newMPC) returns(bool)
func (_Bridge *BridgeTransactor) ChangeMPC(opts *bind.TransactOpts, _newMPC common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeMPC", _newMPC)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address _newMPC) returns(bool)
func (_Bridge *BridgeSession) ChangeMPC(_newMPC common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeMPC(&_Bridge.TransactOpts, _newMPC)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address _newMPC) returns(bool)
func (_Bridge *BridgeTransactorSession) ChangeMPC(_newMPC common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeMPC(&_Bridge.TransactOpts, _newMPC)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _mpc) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, _mpc common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", _mpc)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _mpc) returns()
func (_Bridge *BridgeSession) Initialize(_mpc common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, _mpc)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _mpc) returns()
func (_Bridge *BridgeTransactorSession) Initialize(_mpc common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, _mpc)
}

// ReceiveRequestV2 is a paid mutator transaction binding the contract method 0xf7f1baf0.
//
// Solidity: function receiveRequestV2(bytes _callData, address _receiveSide) returns()
func (_Bridge *BridgeTransactor) ReceiveRequestV2(opts *bind.TransactOpts, _callData []byte, _receiveSide common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "receiveRequestV2", _callData, _receiveSide)
}

// ReceiveRequestV2 is a paid mutator transaction binding the contract method 0xf7f1baf0.
//
// Solidity: function receiveRequestV2(bytes _callData, address _receiveSide) returns()
func (_Bridge *BridgeSession) ReceiveRequestV2(_callData []byte, _receiveSide common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveRequestV2(&_Bridge.TransactOpts, _callData, _receiveSide)
}

// ReceiveRequestV2 is a paid mutator transaction binding the contract method 0xf7f1baf0.
//
// Solidity: function receiveRequestV2(bytes _callData, address _receiveSide) returns()
func (_Bridge *BridgeTransactorSession) ReceiveRequestV2(_callData []byte, _receiveSide common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveRequestV2(&_Bridge.TransactOpts, _callData, _receiveSide)
}

// ReceiveRequestV2Signed is a paid mutator transaction binding the contract method 0x84d61c97.
//
// Solidity: function receiveRequestV2Signed(bytes _callData, address _receiveSide, bytes signature) returns()
func (_Bridge *BridgeTransactor) ReceiveRequestV2Signed(opts *bind.TransactOpts, _callData []byte, _receiveSide common.Address, signature []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "receiveRequestV2Signed", _callData, _receiveSide, signature)
}

// ReceiveRequestV2Signed is a paid mutator transaction binding the contract method 0x84d61c97.
//
// Solidity: function receiveRequestV2Signed(bytes _callData, address _receiveSide, bytes signature) returns()
func (_Bridge *BridgeSession) ReceiveRequestV2Signed(_callData []byte, _receiveSide common.Address, signature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveRequestV2Signed(&_Bridge.TransactOpts, _callData, _receiveSide, signature)
}

// ReceiveRequestV2Signed is a paid mutator transaction binding the contract method 0x84d61c97.
//
// Solidity: function receiveRequestV2Signed(bytes _callData, address _receiveSide, bytes signature) returns()
func (_Bridge *BridgeTransactorSession) ReceiveRequestV2Signed(_callData []byte, _receiveSide common.Address, signature []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveRequestV2Signed(&_Bridge.TransactOpts, _callData, _receiveSide, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetAdminPermission is a paid mutator transaction binding the contract method 0x75f3974b.
//
// Solidity: function setAdminPermission(address _user, bool _permission) returns()
func (_Bridge *BridgeTransactor) SetAdminPermission(opts *bind.TransactOpts, _user common.Address, _permission bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAdminPermission", _user, _permission)
}

// SetAdminPermission is a paid mutator transaction binding the contract method 0x75f3974b.
//
// Solidity: function setAdminPermission(address _user, bool _permission) returns()
func (_Bridge *BridgeSession) SetAdminPermission(_user common.Address, _permission bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdminPermission(&_Bridge.TransactOpts, _user, _permission)
}

// SetAdminPermission is a paid mutator transaction binding the contract method 0x75f3974b.
//
// Solidity: function setAdminPermission(address _user, bool _permission) returns()
func (_Bridge *BridgeTransactorSession) SetAdminPermission(_user common.Address, _permission bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdminPermission(&_Bridge.TransactOpts, _user, _permission)
}

// SetTransmitterStatus is a paid mutator transaction binding the contract method 0x19117d93.
//
// Solidity: function setTransmitterStatus(address _transmitter, bool _status) returns()
func (_Bridge *BridgeTransactor) SetTransmitterStatus(opts *bind.TransactOpts, _transmitter common.Address, _status bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTransmitterStatus", _transmitter, _status)
}

// SetTransmitterStatus is a paid mutator transaction binding the contract method 0x19117d93.
//
// Solidity: function setTransmitterStatus(address _transmitter, bool _status) returns()
func (_Bridge *BridgeSession) SetTransmitterStatus(_transmitter common.Address, _status bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTransmitterStatus(&_Bridge.TransactOpts, _transmitter, _status)
}

// SetTransmitterStatus is a paid mutator transaction binding the contract method 0x19117d93.
//
// Solidity: function setTransmitterStatus(address _transmitter, bool _status) returns()
func (_Bridge *BridgeTransactorSession) SetTransmitterStatus(_transmitter common.Address, _status bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTransmitterStatus(&_Bridge.TransactOpts, _transmitter, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransmitRequestV2 is a paid mutator transaction binding the contract method 0x6cebc9c2.
//
// Solidity: function transmitRequestV2(bytes _callData, address _receiveSide, address _oppositeBridge, uint256 _chainId) returns()
func (_Bridge *BridgeTransactor) TransmitRequestV2(opts *bind.TransactOpts, _callData []byte, _receiveSide common.Address, _oppositeBridge common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transmitRequestV2", _callData, _receiveSide, _oppositeBridge, _chainId)
}

// TransmitRequestV2 is a paid mutator transaction binding the contract method 0x6cebc9c2.
//
// Solidity: function transmitRequestV2(bytes _callData, address _receiveSide, address _oppositeBridge, uint256 _chainId) returns()
func (_Bridge *BridgeSession) TransmitRequestV2(_callData []byte, _receiveSide common.Address, _oppositeBridge common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransmitRequestV2(&_Bridge.TransactOpts, _callData, _receiveSide, _oppositeBridge, _chainId)
}

// TransmitRequestV2 is a paid mutator transaction binding the contract method 0x6cebc9c2.
//
// Solidity: function transmitRequestV2(bytes _callData, address _receiveSide, address _oppositeBridge, uint256 _chainId) returns()
func (_Bridge *BridgeTransactorSession) TransmitRequestV2(_callData []byte, _receiveSide common.Address, _oppositeBridge common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransmitRequestV2(&_Bridge.TransactOpts, _callData, _receiveSide, _oppositeBridge, _chainId)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactor) WithdrawFee(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawFee", token, to, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amount) returns(bool)
func (_Bridge *BridgeSession) WithdrawFee(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawFee(&_Bridge.TransactOpts, token, to, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactorSession) WithdrawFee(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawFee(&_Bridge.TransactOpts, token, to, amount)
}

// BridgeLogChangeMPCIterator is returned from FilterLogChangeMPC and is used to iterate over the raw logs and unpacked data for LogChangeMPC events raised by the Bridge contract.
type BridgeLogChangeMPCIterator struct {
	Event *BridgeLogChangeMPC // Event containing the contract specifics and raw log

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
func (it *BridgeLogChangeMPCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogChangeMPC)
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
		it.Event = new(BridgeLogChangeMPC)
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
func (it *BridgeLogChangeMPCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogChangeMPCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogChangeMPC represents a LogChangeMPC event raised by the Bridge contract.
type BridgeLogChangeMPC struct {
	OldMPC        common.Address
	NewMPC        common.Address
	EffectiveTime *big.Int
	ChainId       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogChangeMPC is a free log retrieval operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainId)
func (_Bridge *BridgeFilterer) FilterLogChangeMPC(opts *bind.FilterOpts, oldMPC []common.Address, newMPC []common.Address, effectiveTime []*big.Int) (*BridgeLogChangeMPCIterator, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLogChangeMPCIterator{contract: _Bridge.contract, event: "LogChangeMPC", logs: logs, sub: sub}, nil
}

// WatchLogChangeMPC is a free log subscription operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainId)
func (_Bridge *BridgeFilterer) WatchLogChangeMPC(opts *bind.WatchOpts, sink chan<- *BridgeLogChangeMPC, oldMPC []common.Address, newMPC []common.Address, effectiveTime []*big.Int) (event.Subscription, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogChangeMPC)
				if err := _Bridge.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
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

// ParseLogChangeMPC is a log parse operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainId)
func (_Bridge *BridgeFilterer) ParseLogChangeMPC(log types.Log) (*BridgeLogChangeMPC, error) {
	event := new(BridgeLogChangeMPC)
	if err := _Bridge.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the Bridge contract.
type BridgeOracleRequestIterator struct {
	Event *BridgeOracleRequest // Event containing the contract specifics and raw log

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
func (it *BridgeOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOracleRequest)
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
		it.Event = new(BridgeOracleRequest)
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
func (it *BridgeOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOracleRequest represents a OracleRequest event raised by the Bridge contract.
type BridgeOracleRequest struct {
	Bridge         common.Address
	CallData       []byte
	ReceiveSide    common.Address
	OppositeBridge common.Address
	ChainId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0x532dbb6d061eee97ab4370060f60ede10b3dc361cc1214c07ae5e34dd86e6aaf.
//
// Solidity: event OracleRequest(address bridge, bytes callData, address receiveSide, address oppositeBridge, uint256 chainId)
func (_Bridge *BridgeFilterer) FilterOracleRequest(opts *bind.FilterOpts) (*BridgeOracleRequestIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return &BridgeOracleRequestIterator{contract: _Bridge.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0x532dbb6d061eee97ab4370060f60ede10b3dc361cc1214c07ae5e34dd86e6aaf.
//
// Solidity: event OracleRequest(address bridge, bytes callData, address receiveSide, address oppositeBridge, uint256 chainId)
func (_Bridge *BridgeFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *BridgeOracleRequest) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOracleRequest)
				if err := _Bridge.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

// ParseOracleRequest is a log parse operation binding the contract event 0x532dbb6d061eee97ab4370060f60ede10b3dc361cc1214c07ae5e34dd86e6aaf.
//
// Solidity: event OracleRequest(address bridge, bytes callData, address receiveSide, address oppositeBridge, uint256 chainId)
func (_Bridge *BridgeFilterer) ParseOracleRequest(log types.Log) (*BridgeOracleRequest, error) {
	event := new(BridgeOracleRequest)
	if err := _Bridge.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetAdminPermissionIterator is returned from FilterSetAdminPermission and is used to iterate over the raw logs and unpacked data for SetAdminPermission events raised by the Bridge contract.
type BridgeSetAdminPermissionIterator struct {
	Event *BridgeSetAdminPermission // Event containing the contract specifics and raw log

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
func (it *BridgeSetAdminPermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetAdminPermission)
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
		it.Event = new(BridgeSetAdminPermission)
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
func (it *BridgeSetAdminPermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetAdminPermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetAdminPermission represents a SetAdminPermission event raised by the Bridge contract.
type BridgeSetAdminPermission struct {
	Admin      common.Address
	Permission bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetAdminPermission is a free log retrieval operation binding the contract event 0x0e7bea53cb2b3130dd1aac8d56b61cc8da7ebab0432e2d1622513523d848f2e7.
//
// Solidity: event SetAdminPermission(address indexed admin, bool permission)
func (_Bridge *BridgeFilterer) FilterSetAdminPermission(opts *bind.FilterOpts, admin []common.Address) (*BridgeSetAdminPermissionIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetAdminPermission", adminRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSetAdminPermissionIterator{contract: _Bridge.contract, event: "SetAdminPermission", logs: logs, sub: sub}, nil
}

// WatchSetAdminPermission is a free log subscription operation binding the contract event 0x0e7bea53cb2b3130dd1aac8d56b61cc8da7ebab0432e2d1622513523d848f2e7.
//
// Solidity: event SetAdminPermission(address indexed admin, bool permission)
func (_Bridge *BridgeFilterer) WatchSetAdminPermission(opts *bind.WatchOpts, sink chan<- *BridgeSetAdminPermission, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetAdminPermission", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetAdminPermission)
				if err := _Bridge.contract.UnpackLog(event, "SetAdminPermission", log); err != nil {
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

// ParseSetAdminPermission is a log parse operation binding the contract event 0x0e7bea53cb2b3130dd1aac8d56b61cc8da7ebab0432e2d1622513523d848f2e7.
//
// Solidity: event SetAdminPermission(address indexed admin, bool permission)
func (_Bridge *BridgeFilterer) ParseSetAdminPermission(log types.Log) (*BridgeSetAdminPermission, error) {
	event := new(BridgeSetAdminPermission)
	if err := _Bridge.contract.UnpackLog(event, "SetAdminPermission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSetTransmitterStatusIterator is returned from FilterSetTransmitterStatus and is used to iterate over the raw logs and unpacked data for SetTransmitterStatus events raised by the Bridge contract.
type BridgeSetTransmitterStatusIterator struct {
	Event *BridgeSetTransmitterStatus // Event containing the contract specifics and raw log

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
func (it *BridgeSetTransmitterStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSetTransmitterStatus)
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
		it.Event = new(BridgeSetTransmitterStatus)
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
func (it *BridgeSetTransmitterStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSetTransmitterStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSetTransmitterStatus represents a SetTransmitterStatus event raised by the Bridge contract.
type BridgeSetTransmitterStatus struct {
	Transmitter common.Address
	Status      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetTransmitterStatus is a free log retrieval operation binding the contract event 0xeeec8b4e2d317fc608f301f859237a6081b9813f150a3fcfb02fd54276c8be40.
//
// Solidity: event SetTransmitterStatus(address indexed transmitter, bool status)
func (_Bridge *BridgeFilterer) FilterSetTransmitterStatus(opts *bind.FilterOpts, transmitter []common.Address) (*BridgeSetTransmitterStatusIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SetTransmitterStatus", transmitterRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSetTransmitterStatusIterator{contract: _Bridge.contract, event: "SetTransmitterStatus", logs: logs, sub: sub}, nil
}

// WatchSetTransmitterStatus is a free log subscription operation binding the contract event 0xeeec8b4e2d317fc608f301f859237a6081b9813f150a3fcfb02fd54276c8be40.
//
// Solidity: event SetTransmitterStatus(address indexed transmitter, bool status)
func (_Bridge *BridgeFilterer) WatchSetTransmitterStatus(opts *bind.WatchOpts, sink chan<- *BridgeSetTransmitterStatus, transmitter []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SetTransmitterStatus", transmitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSetTransmitterStatus)
				if err := _Bridge.contract.UnpackLog(event, "SetTransmitterStatus", log); err != nil {
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

// ParseSetTransmitterStatus is a log parse operation binding the contract event 0xeeec8b4e2d317fc608f301f859237a6081b9813f150a3fcfb02fd54276c8be40.
//
// Solidity: event SetTransmitterStatus(address indexed transmitter, bool status)
func (_Bridge *BridgeFilterer) ParseSetTransmitterStatus(log types.Log) (*BridgeSetTransmitterStatus, error) {
	event := new(BridgeSetTransmitterStatus)
	if err := _Bridge.contract.UnpackLog(event, "SetTransmitterStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
