// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ForwarderPool

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

// ForwarderPoolMetaData contains all meta data concerning the ForwarderPool contract.
var ForwarderPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveAssetForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetForwarder\",\"outputs\":[{\"internalType\":\"contractIAssetForwarder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage\",\"name\":\"relayMessageData\",\"type\":\"tuple\"}],\"name\":\"iRelayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetForwarder\",\"type\":\"address\"}],\"name\":\"setAssetForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setWhitelistedFiller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedFillers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610e39806100a56000396000f3fe6080604052600436106100dd5760003560e01c8063715018a61161007f5780638da5cb5b116100595780638da5cb5b1461023e57806397feb9261461025c5780639b5efa8d1461027c578063f2fde38b1461029c57600080fd5b8063715018a6146101e957806379433d8b146101fe5780637fe689171461020657600080fd5b806344004cc1116100bb57806344004cc114610169578063536c6bfa1461018957806364778c1f146101a95780636fb003da146101c957600080fd5b806318137a41146100e257806323d5dc3614610104578063373292a714610149575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610966565b6102bc565b005b34801561011057600080fd5b5061013461011f366004610966565b60036020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561015557600080fd5b5061010261016436600461098a565b6102e6565b34801561017557600080fd5b506101026101843660046109b6565b61039c565b34801561019557600080fd5b506101026101a436600461098a565b61041d565b3480156101b557600080fd5b506101026101c4366004610a67565b61045b565b3480156101d557600080fd5b506101026101e4366004610ae6565b6105ac565b3480156101f557600080fd5b50610102610783565b610102610797565b34801561021257600080fd5b50600454610226906001600160a01b031681565b6040516001600160a01b039091168152602001610140565b34801561024a57600080fd5b506000546001600160a01b0316610226565b34801561026857600080fd5b5061010261027736600461098a565b6107bd565b34801561028857600080fd5b50610102610297366004610bf9565b610866565b3480156102a857600080fd5b506101026102b7366004610966565b610899565b6102c46108d4565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b3360009081526003602052604090205460ff1661031e5760405162461bcd60e51b815260040161031590610c32565b60405180910390fd5b6004805460405163095ea7b360e01b81526001600160a01b03918216928101929092526024820183905283169063095ea7b3906044016020604051808303816000875af1158015610373573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103979190610c81565b505050565b6103a46108d4565b60405163a9059cbb60e01b81526001600160a01b0383811660048301526024820183905284169063a9059cbb906044016020604051808303816000875af11580156103f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104179190610c81565b50505050565b6104256108d4565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610397573d6000803e3d6000fd5b3360009081526003602052604090205460ff1661048a5760405162461bcd60e51b815260040161031590610c32565b6004546001600160a01b03166104b25760405162461bcd60e51b815260040161031590610c9e565b6104dc81606001516001600160a01b031673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15610547576004805482516040516364778c1f60e01b81526001600160a01b03909216926364778c1f9261051291869101610ce3565b6000604051808303818588803b15801561052b57600080fd5b505af115801561053f573d6000803e3d6000fd5b505050505050565b600480546040516364778c1f60e01b81526001600160a01b03909116916364778c1f9161057691859101610ce3565b600060405180830381600087803b15801561059057600080fd5b505af11580156105a4573d6000803e3d6000fd5b505050505b50565b3360009081526003602052604090205460ff166105db5760405162461bcd60e51b815260040161031590610c32565b6004546001600160a01b03166106035760405162461bcd60e51b815260040161031590610c9e565b61062d81606001516001600160a01b031673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b156106a157805147101561069c5760405162461bcd60e51b815260206004820152603060248201527f466f72776172646572506f6f6c3a20696e73756666696369656e74206e61746960448201526f766520746f6b656e2062616c616e636560801b6064820152608401610315565b610754565b805160608201516001600160a01b031660009081526001602052604090205410156107205760405162461bcd60e51b815260206004820152602960248201527f466f72776172646572506f6f6c3a20696e73756666696369656e7420455243326044820152680c0819195c1bdcda5d60ba1b6064820152608401610315565b805160608201516001600160a01b03166000908152600160205260408120805490919061074e908490610d3d565b90915550505b600480546040516337d801ed60e11b81526001600160a01b0390911691636fb003da9161057691859101610d56565b61078b6108d4565b6107956000610901565b565b33600090815260026020526040812080543492906107b6908490610df0565b9091555050565b6040516323b872dd60e01b8152336004820152306024820152604481018290526001600160a01b038316906323b872dd906064016020604051808303816000875af1158015610810573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108349190610c81565b506001600160a01b0382166000908152600160205260408120805483929061085d908490610df0565b90915550505050565b61086e6108d4565b6001600160a01b03919091166000908152600360205260409020805460ff1916911515919091179055565b6108a16108d4565b6001600160a01b0381166108cb57604051631e4fbdf760e01b815260006004820152602401610315565b6105a981610901565b6000546001600160a01b031633146107955760405163118cdaa760e01b8152336004820152602401610315565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146105a957600080fd5b60006020828403121561097857600080fd5b813561098381610951565b9392505050565b6000806040838503121561099d57600080fd5b82356109a881610951565b946020939093013593505050565b6000806000606084860312156109cb57600080fd5b83356109d681610951565b925060208401356109e681610951565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715610a3057610a306109f7565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610a5f57610a5f6109f7565b604052919050565b600060a08284031215610a7957600080fd5b60405160a0810181811067ffffffffffffffff82111715610a9c57610a9c6109f7565b80604052508235815260208301356020820152604083013560408201526060830135610ac781610951565b60608201526080830135610ada81610951565b60808201529392505050565b60006020808385031215610af957600080fd5b823567ffffffffffffffff80821115610b1157600080fd5b9084019060c08287031215610b2557600080fd5b610b2d610a0d565b823581528383013584820152604083013560408201526060830135610b5181610951565b60608201526080830135610b6481610951565b608082015260a083013582811115610b7b57600080fd5b80840193505086601f840112610b9057600080fd5b823582811115610ba257610ba26109f7565b610bb4601f8201601f19168601610a36565b92508083528785828601011115610bca57600080fd5b808585018685013760009083019094019390935260a0830152509392505050565b80151581146105a957600080fd5b60008060408385031215610c0c57600080fd5b8235610c1781610951565b91506020830135610c2781610beb565b809150509250929050565b6020808252602f908201527f466f72776172646572506f6f6c3a2063616c6c6572206973206e6f742077686960408201526e3a32b634b9ba32b2103334b63632b960891b606082015260800190565b600060208284031215610c9357600080fd5b815161098381610beb565b60208082526025908201527f466f72776172646572506f6f6c3a206173736574466f72776172646572206e6f6040820152641d081cd95d60da1b606082015260800190565b8151815260208083015190820152604080830151908201526060808301516001600160a01b0390811691830191909152608092830151169181019190915260a00190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610d5057610d50610d27565b92915050565b6000602080835283518184015280840151604084015260408401516060840152606084015160018060a01b0380821660808601528060808701511660a0860152505060a084015160c08085015280518060e086015260005b81811015610dcb5782810184015186820161010001528301610dae565b506101009250600083828701015282601f19601f830116860101935050505092915050565b80820180821115610d5057610d50610d2756fea2646970667358221220e3ea7c0c7ce2f36cd5ed281289b75986b38e8cb2209461d0b3f7a23d4e62b32164736f6c63430008140033",
}

// ForwarderPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderPoolMetaData.ABI instead.
var ForwarderPoolABI = ForwarderPoolMetaData.ABI

// ForwarderPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForwarderPoolMetaData.Bin instead.
var ForwarderPoolBin = ForwarderPoolMetaData.Bin

// DeployForwarderPool deploys a new Ethereum contract, binding an instance of ForwarderPool to it.
func DeployForwarderPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ForwarderPool, error) {
	parsed, err := ForwarderPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForwarderPoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForwarderPool{ForwarderPoolCaller: ForwarderPoolCaller{contract: contract}, ForwarderPoolTransactor: ForwarderPoolTransactor{contract: contract}, ForwarderPoolFilterer: ForwarderPoolFilterer{contract: contract}}, nil
}

// ForwarderPool is an auto generated Go binding around an Ethereum contract.
type ForwarderPool struct {
	ForwarderPoolCaller     // Read-only binding to the contract
	ForwarderPoolTransactor // Write-only binding to the contract
	ForwarderPoolFilterer   // Log filterer for contract events
}

// ForwarderPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderPoolSession struct {
	Contract     *ForwarderPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderPoolCallerSession struct {
	Contract *ForwarderPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ForwarderPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderPoolTransactorSession struct {
	Contract     *ForwarderPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ForwarderPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderPoolRaw struct {
	Contract *ForwarderPool // Generic contract binding to access the raw methods on
}

// ForwarderPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderPoolCallerRaw struct {
	Contract *ForwarderPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderPoolTransactorRaw struct {
	Contract *ForwarderPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarderPool creates a new instance of ForwarderPool, bound to a specific deployed contract.
func NewForwarderPool(address common.Address, backend bind.ContractBackend) (*ForwarderPool, error) {
	contract, err := bindForwarderPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForwarderPool{ForwarderPoolCaller: ForwarderPoolCaller{contract: contract}, ForwarderPoolTransactor: ForwarderPoolTransactor{contract: contract}, ForwarderPoolFilterer: ForwarderPoolFilterer{contract: contract}}, nil
}

// NewForwarderPoolCaller creates a new read-only instance of ForwarderPool, bound to a specific deployed contract.
func NewForwarderPoolCaller(address common.Address, caller bind.ContractCaller) (*ForwarderPoolCaller, error) {
	contract, err := bindForwarderPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderPoolCaller{contract: contract}, nil
}

// NewForwarderPoolTransactor creates a new write-only instance of ForwarderPool, bound to a specific deployed contract.
func NewForwarderPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderPoolTransactor, error) {
	contract, err := bindForwarderPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderPoolTransactor{contract: contract}, nil
}

// NewForwarderPoolFilterer creates a new log filterer instance of ForwarderPool, bound to a specific deployed contract.
func NewForwarderPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderPoolFilterer, error) {
	contract, err := bindForwarderPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderPoolFilterer{contract: contract}, nil
}

// bindForwarderPool binds a generic wrapper to an already deployed contract.
func bindForwarderPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ForwarderPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForwarderPool *ForwarderPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForwarderPool.Contract.ForwarderPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForwarderPool *ForwarderPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderPool.Contract.ForwarderPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForwarderPool *ForwarderPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForwarderPool.Contract.ForwarderPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForwarderPool *ForwarderPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForwarderPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForwarderPool *ForwarderPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForwarderPool *ForwarderPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForwarderPool.Contract.contract.Transact(opts, method, params...)
}

// AssetForwarder is a free data retrieval call binding the contract method 0x7fe68917.
//
// Solidity: function assetForwarder() view returns(address)
func (_ForwarderPool *ForwarderPoolCaller) AssetForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ForwarderPool.contract.Call(opts, &out, "assetForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetForwarder is a free data retrieval call binding the contract method 0x7fe68917.
//
// Solidity: function assetForwarder() view returns(address)
func (_ForwarderPool *ForwarderPoolSession) AssetForwarder() (common.Address, error) {
	return _ForwarderPool.Contract.AssetForwarder(&_ForwarderPool.CallOpts)
}

// AssetForwarder is a free data retrieval call binding the contract method 0x7fe68917.
//
// Solidity: function assetForwarder() view returns(address)
func (_ForwarderPool *ForwarderPoolCallerSession) AssetForwarder() (common.Address, error) {
	return _ForwarderPool.Contract.AssetForwarder(&_ForwarderPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForwarderPool *ForwarderPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ForwarderPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForwarderPool *ForwarderPoolSession) Owner() (common.Address, error) {
	return _ForwarderPool.Contract.Owner(&_ForwarderPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForwarderPool *ForwarderPoolCallerSession) Owner() (common.Address, error) {
	return _ForwarderPool.Contract.Owner(&_ForwarderPool.CallOpts)
}

// WhitelistedFillers is a free data retrieval call binding the contract method 0x23d5dc36.
//
// Solidity: function whitelistedFillers(address ) view returns(bool)
func (_ForwarderPool *ForwarderPoolCaller) WhitelistedFillers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ForwarderPool.contract.Call(opts, &out, "whitelistedFillers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedFillers is a free data retrieval call binding the contract method 0x23d5dc36.
//
// Solidity: function whitelistedFillers(address ) view returns(bool)
func (_ForwarderPool *ForwarderPoolSession) WhitelistedFillers(arg0 common.Address) (bool, error) {
	return _ForwarderPool.Contract.WhitelistedFillers(&_ForwarderPool.CallOpts, arg0)
}

// WhitelistedFillers is a free data retrieval call binding the contract method 0x23d5dc36.
//
// Solidity: function whitelistedFillers(address ) view returns(bool)
func (_ForwarderPool *ForwarderPoolCallerSession) WhitelistedFillers(arg0 common.Address) (bool, error) {
	return _ForwarderPool.Contract.WhitelistedFillers(&_ForwarderPool.CallOpts, arg0)
}

// ApproveAssetForwarder is a paid mutator transaction binding the contract method 0x373292a7.
//
// Solidity: function approveAssetForwarder(address token, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactor) ApproveAssetForwarder(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "approveAssetForwarder", token, amount)
}

// ApproveAssetForwarder is a paid mutator transaction binding the contract method 0x373292a7.
//
// Solidity: function approveAssetForwarder(address token, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolSession) ApproveAssetForwarder(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.ApproveAssetForwarder(&_ForwarderPool.TransactOpts, token, amount)
}

// ApproveAssetForwarder is a paid mutator transaction binding the contract method 0x373292a7.
//
// Solidity: function approveAssetForwarder(address token, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) ApproveAssetForwarder(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.ApproveAssetForwarder(&_ForwarderPool.TransactOpts, token, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address token, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "depositERC20", token, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address token, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolSession) DepositERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.DepositERC20(&_ForwarderPool.TransactOpts, token, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address token, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) DepositERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.DepositERC20(&_ForwarderPool.TransactOpts, token, amount)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_ForwarderPool *ForwarderPoolTransactor) DepositNativeToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "depositNativeToken")
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_ForwarderPool *ForwarderPoolSession) DepositNativeToken() (*types.Transaction, error) {
	return _ForwarderPool.Contract.DepositNativeToken(&_ForwarderPool.TransactOpts)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) DepositNativeToken() (*types.Transaction, error) {
	return _ForwarderPool.Contract.DepositNativeToken(&_ForwarderPool.TransactOpts)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) returns()
func (_ForwarderPool *ForwarderPoolTransactor) IRelay(opts *bind.TransactOpts, relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "iRelay", relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) returns()
func (_ForwarderPool *ForwarderPoolSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _ForwarderPool.Contract.IRelay(&_ForwarderPool.TransactOpts, relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _ForwarderPool.Contract.IRelay(&_ForwarderPool.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayMessageData) returns()
func (_ForwarderPool *ForwarderPoolTransactor) IRelayMessage(opts *bind.TransactOpts, relayMessageData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "iRelayMessage", relayMessageData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayMessageData) returns()
func (_ForwarderPool *ForwarderPoolSession) IRelayMessage(relayMessageData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _ForwarderPool.Contract.IRelayMessage(&_ForwarderPool.TransactOpts, relayMessageData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayMessageData) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) IRelayMessage(relayMessageData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _ForwarderPool.Contract.IRelayMessage(&_ForwarderPool.TransactOpts, relayMessageData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForwarderPool *ForwarderPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForwarderPool *ForwarderPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _ForwarderPool.Contract.RenounceOwnership(&_ForwarderPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ForwarderPool.Contract.RenounceOwnership(&_ForwarderPool.TransactOpts)
}

// SetAssetForwarder is a paid mutator transaction binding the contract method 0x18137a41.
//
// Solidity: function setAssetForwarder(address _assetForwarder) returns()
func (_ForwarderPool *ForwarderPoolTransactor) SetAssetForwarder(opts *bind.TransactOpts, _assetForwarder common.Address) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "setAssetForwarder", _assetForwarder)
}

// SetAssetForwarder is a paid mutator transaction binding the contract method 0x18137a41.
//
// Solidity: function setAssetForwarder(address _assetForwarder) returns()
func (_ForwarderPool *ForwarderPoolSession) SetAssetForwarder(_assetForwarder common.Address) (*types.Transaction, error) {
	return _ForwarderPool.Contract.SetAssetForwarder(&_ForwarderPool.TransactOpts, _assetForwarder)
}

// SetAssetForwarder is a paid mutator transaction binding the contract method 0x18137a41.
//
// Solidity: function setAssetForwarder(address _assetForwarder) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) SetAssetForwarder(_assetForwarder common.Address) (*types.Transaction, error) {
	return _ForwarderPool.Contract.SetAssetForwarder(&_ForwarderPool.TransactOpts, _assetForwarder)
}

// SetWhitelistedFiller is a paid mutator transaction binding the contract method 0x9b5efa8d.
//
// Solidity: function setWhitelistedFiller(address filler, bool whitelisted) returns()
func (_ForwarderPool *ForwarderPoolTransactor) SetWhitelistedFiller(opts *bind.TransactOpts, filler common.Address, whitelisted bool) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "setWhitelistedFiller", filler, whitelisted)
}

// SetWhitelistedFiller is a paid mutator transaction binding the contract method 0x9b5efa8d.
//
// Solidity: function setWhitelistedFiller(address filler, bool whitelisted) returns()
func (_ForwarderPool *ForwarderPoolSession) SetWhitelistedFiller(filler common.Address, whitelisted bool) (*types.Transaction, error) {
	return _ForwarderPool.Contract.SetWhitelistedFiller(&_ForwarderPool.TransactOpts, filler, whitelisted)
}

// SetWhitelistedFiller is a paid mutator transaction binding the contract method 0x9b5efa8d.
//
// Solidity: function setWhitelistedFiller(address filler, bool whitelisted) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) SetWhitelistedFiller(filler common.Address, whitelisted bool) (*types.Transaction, error) {
	return _ForwarderPool.Contract.SetWhitelistedFiller(&_ForwarderPool.TransactOpts, filler, whitelisted)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForwarderPool *ForwarderPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForwarderPool *ForwarderPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForwarderPool.Contract.TransferOwnership(&_ForwarderPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForwarderPool.Contract.TransferOwnership(&_ForwarderPool.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address recipient, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "withdrawERC20", token, recipient, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address recipient, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolSession) WithdrawERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.WithdrawERC20(&_ForwarderPool.TransactOpts, token, recipient, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address recipient, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) WithdrawERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.WithdrawERC20(&_ForwarderPool.TransactOpts, token, recipient, amount)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x536c6bfa.
//
// Solidity: function withdrawNativeToken(address recipient, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactor) WithdrawNativeToken(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.contract.Transact(opts, "withdrawNativeToken", recipient, amount)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x536c6bfa.
//
// Solidity: function withdrawNativeToken(address recipient, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolSession) WithdrawNativeToken(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.WithdrawNativeToken(&_ForwarderPool.TransactOpts, recipient, amount)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x536c6bfa.
//
// Solidity: function withdrawNativeToken(address recipient, uint256 amount) returns()
func (_ForwarderPool *ForwarderPoolTransactorSession) WithdrawNativeToken(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderPool.Contract.WithdrawNativeToken(&_ForwarderPool.TransactOpts, recipient, amount)
}

// ForwarderPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ForwarderPool contract.
type ForwarderPoolOwnershipTransferredIterator struct {
	Event *ForwarderPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ForwarderPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderPoolOwnershipTransferred)
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
		it.Event = new(ForwarderPoolOwnershipTransferred)
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
func (it *ForwarderPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderPoolOwnershipTransferred represents a OwnershipTransferred event raised by the ForwarderPool contract.
type ForwarderPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForwarderPool *ForwarderPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ForwarderPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForwarderPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderPoolOwnershipTransferredIterator{contract: _ForwarderPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForwarderPool *ForwarderPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForwarderPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForwarderPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderPoolOwnershipTransferred)
				if err := _ForwarderPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ForwarderPool *ForwarderPoolFilterer) ParseOwnershipTransferred(log types.Log) (*ForwarderPoolOwnershipTransferred, error) {
	event := new(ForwarderPoolOwnershipTransferred)
	if err := _ForwarderPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
