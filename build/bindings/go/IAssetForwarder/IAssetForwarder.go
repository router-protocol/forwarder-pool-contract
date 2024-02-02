// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAssetForwarder

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

// IAssetForwarderDepositData is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderDepositData struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestAmount       *big.Int
	SrcToken         common.Address
	RefundRecipient  common.Address
	DestChainIdBytes [32]byte
}

// IAssetForwarderRelayData is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderRelayData struct {
	Amount     *big.Int
	SrcChainId [32]byte
	DepositId  *big.Int
	DestToken  common.Address
	Recipient  common.Address
}

// IAssetForwarderRelayDataMessage is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderRelayDataMessage struct {
	Amount     *big.Int
	SrcChainId [32]byte
	DepositId  *big.Int
	DestToken  common.Address
	Recipient  common.Address
	Message    []byte
}

// IAssetForwarderMetaData contains all meta data concerning the IAssetForwarder contract.
var IAssetForwarderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"DepositInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDepositedWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"FundsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"execFlag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"execData\",\"type\":\"bytes\"}],\"name\":\"FundsPaidWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"iDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"}],\"name\":\"iDepositInfoUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"iDepositMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"iDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IAssetForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use IAssetForwarderMetaData.ABI instead.
var IAssetForwarderABI = IAssetForwarderMetaData.ABI

// IAssetForwarder is an auto generated Go binding around an Ethereum contract.
type IAssetForwarder struct {
	IAssetForwarderCaller     // Read-only binding to the contract
	IAssetForwarderTransactor // Write-only binding to the contract
	IAssetForwarderFilterer   // Log filterer for contract events
}

// IAssetForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAssetForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAssetForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAssetForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAssetForwarderSession struct {
	Contract     *IAssetForwarder  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAssetForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAssetForwarderCallerSession struct {
	Contract *IAssetForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAssetForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAssetForwarderTransactorSession struct {
	Contract     *IAssetForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAssetForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAssetForwarderRaw struct {
	Contract *IAssetForwarder // Generic contract binding to access the raw methods on
}

// IAssetForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAssetForwarderCallerRaw struct {
	Contract *IAssetForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// IAssetForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAssetForwarderTransactorRaw struct {
	Contract *IAssetForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAssetForwarder creates a new instance of IAssetForwarder, bound to a specific deployed contract.
func NewIAssetForwarder(address common.Address, backend bind.ContractBackend) (*IAssetForwarder, error) {
	contract, err := bindIAssetForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAssetForwarder{IAssetForwarderCaller: IAssetForwarderCaller{contract: contract}, IAssetForwarderTransactor: IAssetForwarderTransactor{contract: contract}, IAssetForwarderFilterer: IAssetForwarderFilterer{contract: contract}}, nil
}

// NewIAssetForwarderCaller creates a new read-only instance of IAssetForwarder, bound to a specific deployed contract.
func NewIAssetForwarderCaller(address common.Address, caller bind.ContractCaller) (*IAssetForwarderCaller, error) {
	contract, err := bindIAssetForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderCaller{contract: contract}, nil
}

// NewIAssetForwarderTransactor creates a new write-only instance of IAssetForwarder, bound to a specific deployed contract.
func NewIAssetForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*IAssetForwarderTransactor, error) {
	contract, err := bindIAssetForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderTransactor{contract: contract}, nil
}

// NewIAssetForwarderFilterer creates a new log filterer instance of IAssetForwarder, bound to a specific deployed contract.
func NewIAssetForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*IAssetForwarderFilterer, error) {
	contract, err := bindIAssetForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderFilterer{contract: contract}, nil
}

// bindIAssetForwarder binds a generic wrapper to an already deployed contract.
func bindIAssetForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAssetForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssetForwarder *IAssetForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssetForwarder.Contract.IAssetForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssetForwarder *IAssetForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IAssetForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssetForwarder *IAssetForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IAssetForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssetForwarder *IAssetForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssetForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssetForwarder *IAssetForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssetForwarder *IAssetForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.contract.Transact(opts, method, params...)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactor) IDeposit(opts *bind.TransactOpts, depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _IAssetForwarder.contract.Transact(opts, "iDeposit", depositData, destToken, recipient)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_IAssetForwarder *IAssetForwarderSession) IDeposit(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDeposit(&_IAssetForwarder.TransactOpts, depositData, destToken, recipient)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactorSession) IDeposit(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDeposit(&_IAssetForwarder.TransactOpts, depositData, destToken, recipient)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactor) IDepositInfoUpdate(opts *bind.TransactOpts, srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _IAssetForwarder.contract.Transact(opts, "iDepositInfoUpdate", srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_IAssetForwarder *IAssetForwarderSession) IDepositInfoUpdate(srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDepositInfoUpdate(&_IAssetForwarder.TransactOpts, srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactorSession) IDepositInfoUpdate(srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDepositInfoUpdate(&_IAssetForwarder.TransactOpts, srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactor) IDepositMessage(opts *bind.TransactOpts, depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _IAssetForwarder.contract.Transact(opts, "iDepositMessage", depositData, destToken, recipient, message)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_IAssetForwarder *IAssetForwarderSession) IDepositMessage(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDepositMessage(&_IAssetForwarder.TransactOpts, depositData, destToken, recipient, message)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactorSession) IDepositMessage(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDepositMessage(&_IAssetForwarder.TransactOpts, depositData, destToken, recipient, message)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactor) IDepositUSDC(opts *bind.TransactOpts, partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _IAssetForwarder.contract.Transact(opts, "iDepositUSDC", partnerId, destChainIdBytes, recipient, amount)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_IAssetForwarder *IAssetForwarderSession) IDepositUSDC(partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDepositUSDC(&_IAssetForwarder.TransactOpts, partnerId, destChainIdBytes, recipient, amount)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactorSession) IDepositUSDC(partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IDepositUSDC(&_IAssetForwarder.TransactOpts, partnerId, destChainIdBytes, recipient, amount)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactor) IRelay(opts *bind.TransactOpts, relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _IAssetForwarder.contract.Transact(opts, "iRelay", relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_IAssetForwarder *IAssetForwarderSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IRelay(&_IAssetForwarder.TransactOpts, relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactorSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IRelay(&_IAssetForwarder.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactor) IRelayMessage(opts *bind.TransactOpts, relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _IAssetForwarder.contract.Transact(opts, "iRelayMessage", relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_IAssetForwarder *IAssetForwarderSession) IRelayMessage(relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IRelayMessage(&_IAssetForwarder.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_IAssetForwarder *IAssetForwarderTransactorSession) IRelayMessage(relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _IAssetForwarder.Contract.IRelayMessage(&_IAssetForwarder.TransactOpts, relayData)
}

// IAssetForwarderDepositInfoUpdateIterator is returned from FilterDepositInfoUpdate and is used to iterate over the raw logs and unpacked data for DepositInfoUpdate events raised by the IAssetForwarder contract.
type IAssetForwarderDepositInfoUpdateIterator struct {
	Event *IAssetForwarderDepositInfoUpdate // Event containing the contract specifics and raw log

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
func (it *IAssetForwarderDepositInfoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetForwarderDepositInfoUpdate)
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
		it.Event = new(IAssetForwarderDepositInfoUpdate)
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
func (it *IAssetForwarderDepositInfoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetForwarderDepositInfoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetForwarderDepositInfoUpdate represents a DepositInfoUpdate event raised by the IAssetForwarder contract.
type IAssetForwarderDepositInfoUpdate struct {
	SrcToken           common.Address
	FeeAmount          *big.Int
	DepositId          *big.Int
	EventNonce         *big.Int
	Initiatewithdrawal bool
	Depositor          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositInfoUpdate is a free log retrieval operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_IAssetForwarder *IAssetForwarderFilterer) FilterDepositInfoUpdate(opts *bind.FilterOpts) (*IAssetForwarderDepositInfoUpdateIterator, error) {

	logs, sub, err := _IAssetForwarder.contract.FilterLogs(opts, "DepositInfoUpdate")
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderDepositInfoUpdateIterator{contract: _IAssetForwarder.contract, event: "DepositInfoUpdate", logs: logs, sub: sub}, nil
}

// WatchDepositInfoUpdate is a free log subscription operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_IAssetForwarder *IAssetForwarderFilterer) WatchDepositInfoUpdate(opts *bind.WatchOpts, sink chan<- *IAssetForwarderDepositInfoUpdate) (event.Subscription, error) {

	logs, sub, err := _IAssetForwarder.contract.WatchLogs(opts, "DepositInfoUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetForwarderDepositInfoUpdate)
				if err := _IAssetForwarder.contract.UnpackLog(event, "DepositInfoUpdate", log); err != nil {
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

// ParseDepositInfoUpdate is a log parse operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_IAssetForwarder *IAssetForwarderFilterer) ParseDepositInfoUpdate(log types.Log) (*IAssetForwarderDepositInfoUpdate, error) {
	event := new(IAssetForwarderDepositInfoUpdate)
	if err := _IAssetForwarder.contract.UnpackLog(event, "DepositInfoUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetForwarderFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the IAssetForwarder contract.
type IAssetForwarderFundsDepositedIterator struct {
	Event *IAssetForwarderFundsDeposited // Event containing the contract specifics and raw log

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
func (it *IAssetForwarderFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetForwarderFundsDeposited)
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
		it.Event = new(IAssetForwarderFundsDeposited)
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
func (it *IAssetForwarderFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetForwarderFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetForwarderFundsDeposited represents a FundsDeposited event raised by the IAssetForwarder contract.
type IAssetForwarderFundsDeposited struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	DestAmount       *big.Int
	DepositId        *big.Int
	SrcToken         common.Address
	Depositor        common.Address
	Recipient        []byte
	DestToken        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_IAssetForwarder *IAssetForwarderFilterer) FilterFundsDeposited(opts *bind.FilterOpts) (*IAssetForwarderFundsDepositedIterator, error) {

	logs, sub, err := _IAssetForwarder.contract.FilterLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderFundsDepositedIterator{contract: _IAssetForwarder.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_IAssetForwarder *IAssetForwarderFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *IAssetForwarderFundsDeposited) (event.Subscription, error) {

	logs, sub, err := _IAssetForwarder.contract.WatchLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetForwarderFundsDeposited)
				if err := _IAssetForwarder.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
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

// ParseFundsDeposited is a log parse operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_IAssetForwarder *IAssetForwarderFilterer) ParseFundsDeposited(log types.Log) (*IAssetForwarderFundsDeposited, error) {
	event := new(IAssetForwarderFundsDeposited)
	if err := _IAssetForwarder.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetForwarderFundsDepositedWithMessageIterator is returned from FilterFundsDepositedWithMessage and is used to iterate over the raw logs and unpacked data for FundsDepositedWithMessage events raised by the IAssetForwarder contract.
type IAssetForwarderFundsDepositedWithMessageIterator struct {
	Event *IAssetForwarderFundsDepositedWithMessage // Event containing the contract specifics and raw log

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
func (it *IAssetForwarderFundsDepositedWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetForwarderFundsDepositedWithMessage)
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
		it.Event = new(IAssetForwarderFundsDepositedWithMessage)
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
func (it *IAssetForwarderFundsDepositedWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetForwarderFundsDepositedWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetForwarderFundsDepositedWithMessage represents a FundsDepositedWithMessage event raised by the IAssetForwarder contract.
type IAssetForwarderFundsDepositedWithMessage struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	DestAmount       *big.Int
	DepositId        *big.Int
	SrcToken         common.Address
	Recipient        []byte
	Depositor        common.Address
	DestToken        []byte
	Message          []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDepositedWithMessage is a free log retrieval operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_IAssetForwarder *IAssetForwarderFilterer) FilterFundsDepositedWithMessage(opts *bind.FilterOpts) (*IAssetForwarderFundsDepositedWithMessageIterator, error) {

	logs, sub, err := _IAssetForwarder.contract.FilterLogs(opts, "FundsDepositedWithMessage")
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderFundsDepositedWithMessageIterator{contract: _IAssetForwarder.contract, event: "FundsDepositedWithMessage", logs: logs, sub: sub}, nil
}

// WatchFundsDepositedWithMessage is a free log subscription operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_IAssetForwarder *IAssetForwarderFilterer) WatchFundsDepositedWithMessage(opts *bind.WatchOpts, sink chan<- *IAssetForwarderFundsDepositedWithMessage) (event.Subscription, error) {

	logs, sub, err := _IAssetForwarder.contract.WatchLogs(opts, "FundsDepositedWithMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetForwarderFundsDepositedWithMessage)
				if err := _IAssetForwarder.contract.UnpackLog(event, "FundsDepositedWithMessage", log); err != nil {
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

// ParseFundsDepositedWithMessage is a log parse operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_IAssetForwarder *IAssetForwarderFilterer) ParseFundsDepositedWithMessage(log types.Log) (*IAssetForwarderFundsDepositedWithMessage, error) {
	event := new(IAssetForwarderFundsDepositedWithMessage)
	if err := _IAssetForwarder.contract.UnpackLog(event, "FundsDepositedWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetForwarderFundsPaidIterator is returned from FilterFundsPaid and is used to iterate over the raw logs and unpacked data for FundsPaid events raised by the IAssetForwarder contract.
type IAssetForwarderFundsPaidIterator struct {
	Event *IAssetForwarderFundsPaid // Event containing the contract specifics and raw log

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
func (it *IAssetForwarderFundsPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetForwarderFundsPaid)
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
		it.Event = new(IAssetForwarderFundsPaid)
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
func (it *IAssetForwarderFundsPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetForwarderFundsPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetForwarderFundsPaid represents a FundsPaid event raised by the IAssetForwarder contract.
type IAssetForwarderFundsPaid struct {
	MessageHash [32]byte
	Forwarder   common.Address
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaid is a free log retrieval operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_IAssetForwarder *IAssetForwarderFilterer) FilterFundsPaid(opts *bind.FilterOpts) (*IAssetForwarderFundsPaidIterator, error) {

	logs, sub, err := _IAssetForwarder.contract.FilterLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderFundsPaidIterator{contract: _IAssetForwarder.contract, event: "FundsPaid", logs: logs, sub: sub}, nil
}

// WatchFundsPaid is a free log subscription operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_IAssetForwarder *IAssetForwarderFilterer) WatchFundsPaid(opts *bind.WatchOpts, sink chan<- *IAssetForwarderFundsPaid) (event.Subscription, error) {

	logs, sub, err := _IAssetForwarder.contract.WatchLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetForwarderFundsPaid)
				if err := _IAssetForwarder.contract.UnpackLog(event, "FundsPaid", log); err != nil {
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

// ParseFundsPaid is a log parse operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_IAssetForwarder *IAssetForwarderFilterer) ParseFundsPaid(log types.Log) (*IAssetForwarderFundsPaid, error) {
	event := new(IAssetForwarderFundsPaid)
	if err := _IAssetForwarder.contract.UnpackLog(event, "FundsPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetForwarderFundsPaidWithMessageIterator is returned from FilterFundsPaidWithMessage and is used to iterate over the raw logs and unpacked data for FundsPaidWithMessage events raised by the IAssetForwarder contract.
type IAssetForwarderFundsPaidWithMessageIterator struct {
	Event *IAssetForwarderFundsPaidWithMessage // Event containing the contract specifics and raw log

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
func (it *IAssetForwarderFundsPaidWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetForwarderFundsPaidWithMessage)
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
		it.Event = new(IAssetForwarderFundsPaidWithMessage)
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
func (it *IAssetForwarderFundsPaidWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetForwarderFundsPaidWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetForwarderFundsPaidWithMessage represents a FundsPaidWithMessage event raised by the IAssetForwarder contract.
type IAssetForwarderFundsPaidWithMessage struct {
	MessageHash [32]byte
	Forwarder   common.Address
	Nonce       *big.Int
	ExecFlag    bool
	ExecData    []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaidWithMessage is a free log retrieval operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_IAssetForwarder *IAssetForwarderFilterer) FilterFundsPaidWithMessage(opts *bind.FilterOpts) (*IAssetForwarderFundsPaidWithMessageIterator, error) {

	logs, sub, err := _IAssetForwarder.contract.FilterLogs(opts, "FundsPaidWithMessage")
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderFundsPaidWithMessageIterator{contract: _IAssetForwarder.contract, event: "FundsPaidWithMessage", logs: logs, sub: sub}, nil
}

// WatchFundsPaidWithMessage is a free log subscription operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_IAssetForwarder *IAssetForwarderFilterer) WatchFundsPaidWithMessage(opts *bind.WatchOpts, sink chan<- *IAssetForwarderFundsPaidWithMessage) (event.Subscription, error) {

	logs, sub, err := _IAssetForwarder.contract.WatchLogs(opts, "FundsPaidWithMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetForwarderFundsPaidWithMessage)
				if err := _IAssetForwarder.contract.UnpackLog(event, "FundsPaidWithMessage", log); err != nil {
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

// ParseFundsPaidWithMessage is a log parse operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_IAssetForwarder *IAssetForwarderFilterer) ParseFundsPaidWithMessage(log types.Log) (*IAssetForwarderFundsPaidWithMessage, error) {
	event := new(IAssetForwarderFundsPaidWithMessage)
	if err := _IAssetForwarder.contract.UnpackLog(event, "FundsPaidWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetForwarderIUSDCDepositedIterator is returned from FilterIUSDCDeposited and is used to iterate over the raw logs and unpacked data for IUSDCDeposited events raised by the IAssetForwarder contract.
type IAssetForwarderIUSDCDepositedIterator struct {
	Event *IAssetForwarderIUSDCDeposited // Event containing the contract specifics and raw log

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
func (it *IAssetForwarderIUSDCDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetForwarderIUSDCDeposited)
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
		it.Event = new(IAssetForwarderIUSDCDeposited)
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
func (it *IAssetForwarderIUSDCDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetForwarderIUSDCDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetForwarderIUSDCDeposited represents a IUSDCDeposited event raised by the IAssetForwarder contract.
type IAssetForwarderIUSDCDeposited struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	UsdcNonce        *big.Int
	SrcToken         common.Address
	Recipient        [32]byte
	Depositor        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIUSDCDeposited is a free log retrieval operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_IAssetForwarder *IAssetForwarderFilterer) FilterIUSDCDeposited(opts *bind.FilterOpts) (*IAssetForwarderIUSDCDepositedIterator, error) {

	logs, sub, err := _IAssetForwarder.contract.FilterLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return &IAssetForwarderIUSDCDepositedIterator{contract: _IAssetForwarder.contract, event: "iUSDCDeposited", logs: logs, sub: sub}, nil
}

// WatchIUSDCDeposited is a free log subscription operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_IAssetForwarder *IAssetForwarderFilterer) WatchIUSDCDeposited(opts *bind.WatchOpts, sink chan<- *IAssetForwarderIUSDCDeposited) (event.Subscription, error) {

	logs, sub, err := _IAssetForwarder.contract.WatchLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetForwarderIUSDCDeposited)
				if err := _IAssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
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

// ParseIUSDCDeposited is a log parse operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_IAssetForwarder *IAssetForwarderFilterer) ParseIUSDCDeposited(log types.Log) (*IAssetForwarderIUSDCDeposited, error) {
	event := new(IAssetForwarderIUSDCDeposited)
	if err := _IAssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
