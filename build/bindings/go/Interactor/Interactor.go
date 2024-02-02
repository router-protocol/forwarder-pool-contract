// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Interactor

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

// InteractorMetaData contains all meta data concerning the Interactor contract.
var InteractorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"handleMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061035c806100206000396000f3fe60806040526004361061001e5760003560e01c8063d00a2d5f14610023575b600080fd5b61003661003136600461018a565b610038565b005b6000808280602001905181019061004f919061024a565b915091506000632d1e0c028260405160240161006b91906102d7565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050509050600080846001600160a01b0316836040516100b4919061030a565b6000604051808303816000865af19150503d80600081146100f1576040519150601f19603f3d011682016040523d82523d6000602084013e6100f6565b606091505b5050505050505050505050565b6001600160a01b038116811461011857600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561015a5761015a61011b565b604052919050565b600067ffffffffffffffff82111561017c5761017c61011b565b50601f01601f191660200190565b60008060006060848603121561019f57600080fd5b83356101aa81610103565b925060208401359150604084013567ffffffffffffffff8111156101cd57600080fd5b8401601f810186136101de57600080fd5b80356101f16101ec82610162565b610131565b81815287602083850101111561020657600080fd5b816020840160208301376000602083830101528093505050509250925092565b60005b83811015610241578181015183820152602001610229565b50506000910152565b6000806040838503121561025d57600080fd5b825161026881610103565b602084015190925067ffffffffffffffff81111561028557600080fd5b8301601f8101851361029657600080fd5b80516102a46101ec82610162565b8181528660208385010111156102b957600080fd5b6102ca826020830160208601610226565b8093505050509250929050565b60208152600082518060208401526102f6816040850160208701610226565b601f01601f19169190910160400192915050565b6000825161031c818460208701610226565b919091019291505056fea2646970667358221220efccf75543cbbd96e782a5fffab1fdad054d4d6b6855dd3997aa26b737c5e0b564736f6c63430008140033",
}

// InteractorABI is the input ABI used to generate the binding from.
// Deprecated: Use InteractorMetaData.ABI instead.
var InteractorABI = InteractorMetaData.ABI

// InteractorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InteractorMetaData.Bin instead.
var InteractorBin = InteractorMetaData.Bin

// DeployInteractor deploys a new Ethereum contract, binding an instance of Interactor to it.
func DeployInteractor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Interactor, error) {
	parsed, err := InteractorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InteractorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Interactor{InteractorCaller: InteractorCaller{contract: contract}, InteractorTransactor: InteractorTransactor{contract: contract}, InteractorFilterer: InteractorFilterer{contract: contract}}, nil
}

// Interactor is an auto generated Go binding around an Ethereum contract.
type Interactor struct {
	InteractorCaller     // Read-only binding to the contract
	InteractorTransactor // Write-only binding to the contract
	InteractorFilterer   // Log filterer for contract events
}

// InteractorCaller is an auto generated read-only Go binding around an Ethereum contract.
type InteractorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InteractorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InteractorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InteractorSession struct {
	Contract     *Interactor       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InteractorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InteractorCallerSession struct {
	Contract *InteractorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InteractorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InteractorTransactorSession struct {
	Contract     *InteractorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InteractorRaw is an auto generated low-level Go binding around an Ethereum contract.
type InteractorRaw struct {
	Contract *Interactor // Generic contract binding to access the raw methods on
}

// InteractorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InteractorCallerRaw struct {
	Contract *InteractorCaller // Generic read-only contract binding to access the raw methods on
}

// InteractorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InteractorTransactorRaw struct {
	Contract *InteractorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInteractor creates a new instance of Interactor, bound to a specific deployed contract.
func NewInteractor(address common.Address, backend bind.ContractBackend) (*Interactor, error) {
	contract, err := bindInteractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Interactor{InteractorCaller: InteractorCaller{contract: contract}, InteractorTransactor: InteractorTransactor{contract: contract}, InteractorFilterer: InteractorFilterer{contract: contract}}, nil
}

// NewInteractorCaller creates a new read-only instance of Interactor, bound to a specific deployed contract.
func NewInteractorCaller(address common.Address, caller bind.ContractCaller) (*InteractorCaller, error) {
	contract, err := bindInteractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InteractorCaller{contract: contract}, nil
}

// NewInteractorTransactor creates a new write-only instance of Interactor, bound to a specific deployed contract.
func NewInteractorTransactor(address common.Address, transactor bind.ContractTransactor) (*InteractorTransactor, error) {
	contract, err := bindInteractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InteractorTransactor{contract: contract}, nil
}

// NewInteractorFilterer creates a new log filterer instance of Interactor, bound to a specific deployed contract.
func NewInteractorFilterer(address common.Address, filterer bind.ContractFilterer) (*InteractorFilterer, error) {
	contract, err := bindInteractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InteractorFilterer{contract: contract}, nil
}

// bindInteractor binds a generic wrapper to an already deployed contract.
func bindInteractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InteractorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interactor *InteractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interactor.Contract.InteractorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interactor *InteractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interactor.Contract.InteractorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interactor *InteractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interactor.Contract.InteractorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interactor *InteractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interactor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interactor *InteractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interactor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interactor *InteractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interactor.Contract.contract.Transact(opts, method, params...)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) payable returns()
func (_Interactor *InteractorTransactor) HandleMessage(opts *bind.TransactOpts, tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Interactor.contract.Transact(opts, "handleMessage", tokenSent, amount, message)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) payable returns()
func (_Interactor *InteractorSession) HandleMessage(tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Interactor.Contract.HandleMessage(&_Interactor.TransactOpts, tokenSent, amount, message)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) payable returns()
func (_Interactor *InteractorTransactorSession) HandleMessage(tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Interactor.Contract.HandleMessage(&_Interactor.TransactOpts, tokenSent, amount, message)
}
