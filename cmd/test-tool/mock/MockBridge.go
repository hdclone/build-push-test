// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock

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

// MockMetaData contains all meta data concerning the Mock contract.
var MockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mpc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"mpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveRequestV2Signed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storedMpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610ced380380610ced833981810160405281019061003291906100db565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100a88261007d565b9050919050565b6100b88161009d565b81146100c357600080fd5b50565b6000815190506100d5816100af565b92915050565b6000602082840312156100f1576100f0610078565b5b60006100ff848285016100c6565b91505092915050565b610bd6806101176000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637c340c741461004657806384d61c9714610064578063f75c266414610080575b600080fd5b61004e61009e565b60405161005b9190610585565b60405180910390f35b61007e60048036038101906100799190610726565b6100c2565b005b610088610147565b6040516100959190610585565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b828260601b6040516020016100d892919061089e565b60405160208183030381529060405280519060200120816101016100fa610147565b8383610170565b610140576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013790610932565b60405180910390fd5b5050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080600061017f8585610355565b915091506000600481111561019757610196610952565b5b8160048111156101aa576101a9610952565b5b1480156101e257508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b156101f25760019250505061034e565b6000808773ffffffffffffffffffffffffffffffffffffffff16631626ba7e60e01b88886040516024016102279291906109e4565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102919190610a14565b600060405180830381855afa9150503d80600081146102cc576040519150601f19603f3d011682016040523d82523d6000602084013e6102d1565b606091505b50915091508180156102e4575060208151145b80156103475750631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916818060200190518101906103269190610a83565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9450505050505b9392505050565b6000806041835114156103975760008060006020860151925060408601519150606086015160001a905061038b878285856103d8565b945094505050506103d1565b6040835114156103c85760008060208501519150604085015190506103bd8683836104e5565b9350935050506103d1565b60006002915091505b9250929050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c11156104135760006003915091506104dc565b601b8560ff161415801561042b5750601c8560ff1614155b1561043d5760006004915091506104dc565b6000600187878787604051600081526020016040526040516104629493929190610acc565b6020604051602081039080840390855afa158015610484573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104d3576000600192509250506104dc565b80600092509250505b94509492505050565b60008060007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b841690506000601b60ff8660001c901c6105289190610b4a565b9050610536878288856103d8565b935093505050935093915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061056f82610544565b9050919050565b61057f81610564565b82525050565b600060208201905061059a6000830184610576565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610607826105be565b810181811067ffffffffffffffff82111715610626576106256105cf565b5b80604052505050565b60006106396105a0565b905061064582826105fe565b919050565b600067ffffffffffffffff821115610665576106646105cf565b5b61066e826105be565b9050602081019050919050565b82818337600083830152505050565b600061069d6106988461064a565b61062f565b9050828152602081018484840111156106b9576106b86105b9565b5b6106c484828561067b565b509392505050565b600082601f8301126106e1576106e06105b4565b5b81356106f184826020860161068a565b91505092915050565b61070381610564565b811461070e57600080fd5b50565b600081359050610720816106fa565b92915050565b60008060006060848603121561073f5761073e6105aa565b5b600084013567ffffffffffffffff81111561075d5761075c6105af565b5b610769868287016106cc565b935050602061077a86828701610711565b925050604084013567ffffffffffffffff81111561079b5761079a6105af565b5b6107a7868287016106cc565b9150509250925092565b7f7265636569766552657175657374563200000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b8381101561080b5780820151818401526020810190506107f0565b8381111561081a576000848401525b50505050565b600061082b826107d7565b61083581856107e2565b93506108458185602086016107ed565b80840191505092915050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b6000819050919050565b61089861089382610851565b61087d565b82525050565b60006108a9826107b1565b6010820191506108b98285610820565b91506108c58284610887565b6014820191508190509392505050565b600082825260208201905092915050565b7f4d6f636b4272696467653a20696e76616c6964207369676e6174757265000000600082015250565b600061091c601d836108d5565b9150610927826108e6565b602082019050919050565b6000602082019050818103600083015261094b8161090f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000819050919050565b61099481610981565b82525050565b600082825260208201905092915050565b60006109b6826107d7565b6109c0818561099a565b93506109d08185602086016107ed565b6109d9816105be565b840191505092915050565b60006040820190506109f9600083018561098b565b8181036020830152610a0b81846109ab565b90509392505050565b6000610a208284610820565b915081905092915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610a6081610a2b565b8114610a6b57600080fd5b50565b600081519050610a7d81610a57565b92915050565b600060208284031215610a9957610a986105aa565b5b6000610aa784828501610a6e565b91505092915050565b600060ff82169050919050565b610ac681610ab0565b82525050565b6000608082019050610ae1600083018761098b565b610aee6020830186610abd565b610afb604083018561098b565b610b08606083018461098b565b95945050505050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b5582610b11565b9150610b6083610b11565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610b9557610b94610b1b565b5b82820190509291505056fea2646970667358221220c5cdf58a44260938b63a2513aff130df712b501869e8e12a73334738bc4a71ff64736f6c634300080b0033",
}

// MockABI is the input ABI used to generate the binding from.
// Deprecated: Use MockMetaData.ABI instead.
var MockABI = MockMetaData.ABI

// MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockMetaData.Bin instead.
var MockBin = MockMetaData.Bin

// DeployMock deploys a new Ethereum contract, binding an instance of Mock to it.
func DeployMock(auth *bind.TransactOpts, backend bind.ContractBackend, _mpc common.Address) (common.Address, *types.Transaction, *Mock, error) {
	parsed, err := MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockBin), backend, _mpc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mock{MockCaller: MockCaller{contract: contract}, MockTransactor: MockTransactor{contract: contract}, MockFilterer: MockFilterer{contract: contract}}, nil
}

// Mock is an auto generated Go binding around an Ethereum contract.
type Mock struct {
	MockCaller     // Read-only binding to the contract
	MockTransactor // Write-only binding to the contract
	MockFilterer   // Log filterer for contract events
}

// MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockSession struct {
	Contract     *Mock             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Try options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockCallerSession struct {
	Contract *MockCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Try options to use throughout this session
}

// MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTransactorSession struct {
	Contract     *MockTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockRaw struct {
	Contract *Mock // Generic contract binding to access the raw methods on
}

// MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockCallerRaw struct {
	Contract *MockCaller // Generic read-only contract binding to access the raw methods on
}

// MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTransactorRaw struct {
	Contract *MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMock creates a new instance of Mock, bound to a specific deployed contract.
func NewMock(address common.Address, backend bind.ContractBackend) (*Mock, error) {
	contract, err := bindMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mock{MockCaller: MockCaller{contract: contract}, MockTransactor: MockTransactor{contract: contract}, MockFilterer: MockFilterer{contract: contract}}, nil
}

// NewMockCaller creates a new read-only instance of Mock, bound to a specific deployed contract.
func NewMockCaller(address common.Address, caller bind.ContractCaller) (*MockCaller, error) {
	contract, err := bindMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockCaller{contract: contract}, nil
}

// NewMockTransactor creates a new write-only instance of Mock, bound to a specific deployed contract.
func NewMockTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTransactor, error) {
	contract, err := bindMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTransactor{contract: contract}, nil
}

// NewMockFilterer creates a new log filterer instance of Mock, bound to a specific deployed contract.
func NewMockFilterer(address common.Address, filterer bind.ContractFilterer) (*MockFilterer, error) {
	contract, err := bindMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockFilterer{contract: contract}, nil
}

// bindMock binds a generic wrapper to an already deployed contract.
func bindMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mock *MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mock.Contract.MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mock *MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.Contract.MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mock *MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mock.Contract.MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mock *MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mock *MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mock *MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mock.Contract.contract.Transact(opts, method, params...)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Mock *MockCaller) Mpc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "mpc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Mock *MockSession) Mpc() (common.Address, error) {
	return _Mock.Contract.Mpc(&_Mock.CallOpts)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Mock *MockCallerSession) Mpc() (common.Address, error) {
	return _Mock.Contract.Mpc(&_Mock.CallOpts)
}

// StoredMpc is a free data retrieval call binding the contract method 0x7c340c74.
//
// Solidity: function storedMpc() view returns(address)
func (_Mock *MockCaller) StoredMpc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "storedMpc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StoredMpc is a free data retrieval call binding the contract method 0x7c340c74.
//
// Solidity: function storedMpc() view returns(address)
func (_Mock *MockSession) StoredMpc() (common.Address, error) {
	return _Mock.Contract.StoredMpc(&_Mock.CallOpts)
}

// StoredMpc is a free data retrieval call binding the contract method 0x7c340c74.
//
// Solidity: function storedMpc() view returns(address)
func (_Mock *MockCallerSession) StoredMpc() (common.Address, error) {
	return _Mock.Contract.StoredMpc(&_Mock.CallOpts)
}

// ReceiveRequestV2Signed is a paid mutator transaction binding the contract method 0x84d61c97.
//
// Solidity: function receiveRequestV2Signed(bytes _callData, address _receiveSide, bytes signature) returns()
func (_Mock *MockTransactor) ReceiveRequestV2Signed(opts *bind.TransactOpts, _callData []byte, _receiveSide common.Address, signature []byte) (*types.Transaction, error) {
	return _Mock.contract.Transact(opts, "receiveRequestV2Signed", _callData, _receiveSide, signature)
}

// ReceiveRequestV2Signed is a paid mutator transaction binding the contract method 0x84d61c97.
//
// Solidity: function receiveRequestV2Signed(bytes _callData, address _receiveSide, bytes signature) returns()
func (_Mock *MockSession) ReceiveRequestV2Signed(_callData []byte, _receiveSide common.Address, signature []byte) (*types.Transaction, error) {
	return _Mock.Contract.ReceiveRequestV2Signed(&_Mock.TransactOpts, _callData, _receiveSide, signature)
}

// ReceiveRequestV2Signed is a paid mutator transaction binding the contract method 0x84d61c97.
//
// Solidity: function receiveRequestV2Signed(bytes _callData, address _receiveSide, bytes signature) returns()
func (_Mock *MockTransactorSession) ReceiveRequestV2Signed(_callData []byte, _receiveSide common.Address, signature []byte) (*types.Transaction, error) {
	return _Mock.Contract.ReceiveRequestV2Signed(&_Mock.TransactOpts, _callData, _receiveSide, signature)
}
