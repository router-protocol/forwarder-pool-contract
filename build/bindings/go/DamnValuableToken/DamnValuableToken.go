// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DamnValuableToken

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

// DamnValuableTokenMetaData contains all meta data concerning the DamnValuableToken contract.
var DamnValuableTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051806040016040528060118152602001702230b6b72b30b63ab0b13632aa37b5b2b760791b8152506040518060400160405280600381526020016211159560ea1b8152508160039081620000699190620002af565b506004620000788282620002af565b5050506200008f336000196200009560201b60201c565b620003a3565b6001600160a01b038216620000c55760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b620000d360008383620000d7565b5050565b6001600160a01b03831662000106578060026000828254620000fa91906200037b565b909155506200017a9050565b6001600160a01b038316600090815260208190526040902054818110156200015b5760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000bc565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166200019857600280548290039055620001b7565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620001fd91815260200190565b60405180910390a3505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200023557607f821691505b6020821081036200025657634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002aa57600081815260208120601f850160051c81016020861015620002855750805b601f850160051c820191505b81811015620002a65782815560010162000291565b5050505b505050565b81516001600160401b03811115620002cb57620002cb6200020a565b620002e381620002dc845462000220565b846200025c565b602080601f8311600181146200031b5760008415620003025750858301515b600019600386901b1c1916600185901b178555620002a6565b600085815260208120601f198616915b828110156200034c578886015182559484019460019091019084016200032b565b50858210156200036b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200039d57634e487b7160e01b600052601160045260246000fd5b92915050565b61072080620003b36000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063313ce56711610066578063313ce567146100fe57806370a082311461010d57806395d89b4114610136578063a9059cbb1461013e578063dd62ed3e1461015157600080fd5b806306fdde0314610098578063095ea7b3146100b657806318160ddd146100d957806323b872dd146100eb575b600080fd5b6100a061018a565b6040516100ad919061056a565b60405180910390f35b6100c96100c43660046105d4565b61021c565b60405190151581526020016100ad565b6002545b6040519081526020016100ad565b6100c96100f93660046105fe565b610236565b604051601281526020016100ad565b6100dd61011b36600461063a565b6001600160a01b031660009081526020819052604090205490565b6100a061025a565b6100c961014c3660046105d4565b610269565b6100dd61015f36600461065c565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101999061068f565b80601f01602080910402602001604051908101604052809291908181526020018280546101c59061068f565b80156102125780601f106101e757610100808354040283529160200191610212565b820191906000526020600020905b8154815290600101906020018083116101f557829003601f168201915b5050505050905090565b60003361022a818585610277565b60019150505b92915050565b600033610244858285610289565b61024f85858561030c565b506001949350505050565b6060600480546101999061068f565b60003361022a81858561030c565b610284838383600161036b565b505050565b6001600160a01b03838116600090815260016020908152604080832093861683529290522054600019811461030657818110156102f757604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b6103068484848403600061036b565b50505050565b6001600160a01b03831661033657604051634b637e8f60e11b8152600060048201526024016102ee565b6001600160a01b0382166103605760405163ec442f0560e01b8152600060048201526024016102ee565b610284838383610440565b6001600160a01b0384166103955760405163e602df0560e01b8152600060048201526024016102ee565b6001600160a01b0383166103bf57604051634a1406b160e11b8152600060048201526024016102ee565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561030657826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161043291815260200190565b60405180910390a350505050565b6001600160a01b03831661046b57806002600082825461046091906106c9565b909155506104dd9050565b6001600160a01b038316600090815260208190526040902054818110156104be5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016102ee565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166104f957600280548290039055610518565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161055d91815260200190565b60405180910390a3505050565b600060208083528351808285015260005b818110156105975785810183015185820160400152820161057b565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146105cf57600080fd5b919050565b600080604083850312156105e757600080fd5b6105f0836105b8565b946020939093013593505050565b60008060006060848603121561061357600080fd5b61061c846105b8565b925061062a602085016105b8565b9150604084013590509250925092565b60006020828403121561064c57600080fd5b610655826105b8565b9392505050565b6000806040838503121561066f57600080fd5b610678836105b8565b9150610686602084016105b8565b90509250929050565b600181811c908216806106a357607f821691505b6020821081036106c357634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561023057634e487b7160e01b600052601160045260246000fdfea2646970667358221220a2b1ea38d139ca173770080b3eb388fb391e38baf92a75a959acebbf626c852964736f6c63430008140033",
}

// DamnValuableTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use DamnValuableTokenMetaData.ABI instead.
var DamnValuableTokenABI = DamnValuableTokenMetaData.ABI

// DamnValuableTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DamnValuableTokenMetaData.Bin instead.
var DamnValuableTokenBin = DamnValuableTokenMetaData.Bin

// DeployDamnValuableToken deploys a new Ethereum contract, binding an instance of DamnValuableToken to it.
func DeployDamnValuableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DamnValuableToken, error) {
	parsed, err := DamnValuableTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DamnValuableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DamnValuableToken{DamnValuableTokenCaller: DamnValuableTokenCaller{contract: contract}, DamnValuableTokenTransactor: DamnValuableTokenTransactor{contract: contract}, DamnValuableTokenFilterer: DamnValuableTokenFilterer{contract: contract}}, nil
}

// DamnValuableToken is an auto generated Go binding around an Ethereum contract.
type DamnValuableToken struct {
	DamnValuableTokenCaller     // Read-only binding to the contract
	DamnValuableTokenTransactor // Write-only binding to the contract
	DamnValuableTokenFilterer   // Log filterer for contract events
}

// DamnValuableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DamnValuableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DamnValuableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DamnValuableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DamnValuableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DamnValuableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DamnValuableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DamnValuableTokenSession struct {
	Contract     *DamnValuableToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DamnValuableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DamnValuableTokenCallerSession struct {
	Contract *DamnValuableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DamnValuableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DamnValuableTokenTransactorSession struct {
	Contract     *DamnValuableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DamnValuableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DamnValuableTokenRaw struct {
	Contract *DamnValuableToken // Generic contract binding to access the raw methods on
}

// DamnValuableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DamnValuableTokenCallerRaw struct {
	Contract *DamnValuableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// DamnValuableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DamnValuableTokenTransactorRaw struct {
	Contract *DamnValuableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDamnValuableToken creates a new instance of DamnValuableToken, bound to a specific deployed contract.
func NewDamnValuableToken(address common.Address, backend bind.ContractBackend) (*DamnValuableToken, error) {
	contract, err := bindDamnValuableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DamnValuableToken{DamnValuableTokenCaller: DamnValuableTokenCaller{contract: contract}, DamnValuableTokenTransactor: DamnValuableTokenTransactor{contract: contract}, DamnValuableTokenFilterer: DamnValuableTokenFilterer{contract: contract}}, nil
}

// NewDamnValuableTokenCaller creates a new read-only instance of DamnValuableToken, bound to a specific deployed contract.
func NewDamnValuableTokenCaller(address common.Address, caller bind.ContractCaller) (*DamnValuableTokenCaller, error) {
	contract, err := bindDamnValuableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DamnValuableTokenCaller{contract: contract}, nil
}

// NewDamnValuableTokenTransactor creates a new write-only instance of DamnValuableToken, bound to a specific deployed contract.
func NewDamnValuableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DamnValuableTokenTransactor, error) {
	contract, err := bindDamnValuableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DamnValuableTokenTransactor{contract: contract}, nil
}

// NewDamnValuableTokenFilterer creates a new log filterer instance of DamnValuableToken, bound to a specific deployed contract.
func NewDamnValuableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DamnValuableTokenFilterer, error) {
	contract, err := bindDamnValuableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DamnValuableTokenFilterer{contract: contract}, nil
}

// bindDamnValuableToken binds a generic wrapper to an already deployed contract.
func bindDamnValuableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DamnValuableTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DamnValuableToken *DamnValuableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DamnValuableToken.Contract.DamnValuableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DamnValuableToken *DamnValuableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.DamnValuableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DamnValuableToken *DamnValuableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.DamnValuableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DamnValuableToken *DamnValuableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DamnValuableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DamnValuableToken *DamnValuableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DamnValuableToken *DamnValuableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DamnValuableToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DamnValuableToken.Contract.Allowance(&_DamnValuableToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DamnValuableToken.Contract.Allowance(&_DamnValuableToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DamnValuableToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DamnValuableToken.Contract.BalanceOf(&_DamnValuableToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DamnValuableToken.Contract.BalanceOf(&_DamnValuableToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DamnValuableToken *DamnValuableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DamnValuableToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DamnValuableToken *DamnValuableTokenSession) Decimals() (uint8, error) {
	return _DamnValuableToken.Contract.Decimals(&_DamnValuableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DamnValuableToken *DamnValuableTokenCallerSession) Decimals() (uint8, error) {
	return _DamnValuableToken.Contract.Decimals(&_DamnValuableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DamnValuableToken *DamnValuableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DamnValuableToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DamnValuableToken *DamnValuableTokenSession) Name() (string, error) {
	return _DamnValuableToken.Contract.Name(&_DamnValuableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DamnValuableToken *DamnValuableTokenCallerSession) Name() (string, error) {
	return _DamnValuableToken.Contract.Name(&_DamnValuableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DamnValuableToken *DamnValuableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DamnValuableToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DamnValuableToken *DamnValuableTokenSession) Symbol() (string, error) {
	return _DamnValuableToken.Contract.Symbol(&_DamnValuableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DamnValuableToken *DamnValuableTokenCallerSession) Symbol() (string, error) {
	return _DamnValuableToken.Contract.Symbol(&_DamnValuableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DamnValuableToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenSession) TotalSupply() (*big.Int, error) {
	return _DamnValuableToken.Contract.TotalSupply(&_DamnValuableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DamnValuableToken *DamnValuableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DamnValuableToken.Contract.TotalSupply(&_DamnValuableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.Approve(&_DamnValuableToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.Approve(&_DamnValuableToken.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.Transfer(&_DamnValuableToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.Transfer(&_DamnValuableToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.TransferFrom(&_DamnValuableToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DamnValuableToken *DamnValuableTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DamnValuableToken.Contract.TransferFrom(&_DamnValuableToken.TransactOpts, from, to, value)
}

// DamnValuableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DamnValuableToken contract.
type DamnValuableTokenApprovalIterator struct {
	Event *DamnValuableTokenApproval // Event containing the contract specifics and raw log

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
func (it *DamnValuableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DamnValuableTokenApproval)
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
		it.Event = new(DamnValuableTokenApproval)
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
func (it *DamnValuableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DamnValuableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DamnValuableTokenApproval represents a Approval event raised by the DamnValuableToken contract.
type DamnValuableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DamnValuableToken *DamnValuableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DamnValuableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DamnValuableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DamnValuableTokenApprovalIterator{contract: _DamnValuableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DamnValuableToken *DamnValuableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DamnValuableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DamnValuableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DamnValuableTokenApproval)
				if err := _DamnValuableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DamnValuableToken *DamnValuableTokenFilterer) ParseApproval(log types.Log) (*DamnValuableTokenApproval, error) {
	event := new(DamnValuableTokenApproval)
	if err := _DamnValuableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DamnValuableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DamnValuableToken contract.
type DamnValuableTokenTransferIterator struct {
	Event *DamnValuableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *DamnValuableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DamnValuableTokenTransfer)
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
		it.Event = new(DamnValuableTokenTransfer)
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
func (it *DamnValuableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DamnValuableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DamnValuableTokenTransfer represents a Transfer event raised by the DamnValuableToken contract.
type DamnValuableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DamnValuableToken *DamnValuableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DamnValuableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DamnValuableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DamnValuableTokenTransferIterator{contract: _DamnValuableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DamnValuableToken *DamnValuableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DamnValuableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DamnValuableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DamnValuableTokenTransfer)
				if err := _DamnValuableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DamnValuableToken *DamnValuableTokenFilterer) ParseTransfer(log types.Log) (*DamnValuableTokenTransfer, error) {
	event := new(DamnValuableTokenTransfer)
	if err := _DamnValuableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
