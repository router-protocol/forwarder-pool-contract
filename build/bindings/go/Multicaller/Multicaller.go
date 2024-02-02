// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Multicaller

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

// MulticallerMetaData contains all meta data concerning the Multicaller contract.
var MulticallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061048f806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ac9650d814610030575b600080fd5b61004361003e3660046101ba565b610059565b604051610050919061027f565b60405180910390f35b60608167ffffffffffffffff811115610074576100746102e1565b6040519080825280602002602001820160405280156100a757816020015b60608152602001906001900390816100925790505b50905060005b828110156101b357600080308686858181106100cb576100cb6102f7565b90506020028101906100dd919061030d565b6040516100eb92919061035b565b600060405180830381855af49150503d8060008114610126576040519150601f19603f3d011682016040523d82523d6000602084013e61012b565b606091505b5091509150816101805760448151101561014457600080fd5b6004810190508080602001905181019061015e919061036b565b60405162461bcd60e51b81526004016101779190610418565b60405180910390fd5b80848481518110610193576101936102f7565b6020026020010181905250505080806101ab90610432565b9150506100ad565b5092915050565b600080602083850312156101cd57600080fd5b823567ffffffffffffffff808211156101e557600080fd5b818501915085601f8301126101f957600080fd5b81358181111561020857600080fd5b8660208260051b850101111561021d57600080fd5b60209290920196919550909350505050565b60005b8381101561024a578181015183820152602001610232565b50506000910152565b6000815180845261026b81602086016020860161022f565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156102d457603f198886030184526102c2858351610253565b945092850192908501906001016102a6565b5092979650505050505050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261032457600080fd5b83018035915067ffffffffffffffff82111561033f57600080fd5b60200191503681900382131561035457600080fd5b9250929050565b8183823760009101908152919050565b60006020828403121561037d57600080fd5b815167ffffffffffffffff8082111561039557600080fd5b818401915084601f8301126103a957600080fd5b8151818111156103bb576103bb6102e1565b604051601f8201601f19908116603f011681019083821181831017156103e3576103e36102e1565b816040528281528760208487010111156103fc57600080fd5b61040d83602083016020880161022f565b979650505050505050565b60208152600061042b6020830184610253565b9392505050565b60006001820161045257634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220e57d16eb18117d34c29af49556a408a8c2e19f6165763fdf4b8aa64575aa579464736f6c63430008140033",
}

// MulticallerABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallerMetaData.ABI instead.
var MulticallerABI = MulticallerMetaData.ABI

// MulticallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MulticallerMetaData.Bin instead.
var MulticallerBin = MulticallerMetaData.Bin

// DeployMulticaller deploys a new Ethereum contract, binding an instance of Multicaller to it.
func DeployMulticaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multicaller, error) {
	parsed, err := MulticallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MulticallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multicaller{MulticallerCaller: MulticallerCaller{contract: contract}, MulticallerTransactor: MulticallerTransactor{contract: contract}, MulticallerFilterer: MulticallerFilterer{contract: contract}}, nil
}

// Multicaller is an auto generated Go binding around an Ethereum contract.
type Multicaller struct {
	MulticallerCaller     // Read-only binding to the contract
	MulticallerTransactor // Write-only binding to the contract
	MulticallerFilterer   // Log filterer for contract events
}

// MulticallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallerSession struct {
	Contract     *Multicaller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallerCallerSession struct {
	Contract *MulticallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MulticallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallerTransactorSession struct {
	Contract     *MulticallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MulticallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallerRaw struct {
	Contract *Multicaller // Generic contract binding to access the raw methods on
}

// MulticallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallerCallerRaw struct {
	Contract *MulticallerCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallerTransactorRaw struct {
	Contract *MulticallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticaller creates a new instance of Multicaller, bound to a specific deployed contract.
func NewMulticaller(address common.Address, backend bind.ContractBackend) (*Multicaller, error) {
	contract, err := bindMulticaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicaller{MulticallerCaller: MulticallerCaller{contract: contract}, MulticallerTransactor: MulticallerTransactor{contract: contract}, MulticallerFilterer: MulticallerFilterer{contract: contract}}, nil
}

// NewMulticallerCaller creates a new read-only instance of Multicaller, bound to a specific deployed contract.
func NewMulticallerCaller(address common.Address, caller bind.ContractCaller) (*MulticallerCaller, error) {
	contract, err := bindMulticaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallerCaller{contract: contract}, nil
}

// NewMulticallerTransactor creates a new write-only instance of Multicaller, bound to a specific deployed contract.
func NewMulticallerTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallerTransactor, error) {
	contract, err := bindMulticaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallerTransactor{contract: contract}, nil
}

// NewMulticallerFilterer creates a new log filterer instance of Multicaller, bound to a specific deployed contract.
func NewMulticallerFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallerFilterer, error) {
	contract, err := bindMulticaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallerFilterer{contract: contract}, nil
}

// bindMulticaller binds a generic wrapper to an already deployed contract.
func bindMulticaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicaller *MulticallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicaller.Contract.MulticallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicaller *MulticallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicaller.Contract.MulticallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicaller *MulticallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicaller.Contract.MulticallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicaller *MulticallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicaller *MulticallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicaller *MulticallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicaller.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multicaller *MulticallerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Multicaller.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multicaller *MulticallerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicaller.Contract.Multicall(&_Multicaller.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multicaller *MulticallerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicaller.Contract.Multicall(&_Multicaller.TransactOpts, data)
}
