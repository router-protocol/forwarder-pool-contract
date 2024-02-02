// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IMessageHandler

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

// IMessageHandlerMetaData contains all meta data concerning the IMessageHandler contract.
var IMessageHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"handleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMessageHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageHandlerMetaData.ABI instead.
var IMessageHandlerABI = IMessageHandlerMetaData.ABI

// IMessageHandler is an auto generated Go binding around an Ethereum contract.
type IMessageHandler struct {
	IMessageHandlerCaller     // Read-only binding to the contract
	IMessageHandlerTransactor // Write-only binding to the contract
	IMessageHandlerFilterer   // Log filterer for contract events
}

// IMessageHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageHandlerSession struct {
	Contract     *IMessageHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageHandlerCallerSession struct {
	Contract *IMessageHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IMessageHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageHandlerTransactorSession struct {
	Contract     *IMessageHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IMessageHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageHandlerRaw struct {
	Contract *IMessageHandler // Generic contract binding to access the raw methods on
}

// IMessageHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageHandlerCallerRaw struct {
	Contract *IMessageHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageHandlerTransactorRaw struct {
	Contract *IMessageHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageHandler creates a new instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandler(address common.Address, backend bind.ContractBackend) (*IMessageHandler, error) {
	contract, err := bindIMessageHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageHandler{IMessageHandlerCaller: IMessageHandlerCaller{contract: contract}, IMessageHandlerTransactor: IMessageHandlerTransactor{contract: contract}, IMessageHandlerFilterer: IMessageHandlerFilterer{contract: contract}}, nil
}

// NewIMessageHandlerCaller creates a new read-only instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerCaller(address common.Address, caller bind.ContractCaller) (*IMessageHandlerCaller, error) {
	contract, err := bindIMessageHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerCaller{contract: contract}, nil
}

// NewIMessageHandlerTransactor creates a new write-only instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageHandlerTransactor, error) {
	contract, err := bindIMessageHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerTransactor{contract: contract}, nil
}

// NewIMessageHandlerFilterer creates a new log filterer instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageHandlerFilterer, error) {
	contract, err := bindIMessageHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerFilterer{contract: contract}, nil
}

// bindIMessageHandler binds a generic wrapper to an already deployed contract.
func bindIMessageHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageHandler *IMessageHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageHandler.Contract.IMessageHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageHandler *IMessageHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageHandler.Contract.IMessageHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageHandler *IMessageHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageHandler.Contract.IMessageHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageHandler *IMessageHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageHandler *IMessageHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageHandler *IMessageHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageHandler.Contract.contract.Transact(opts, method, params...)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) returns()
func (_IMessageHandler *IMessageHandlerTransactor) HandleMessage(opts *bind.TransactOpts, tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _IMessageHandler.contract.Transact(opts, "handleMessage", tokenSent, amount, message)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) returns()
func (_IMessageHandler *IMessageHandlerSession) HandleMessage(tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _IMessageHandler.Contract.HandleMessage(&_IMessageHandler.TransactOpts, tokenSent, amount, message)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) returns()
func (_IMessageHandler *IMessageHandlerTransactorSession) HandleMessage(tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _IMessageHandler.Contract.HandleMessage(&_IMessageHandler.TransactOpts, tokenSent, amount, message)
}
