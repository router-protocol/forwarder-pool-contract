// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vm

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

// VmSafeAccountAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeAccountAccess struct {
	ChainInfo       VmSafeChainInfo
	Kind            uint8
	Account         common.Address
	Accessor        common.Address
	Initialized     bool
	OldBalance      *big.Int
	NewBalance      *big.Int
	DeployedCode    []byte
	Value           *big.Int
	Data            []byte
	Reverted        bool
	StorageAccesses []VmSafeStorageAccess
}

// VmSafeChainInfo is an auto generated low-level Go binding around an user-defined struct.
type VmSafeChainInfo struct {
	ForkId  *big.Int
	ChainId *big.Int
}

// VmSafeDirEntry is an auto generated low-level Go binding around an user-defined struct.
type VmSafeDirEntry struct {
	ErrorMessage string
	Path         string
	Depth        uint64
	IsDir        bool
	IsSymlink    bool
}

// VmSafeEthGetLogs is an auto generated low-level Go binding around an user-defined struct.
type VmSafeEthGetLogs struct {
	Emitter          common.Address
	Topics           [][32]byte
	Data             []byte
	BlockHash        [32]byte
	BlockNumber      uint64
	TransactionHash  [32]byte
	TransactionIndex uint64
	LogIndex         *big.Int
	Removed          bool
}

// VmSafeFfiResult is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFfiResult struct {
	ExitCode int32
	Stdout   []byte
	Stderr   []byte
}

// VmSafeFsMetadata is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFsMetadata struct {
	IsDir     bool
	IsSymlink bool
	Length    *big.Int
	ReadOnly  bool
	Modified  *big.Int
	Accessed  *big.Int
	Created   *big.Int
}

// VmSafeLog is an auto generated low-level Go binding around an user-defined struct.
type VmSafeLog struct {
	Topics  [][32]byte
	Data    []byte
	Emitter common.Address
}

// VmSafeRpc is an auto generated low-level Go binding around an user-defined struct.
type VmSafeRpc struct {
	Key string
	Url string
}

// VmSafeStorageAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeStorageAccess struct {
	Account       common.Address
	Slot          [32]byte
	IsWrite       bool
	PreviousValue [32]byte
	NewValue      [32]byte
	Reverted      bool
}

// VmSafeWallet is an auto generated low-level Go binding around an user-defined struct.
type VmSafeWallet struct {
	Addr       common.Address
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	PrivateKey *big.Int
}

// VmMetaData contains all meta data concerning the Vm contract.
var VmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"accesses\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"readSlots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"writeSlots\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"keyAddr\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"allowCheatcodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"condition\",\"type\":\"bool\"}],\"name\":\"assume\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"char\",\"type\":\"string\"}],\"name\":\"breakpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"char\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"breakpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainId\",\"type\":\"uint256\"}],\"name\":\"chainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearMockedCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"closeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCoinbase\",\"type\":\"address\"}],\"name\":\"coinbase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initCodeHash\",\"type\":\"bytes32\"}],\"name\":\"computeCreate2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"computeCreate2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"computeCreateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"}],\"name\":\"copyFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"copied\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"recursive\",\"type\":\"bool\"}],\"name\":\"createDir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"}],\"name\":\"createFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"createFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"createFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"createSelectFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"createSelectFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"urlOrAlias\",\"type\":\"string\"}],\"name\":\"createSelectFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"walletLabel\",\"type\":\"string\"}],\"name\":\"createWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"createWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"walletLabel\",\"type\":\"string\"}],\"name\":\"createWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"deleteSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deleteSnapshots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"derivationPath\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"language\",\"type\":\"string\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"language\",\"type\":\"string\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mnemonic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"derivationPath\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"deriveKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDifficulty\",\"type\":\"uint256\"}],\"name\":\"difficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pathToStateJson\",\"type\":\"string\"}],\"name\":\"dumpState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envBool\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envBytes\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envBytes32\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envInt\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"defaultValue\",\"type\":\"bytes32[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"defaultValue\",\"type\":\"int256[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"defaultValue\",\"type\":\"bool\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"defaultValue\",\"type\":\"address\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"defaultValue\",\"type\":\"uint256\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"defaultValue\",\"type\":\"bytes[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"defaultValue\",\"type\":\"uint256[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"defaultValue\",\"type\":\"string[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"defaultValue\",\"type\":\"bytes\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"defaultValue\",\"type\":\"bytes32\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"defaultValue\",\"type\":\"int256\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"defaultValue\",\"type\":\"address[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"defaultValue\",\"type\":\"string\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"defaultValue\",\"type\":\"bool[]\"}],\"name\":\"envOr\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envString\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"envUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"delim\",\"type\":\"string\"}],\"name\":\"envUint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"newRuntimeBytecode\",\"type\":\"bytes\"}],\"name\":\"etch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"}],\"name\":\"eth_getLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"transactionIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"removed\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.EthGetLogs[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"expectCallMinGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"expectCallMinGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"checkTopic1\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic2\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic3\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkData\",\"type\":\"bool\"}],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"checkTopic1\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic2\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkTopic3\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"checkData\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"expectEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"revertData\",\"type\":\"bytes4\"}],\"name\":\"expectRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"expectRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"name\":\"expectSafeMemory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"name\":\"expectSafeMemoryCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBasefee\",\"type\":\"uint256\"}],\"name\":\"fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"commandInput\",\"type\":\"string[]\"}],\"name\":\"ffi\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"fsMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"readOnly\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"modified\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"created\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.FsMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"artifactPath\",\"type\":\"string\"}],\"name\":\"getCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"creationBytecode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"artifactPath\",\"type\":\"string\"}],\"name\":\"getDeployedCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"runtimeBytecode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"currentLabel\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"elementSlot\",\"type\":\"bytes32\"}],\"name\":\"getMappingKeyAndParentOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parent\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"mappingSlot\",\"type\":\"bytes32\"}],\"name\":\"getMappingLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"mappingSlot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMappingSlotAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecordedLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"internalType\":\"structVmSafe.Log[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"isDir\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"isFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPersistent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"persistent\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"keyExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"newLabel\",\"type\":\"string\"}],\"name\":\"label\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"load\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pathToAllocsJson\",\"type\":\"string\"}],\"name\":\"loadAllocs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account1\",\"type\":\"address\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account2\",\"type\":\"address\"}],\"name\":\"makePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"mockCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"mockCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"mockCallRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callee\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"mockCallRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"parsedValue\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"parsedValue\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"parsedValue\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parsedValue\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"parsedValue\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"name\":\"parseJson\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"abiEncodedData\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJson\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"abiEncodedData\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBoolArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytes32Array\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonBytesArray\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonIntArray\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonStringArray\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"parseJsonUintArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"name\":\"parseUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"parsedValue\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGasMetering\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"}],\"name\":\"prank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"prank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newPrevrandao\",\"type\":\"bytes32\"}],\"name\":\"prevrandao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectRoot\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readCallers\",\"outputs\":[{\"internalType\":\"enumVmSafe.CallerMode\",\"name\":\"callerMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"maxDepth\",\"type\":\"uint64\"}],\"name\":\"readDir\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.DirEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"maxDepth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"followLinks\",\"type\":\"bool\"}],\"name\":\"readDir\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.DirEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readDir\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"depth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDir\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSymlink\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.DirEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readFile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readFileBinary\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"readLine\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"line\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"linkPath\",\"type\":\"string\"}],\"name\":\"readLink\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"targetPath\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"record\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordLogs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"rememberKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"keyAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"recursive\",\"type\":\"bool\"}],\"name\":\"removeDir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"removeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"resetNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeGasMetering\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"revertTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"revertToAndDelete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokePersistent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newHeight\",\"type\":\"uint256\"}],\"name\":\"roll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"rollFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"params\",\"type\":\"string\"}],\"name\":\"rpc\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rpcAlias\",\"type\":\"string\"}],\"name\":\"rpcUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rpcUrlStructs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structVmSafe.Rpc[]\",\"name\":\"urls\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rpcUrls\",\"outputs\":[{\"internalType\":\"string[2][]\",\"name\":\"urls\",\"type\":\"string[2][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"}],\"name\":\"selectFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"values\",\"type\":\"address[]\"}],\"name\":\"serializeAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"serializeAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"serializeBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"serializeBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"values\",\"type\":\"bytes[]\"}],\"name\":\"serializeBytes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"serializeBytes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"values\",\"type\":\"bytes32[]\"}],\"name\":\"serializeBytes32\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"serializeBytes32\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"serializeInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"}],\"name\":\"serializeInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"serializeJson\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"serializeString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"serializeString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"serializeUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"serializeUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setEnv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newNonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newNonce\",\"type\":\"uint64\"}],\"name\":\"setNonceUnsafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.Wallet\",\"name\":\"wallet\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"signP256\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"skipTest\",\"type\":\"bool\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"sleep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"startBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"name\":\"startBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startMappingRecording\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"startPrank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"}],\"name\":\"startPrank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startStateDiffRecording\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAndReturnStateDiff\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structVmSafe.ChainInfo\",\"name\":\"chainInfo\",\"type\":\"tuple\"},{\"internalType\":\"enumVmSafe.AccountAccessKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accessor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"deployedCode\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"reverted\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isWrite\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"previousValue\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newValue\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"reverted\",\"type\":\"bool\"}],\"internalType\":\"structVmSafe.StorageAccess[]\",\"name\":\"storageAccesses\",\"type\":\"tuple[]\"}],\"internalType\":\"structVmSafe.AccountAccess[]\",\"name\":\"accountAccesses\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopMappingRecording\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopPrank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"toBase64\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toBase64\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"toBase64URL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toBase64URL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"toString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stringifiedValue\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forkId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"transact\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"transact\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"commandInput\",\"type\":\"string[]\"}],\"name\":\"tryFfi\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"exitCode\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"stdout\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stderr\",\"type\":\"bytes\"}],\"internalType\":\"structVmSafe.FfiResult\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasPrice\",\"type\":\"uint256\"}],\"name\":\"txGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unixTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"milliseconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"warp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"writeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"writeFileBinary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueKey\",\"type\":\"string\"}],\"name\":\"writeJson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"writeJson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"writeLine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VmABI is the input ABI used to generate the binding from.
// Deprecated: Use VmMetaData.ABI instead.
var VmABI = VmMetaData.ABI

// Vm is an auto generated Go binding around an Ethereum contract.
type Vm struct {
	VmCaller     // Read-only binding to the contract
	VmTransactor // Write-only binding to the contract
	VmFilterer   // Log filterer for contract events
}

// VmCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmSession struct {
	Contract     *Vm               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmCallerSession struct {
	Contract *VmCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmTransactorSession struct {
	Contract     *VmTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmRaw struct {
	Contract *Vm // Generic contract binding to access the raw methods on
}

// VmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmCallerRaw struct {
	Contract *VmCaller // Generic read-only contract binding to access the raw methods on
}

// VmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmTransactorRaw struct {
	Contract *VmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVm creates a new instance of Vm, bound to a specific deployed contract.
func NewVm(address common.Address, backend bind.ContractBackend) (*Vm, error) {
	contract, err := bindVm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vm{VmCaller: VmCaller{contract: contract}, VmTransactor: VmTransactor{contract: contract}, VmFilterer: VmFilterer{contract: contract}}, nil
}

// NewVmCaller creates a new read-only instance of Vm, bound to a specific deployed contract.
func NewVmCaller(address common.Address, caller bind.ContractCaller) (*VmCaller, error) {
	contract, err := bindVm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmCaller{contract: contract}, nil
}

// NewVmTransactor creates a new write-only instance of Vm, bound to a specific deployed contract.
func NewVmTransactor(address common.Address, transactor bind.ContractTransactor) (*VmTransactor, error) {
	contract, err := bindVm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmTransactor{contract: contract}, nil
}

// NewVmFilterer creates a new log filterer instance of Vm, bound to a specific deployed contract.
func NewVmFilterer(address common.Address, filterer bind.ContractFilterer) (*VmFilterer, error) {
	contract, err := bindVm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmFilterer{contract: contract}, nil
}

// bindVm binds a generic wrapper to an already deployed contract.
func bindVm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vm *VmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vm.Contract.VmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vm *VmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.Contract.VmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vm *VmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vm.Contract.VmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vm *VmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vm *VmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vm *VmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vm.Contract.contract.Transact(opts, method, params...)
}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmCaller) ActiveFork(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "activeFork")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmSession) ActiveFork() (*big.Int, error) {
	return _Vm.Contract.ActiveFork(&_Vm.CallOpts)
}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmCallerSession) ActiveFork() (*big.Int, error) {
	return _Vm.Contract.ActiveFork(&_Vm.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmCaller) Addr(opts *bind.CallOpts, privateKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "addr", privateKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _Vm.Contract.Addr(&_Vm.CallOpts, privateKey)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmCallerSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _Vm.Contract.Addr(&_Vm.CallOpts, privateKey)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmCaller) Assume(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assume", condition)

	if err != nil {
		return err
	}

	return err

}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmSession) Assume(condition bool) error {
	return _Vm.Contract.Assume(&_Vm.CallOpts, condition)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmCallerSession) Assume(condition bool) error {
	return _Vm.Contract.Assume(&_Vm.CallOpts, condition)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmCaller) ComputeCreate2Address(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreate2Address", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address(&_Vm.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address(&_Vm.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmCaller) ComputeCreate2Address0(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreate2Address0", salt, initCodeHash, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address0(&_Vm.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address0(&_Vm.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmCaller) ComputeCreateAddress(opts *bind.CallOpts, deployer common.Address, nonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreateAddress", deployer, nonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _Vm.Contract.ComputeCreateAddress(&_Vm.CallOpts, deployer, nonce)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _Vm.Contract.ComputeCreateAddress(&_Vm.CallOpts, deployer, nonce)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey", mnemonic, derivationPath, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey(&_Vm.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey(&_Vm.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey0(opts *bind.CallOpts, mnemonic string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey0", mnemonic, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey0(&_Vm.CallOpts, mnemonic, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey0(&_Vm.CallOpts, mnemonic, index, language)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey1(opts *bind.CallOpts, mnemonic string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey1", mnemonic, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey1(&_Vm.CallOpts, mnemonic, index)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey1(&_Vm.CallOpts, mnemonic, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey2(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey2", mnemonic, derivationPath, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey2(&_Vm.CallOpts, mnemonic, derivationPath, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey2(&_Vm.CallOpts, mnemonic, derivationPath, index)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmCaller) EnvAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmSession) EnvAddress(name string) (common.Address, error) {
	return _Vm.Contract.EnvAddress(&_Vm.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmCallerSession) EnvAddress(name string) (common.Address, error) {
	return _Vm.Contract.EnvAddress(&_Vm.CallOpts, name)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmCaller) EnvAddress0(opts *bind.CallOpts, name string, delim string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envAddress0", name, delim)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _Vm.Contract.EnvAddress0(&_Vm.CallOpts, name, delim)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmCallerSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _Vm.Contract.EnvAddress0(&_Vm.CallOpts, name, delim)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmCaller) EnvBool(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBool", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmSession) EnvBool(name string) (bool, error) {
	return _Vm.Contract.EnvBool(&_Vm.CallOpts, name)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmCallerSession) EnvBool(name string) (bool, error) {
	return _Vm.Contract.EnvBool(&_Vm.CallOpts, name)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmCaller) EnvBool0(opts *bind.CallOpts, name string, delim string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBool0", name, delim)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _Vm.Contract.EnvBool0(&_Vm.CallOpts, name, delim)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmCallerSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _Vm.Contract.EnvBool0(&_Vm.CallOpts, name, delim)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmCaller) EnvBytes(opts *bind.CallOpts, name string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmSession) EnvBytes(name string) ([]byte, error) {
	return _Vm.Contract.EnvBytes(&_Vm.CallOpts, name)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmCallerSession) EnvBytes(name string) ([]byte, error) {
	return _Vm.Contract.EnvBytes(&_Vm.CallOpts, name)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmCaller) EnvBytes0(opts *bind.CallOpts, name string, delim string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes0", name, delim)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _Vm.Contract.EnvBytes0(&_Vm.CallOpts, name, delim)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmCallerSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _Vm.Contract.EnvBytes0(&_Vm.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmCaller) EnvBytes32(opts *bind.CallOpts, name string, delim string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes32", name, delim)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _Vm.Contract.EnvBytes32(&_Vm.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmCallerSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _Vm.Contract.EnvBytes32(&_Vm.CallOpts, name, delim)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmCaller) EnvBytes320(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes320", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmSession) EnvBytes320(name string) ([32]byte, error) {
	return _Vm.Contract.EnvBytes320(&_Vm.CallOpts, name)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmCallerSession) EnvBytes320(name string) ([32]byte, error) {
	return _Vm.Contract.EnvBytes320(&_Vm.CallOpts, name)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmCaller) EnvInt(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envInt", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvInt(&_Vm.CallOpts, name, delim)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmCallerSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvInt(&_Vm.CallOpts, name, delim)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmCaller) EnvInt0(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envInt0", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmSession) EnvInt0(name string) (*big.Int, error) {
	return _Vm.Contract.EnvInt0(&_Vm.CallOpts, name)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmCallerSession) EnvInt0(name string) (*big.Int, error) {
	return _Vm.Contract.EnvInt0(&_Vm.CallOpts, name)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmCaller) EnvOr(opts *bind.CallOpts, name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr", name, delim, defaultValue)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _Vm.Contract.EnvOr(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmCallerSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _Vm.Contract.EnvOr(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmCaller) EnvOr0(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr0", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr0(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmCallerSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr0(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmCaller) EnvOr1(opts *bind.CallOpts, name string, defaultValue bool) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr1", name, defaultValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _Vm.Contract.EnvOr1(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmCallerSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _Vm.Contract.EnvOr1(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmCaller) EnvOr10(opts *bind.CallOpts, name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr10", name, delim, defaultValue)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _Vm.Contract.EnvOr10(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmCallerSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _Vm.Contract.EnvOr10(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmCaller) EnvOr11(opts *bind.CallOpts, name string, defaultValue string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr11", name, defaultValue)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _Vm.Contract.EnvOr11(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmCallerSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _Vm.Contract.EnvOr11(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmCaller) EnvOr12(opts *bind.CallOpts, name string, delim string, defaultValue []bool) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr12", name, delim, defaultValue)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _Vm.Contract.EnvOr12(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmCallerSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _Vm.Contract.EnvOr12(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmCaller) EnvOr2(opts *bind.CallOpts, name string, defaultValue common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr2", name, defaultValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _Vm.Contract.EnvOr2(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmCallerSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _Vm.Contract.EnvOr2(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmCaller) EnvOr3(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr3", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr3(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmCallerSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr3(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmCaller) EnvOr4(opts *bind.CallOpts, name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr4", name, delim, defaultValue)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _Vm.Contract.EnvOr4(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmCallerSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _Vm.Contract.EnvOr4(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmCaller) EnvOr5(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr5", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr5(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmCallerSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr5(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmCaller) EnvOr6(opts *bind.CallOpts, name string, delim string, defaultValue []string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr6", name, delim, defaultValue)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _Vm.Contract.EnvOr6(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmCallerSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _Vm.Contract.EnvOr6(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmCaller) EnvOr7(opts *bind.CallOpts, name string, defaultValue []byte) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr7", name, defaultValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _Vm.Contract.EnvOr7(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmCallerSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _Vm.Contract.EnvOr7(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmCaller) EnvOr8(opts *bind.CallOpts, name string, defaultValue [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr8", name, defaultValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _Vm.Contract.EnvOr8(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmCallerSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _Vm.Contract.EnvOr8(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmCaller) EnvOr9(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr9", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr9(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmCallerSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr9(&_Vm.CallOpts, name, defaultValue)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmCaller) EnvString(opts *bind.CallOpts, name string, delim string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envString", name, delim)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmSession) EnvString(name string, delim string) ([]string, error) {
	return _Vm.Contract.EnvString(&_Vm.CallOpts, name, delim)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmCallerSession) EnvString(name string, delim string) ([]string, error) {
	return _Vm.Contract.EnvString(&_Vm.CallOpts, name, delim)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmCaller) EnvString0(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envString0", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmSession) EnvString0(name string) (string, error) {
	return _Vm.Contract.EnvString0(&_Vm.CallOpts, name)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmCallerSession) EnvString0(name string) (string, error) {
	return _Vm.Contract.EnvString0(&_Vm.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmCaller) EnvUint(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envUint", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmSession) EnvUint(name string) (*big.Int, error) {
	return _Vm.Contract.EnvUint(&_Vm.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmCallerSession) EnvUint(name string) (*big.Int, error) {
	return _Vm.Contract.EnvUint(&_Vm.CallOpts, name)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmCaller) EnvUint0(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envUint0", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvUint0(&_Vm.CallOpts, name, delim)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmCallerSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvUint0(&_Vm.CallOpts, name, delim)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmCaller) FsMetadata(opts *bind.CallOpts, path string) (VmSafeFsMetadata, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "fsMetadata", path)

	if err != nil {
		return *new(VmSafeFsMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeFsMetadata)).(*VmSafeFsMetadata)

	return out0, err

}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _Vm.Contract.FsMetadata(&_Vm.CallOpts, path)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmCallerSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _Vm.Contract.FsMetadata(&_Vm.CallOpts, path)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmSession) GetBlockNumber() (*big.Int, error) {
	return _Vm.Contract.GetBlockNumber(&_Vm.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Vm.Contract.GetBlockNumber(&_Vm.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmSession) GetBlockTimestamp() (*big.Int, error) {
	return _Vm.Contract.GetBlockTimestamp(&_Vm.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _Vm.Contract.GetBlockTimestamp(&_Vm.CallOpts)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmCaller) GetCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmSession) GetCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetCode(&_Vm.CallOpts, artifactPath)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmCallerSession) GetCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmCaller) GetDeployedCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getDeployedCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetDeployedCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmCallerSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetDeployedCode(&_Vm.CallOpts, artifactPath)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmCaller) GetLabel(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getLabel", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmSession) GetLabel(account common.Address) (string, error) {
	return _Vm.Contract.GetLabel(&_Vm.CallOpts, account)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmCallerSession) GetLabel(account common.Address) (string, error) {
	return _Vm.Contract.GetLabel(&_Vm.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmCaller) GetNonce(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getNonce", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmSession) GetNonce(account common.Address) (uint64, error) {
	return _Vm.Contract.GetNonce(&_Vm.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmCallerSession) GetNonce(account common.Address) (uint64, error) {
	return _Vm.Contract.GetNonce(&_Vm.CallOpts, account)
}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmCaller) IsPersistent(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isPersistent", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmSession) IsPersistent(account common.Address) (bool, error) {
	return _Vm.Contract.IsPersistent(&_Vm.CallOpts, account)
}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmCallerSession) IsPersistent(account common.Address) (bool, error) {
	return _Vm.Contract.IsPersistent(&_Vm.CallOpts, account)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmCaller) KeyExists(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExists", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmSession) KeyExists(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExists(&_Vm.CallOpts, json, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExists(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExists(&_Vm.CallOpts, json, key)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmCaller) Load(opts *bind.CallOpts, target common.Address, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "load", target, slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _Vm.Contract.Load(&_Vm.CallOpts, target, slot)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmCallerSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _Vm.Contract.Load(&_Vm.CallOpts, target, slot)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmCaller) ParseAddress(opts *bind.CallOpts, stringifiedValue string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseAddress", stringifiedValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _Vm.Contract.ParseAddress(&_Vm.CallOpts, stringifiedValue)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmCallerSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _Vm.Contract.ParseAddress(&_Vm.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmCaller) ParseBool(opts *bind.CallOpts, stringifiedValue string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBool", stringifiedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmSession) ParseBool(stringifiedValue string) (bool, error) {
	return _Vm.Contract.ParseBool(&_Vm.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmCallerSession) ParseBool(stringifiedValue string) (bool, error) {
	return _Vm.Contract.ParseBool(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmCaller) ParseBytes(opts *bind.CallOpts, stringifiedValue string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBytes", stringifiedValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _Vm.Contract.ParseBytes(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmCallerSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _Vm.Contract.ParseBytes(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmCaller) ParseBytes32(opts *bind.CallOpts, stringifiedValue string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBytes32", stringifiedValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _Vm.Contract.ParseBytes32(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmCallerSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _Vm.Contract.ParseBytes32(&_Vm.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmCaller) ParseInt(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseInt", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseInt(&_Vm.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmCallerSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseInt(&_Vm.CallOpts, stringifiedValue)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseJson(opts *bind.CallOpts, json string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJson", json)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseJson(json string) ([]byte, error) {
	return _Vm.Contract.ParseJson(&_Vm.CallOpts, json)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseJson(json string) ([]byte, error) {
	return _Vm.Contract.ParseJson(&_Vm.CallOpts, json)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseJson0(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJson0", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseJson0(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJson0(&_Vm.CallOpts, json, key)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseJson0(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJson0(&_Vm.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmCaller) ParseJsonAddress(opts *bind.CallOpts, json string, key string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonAddress", json, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _Vm.Contract.ParseJsonAddress(&_Vm.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmCallerSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _Vm.Contract.ParseJsonAddress(&_Vm.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmCaller) ParseJsonAddressArray(opts *bind.CallOpts, json string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonAddressArray", json, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseJsonAddressArray(&_Vm.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmCallerSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseJsonAddressArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmCaller) ParseJsonBool(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBool", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmSession) ParseJsonBool(json string, key string) (bool, error) {
	return _Vm.Contract.ParseJsonBool(&_Vm.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmCallerSession) ParseJsonBool(json string, key string) (bool, error) {
	return _Vm.Contract.ParseJsonBool(&_Vm.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmCaller) ParseJsonBoolArray(opts *bind.CallOpts, json string, key string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBoolArray", json, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _Vm.Contract.ParseJsonBoolArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmCallerSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _Vm.Contract.ParseJsonBoolArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmCaller) ParseJsonBytes(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJsonBytes(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmCallerSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJsonBytes(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmCaller) ParseJsonBytes32(opts *bind.CallOpts, json string, key string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes32", json, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmCallerSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmCaller) ParseJsonBytes32Array(opts *bind.CallOpts, json string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes32Array", json, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32Array(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmCallerSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32Array(&_Vm.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmCaller) ParseJsonBytesArray(opts *bind.CallOpts, json string, key string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytesArray", json, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseJsonBytesArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmCallerSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseJsonBytesArray(&_Vm.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmCaller) ParseJsonInt(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonInt", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonInt(&_Vm.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmCallerSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonInt(&_Vm.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmCaller) ParseJsonIntArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonIntArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonIntArray(&_Vm.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmCallerSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonIntArray(&_Vm.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmCaller) ParseJsonKeys(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonKeys", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonKeys(&_Vm.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmCallerSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonKeys(&_Vm.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmCaller) ParseJsonString(opts *bind.CallOpts, json string, key string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonString", json, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmSession) ParseJsonString(json string, key string) (string, error) {
	return _Vm.Contract.ParseJsonString(&_Vm.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmCallerSession) ParseJsonString(json string, key string) (string, error) {
	return _Vm.Contract.ParseJsonString(&_Vm.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmCaller) ParseJsonStringArray(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonStringArray", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonStringArray(&_Vm.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmCallerSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonStringArray(&_Vm.CallOpts, json, key)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmCaller) ParseJsonUint(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonUint", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonUint(&_Vm.CallOpts, json, key)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmCallerSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonUint(&_Vm.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmCaller) ParseJsonUintArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonUintArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonUintArray(&_Vm.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmCallerSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonUintArray(&_Vm.CallOpts, json, key)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmCaller) ParseUint(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseUint", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseUint(&_Vm.CallOpts, stringifiedValue)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmCallerSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseUint(&_Vm.CallOpts, stringifiedValue)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmCaller) ProjectRoot(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "projectRoot")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmSession) ProjectRoot() (string, error) {
	return _Vm.Contract.ProjectRoot(&_Vm.CallOpts)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmCallerSession) ProjectRoot() (string, error) {
	return _Vm.Contract.ProjectRoot(&_Vm.CallOpts)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir(opts *bind.CallOpts, path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir", path, maxDepth)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir(&_Vm.CallOpts, path, maxDepth)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir(&_Vm.CallOpts, path, maxDepth)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir0(opts *bind.CallOpts, path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir0", path, maxDepth, followLinks)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir0(&_Vm.CallOpts, path, maxDepth, followLinks)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir0(&_Vm.CallOpts, path, maxDepth, followLinks)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir1(opts *bind.CallOpts, path string) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir1", path)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir1(&_Vm.CallOpts, path)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir1(&_Vm.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmCaller) ReadFile(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readFile", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmSession) ReadFile(path string) (string, error) {
	return _Vm.Contract.ReadFile(&_Vm.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmCallerSession) ReadFile(path string) (string, error) {
	return _Vm.Contract.ReadFile(&_Vm.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmCaller) ReadFileBinary(opts *bind.CallOpts, path string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readFileBinary", path)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmSession) ReadFileBinary(path string) ([]byte, error) {
	return _Vm.Contract.ReadFileBinary(&_Vm.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmCallerSession) ReadFileBinary(path string) ([]byte, error) {
	return _Vm.Contract.ReadFileBinary(&_Vm.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmCaller) ReadLine(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readLine", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmSession) ReadLine(path string) (string, error) {
	return _Vm.Contract.ReadLine(&_Vm.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmCallerSession) ReadLine(path string) (string, error) {
	return _Vm.Contract.ReadLine(&_Vm.CallOpts, path)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmCaller) ReadLink(opts *bind.CallOpts, linkPath string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readLink", linkPath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmSession) ReadLink(linkPath string) (string, error) {
	return _Vm.Contract.ReadLink(&_Vm.CallOpts, linkPath)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmCallerSession) ReadLink(linkPath string) (string, error) {
	return _Vm.Contract.ReadLink(&_Vm.CallOpts, linkPath)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmCaller) RpcUrl(opts *bind.CallOpts, rpcAlias string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrl", rpcAlias)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmSession) RpcUrl(rpcAlias string) (string, error) {
	return _Vm.Contract.RpcUrl(&_Vm.CallOpts, rpcAlias)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmCallerSession) RpcUrl(rpcAlias string) (string, error) {
	return _Vm.Contract.RpcUrl(&_Vm.CallOpts, rpcAlias)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmCaller) RpcUrlStructs(opts *bind.CallOpts) ([]VmSafeRpc, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrlStructs")

	if err != nil {
		return *new([]VmSafeRpc), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeRpc)).(*[]VmSafeRpc)

	return out0, err

}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _Vm.Contract.RpcUrlStructs(&_Vm.CallOpts)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmCallerSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _Vm.Contract.RpcUrlStructs(&_Vm.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmCaller) RpcUrls(opts *bind.CallOpts) ([][2]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrls")

	if err != nil {
		return *new([][2]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]string)).(*[][2]string)

	return out0, err

}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmSession) RpcUrls() ([][2]string, error) {
	return _Vm.Contract.RpcUrls(&_Vm.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmCallerSession) RpcUrls() ([][2]string, error) {
	return _Vm.Contract.RpcUrls(&_Vm.CallOpts)
}

// Sign0 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign0(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign0", privateKey, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign0 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign0(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign0(&_Vm.CallOpts, privateKey, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign0(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign0(&_Vm.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmCaller) SignP256(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "signP256", privateKey, digest)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.SignP256(&_Vm.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.SignP256(&_Vm.CallOpts, privateKey, digest)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmCaller) ToBase64(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmSession) ToBase64(data string) (string, error) {
	return _Vm.Contract.ToBase64(&_Vm.CallOpts, data)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64(data string) (string, error) {
	return _Vm.Contract.ToBase64(&_Vm.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmCaller) ToBase640(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase640", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmSession) ToBase640(data []byte) (string, error) {
	return _Vm.Contract.ToBase640(&_Vm.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmCallerSession) ToBase640(data []byte) (string, error) {
	return _Vm.Contract.ToBase640(&_Vm.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmCaller) ToBase64URL(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64URL", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmSession) ToBase64URL(data string) (string, error) {
	return _Vm.Contract.ToBase64URL(&_Vm.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64URL(data string) (string, error) {
	return _Vm.Contract.ToBase64URL(&_Vm.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmCaller) ToBase64URL0(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64URL0", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmSession) ToBase64URL0(data []byte) (string, error) {
	return _Vm.Contract.ToBase64URL0(&_Vm.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64URL0(data []byte) (string, error) {
	return _Vm.Contract.ToBase64URL0(&_Vm.CallOpts, data)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString(opts *bind.CallOpts, value common.Address) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString(value common.Address) (string, error) {
	return _Vm.Contract.ToString(&_Vm.CallOpts, value)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString(value common.Address) (string, error) {
	return _Vm.Contract.ToString(&_Vm.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString0(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString0", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString0(value *big.Int) (string, error) {
	return _Vm.Contract.ToString0(&_Vm.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString0(value *big.Int) (string, error) {
	return _Vm.Contract.ToString0(&_Vm.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString1(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString1", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString1(value []byte) (string, error) {
	return _Vm.Contract.ToString1(&_Vm.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString1(value []byte) (string, error) {
	return _Vm.Contract.ToString1(&_Vm.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString2(opts *bind.CallOpts, value bool) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString2", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString2(value bool) (string, error) {
	return _Vm.Contract.ToString2(&_Vm.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString2(value bool) (string, error) {
	return _Vm.Contract.ToString2(&_Vm.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString3(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString3", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString3(value *big.Int) (string, error) {
	return _Vm.Contract.ToString3(&_Vm.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString3(value *big.Int) (string, error) {
	return _Vm.Contract.ToString3(&_Vm.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString4(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString4", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString4(value [32]byte) (string, error) {
	return _Vm.Contract.ToString4(&_Vm.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString4(value [32]byte) (string, error) {
	return _Vm.Contract.ToString4(&_Vm.CallOpts, value)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmTransactor) Accesses(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "accesses", target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Accesses(&_Vm.TransactOpts, target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmTransactorSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Accesses(&_Vm.TransactOpts, target)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmTransactor) AllowCheatcodes(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "allowCheatcodes", account)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmSession) AllowCheatcodes(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.AllowCheatcodes(&_Vm.TransactOpts, account)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmTransactorSession) AllowCheatcodes(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.AllowCheatcodes(&_Vm.TransactOpts, account)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_Vm *VmTransactor) Breakpoint(opts *bind.TransactOpts, char string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "breakpoint", char)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_Vm *VmSession) Breakpoint(char string) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint(&_Vm.TransactOpts, char)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_Vm *VmTransactorSession) Breakpoint(char string) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint(&_Vm.TransactOpts, char)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_Vm *VmTransactor) Breakpoint0(opts *bind.TransactOpts, char string, value bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "breakpoint0", char, value)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_Vm *VmSession) Breakpoint0(char string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint0(&_Vm.TransactOpts, char, value)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_Vm *VmTransactorSession) Breakpoint0(char string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.Breakpoint0(&_Vm.TransactOpts, char, value)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmTransactor) Broadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast")
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmSession) Broadcast() (*types.Transaction, error) {
	return _Vm.Contract.Broadcast(&_Vm.TransactOpts)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmTransactorSession) Broadcast() (*types.Transaction, error) {
	return _Vm.Contract.Broadcast(&_Vm.TransactOpts)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmTransactor) Broadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast0", signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast0(&_Vm.TransactOpts, signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmTransactorSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast0(&_Vm.TransactOpts, signer)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmTransactor) Broadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast1", privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast1(&_Vm.TransactOpts, privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmTransactorSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast1(&_Vm.TransactOpts, privateKey)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmTransactor) ChainId(opts *bind.TransactOpts, newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "chainId", newChainId)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmSession) ChainId(newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.ChainId(&_Vm.TransactOpts, newChainId)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmTransactorSession) ChainId(newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.ChainId(&_Vm.TransactOpts, newChainId)
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmTransactor) ClearMockedCalls(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "clearMockedCalls")
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmSession) ClearMockedCalls() (*types.Transaction, error) {
	return _Vm.Contract.ClearMockedCalls(&_Vm.TransactOpts)
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmTransactorSession) ClearMockedCalls() (*types.Transaction, error) {
	return _Vm.Contract.ClearMockedCalls(&_Vm.TransactOpts)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmTransactor) CloseFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "closeFile", path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmSession) CloseFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.CloseFile(&_Vm.TransactOpts, path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmTransactorSession) CloseFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.CloseFile(&_Vm.TransactOpts, path)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmTransactor) Coinbase(opts *bind.TransactOpts, newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "coinbase", newCoinbase)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmSession) Coinbase(newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Coinbase(&_Vm.TransactOpts, newCoinbase)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmTransactorSession) Coinbase(newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Coinbase(&_Vm.TransactOpts, newCoinbase)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmTransactor) CopyFile(opts *bind.TransactOpts, from string, to string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "copyFile", from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _Vm.Contract.CopyFile(&_Vm.TransactOpts, from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmTransactorSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _Vm.Contract.CopyFile(&_Vm.TransactOpts, from, to)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmTransactor) CreateDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createDir", path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.CreateDir(&_Vm.TransactOpts, path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmTransactorSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.CreateDir(&_Vm.TransactOpts, path, recursive)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork(opts *bind.TransactOpts, urlOrAlias string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork", urlOrAlias)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork(&_Vm.TransactOpts, urlOrAlias)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork(&_Vm.TransactOpts, urlOrAlias)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork0(opts *bind.TransactOpts, urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork0", urlOrAlias, blockNumber)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork0(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork0(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork0(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork0(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork1(opts *bind.TransactOpts, urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork1", urlOrAlias, txHash)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork1(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork1(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork1(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork1(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork(opts *bind.TransactOpts, urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork", urlOrAlias, blockNumber)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork0(opts *bind.TransactOpts, urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork0", urlOrAlias, txHash)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork0(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork0(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork0(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork0(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork1(opts *bind.TransactOpts, urlOrAlias string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork1", urlOrAlias)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork1(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork1(&_Vm.TransactOpts, urlOrAlias)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork1(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork1(&_Vm.TransactOpts, urlOrAlias)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet(opts *bind.TransactOpts, walletLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet", walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet(&_Vm.TransactOpts, walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet(&_Vm.TransactOpts, walletLabel)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet0(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet0", privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet0(&_Vm.TransactOpts, privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet0(&_Vm.TransactOpts, privateKey)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet1(opts *bind.TransactOpts, privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet1", privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet1(&_Vm.TransactOpts, privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet1(&_Vm.TransactOpts, privateKey, walletLabel)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmTransactor) Deal(opts *bind.TransactOpts, account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deal", account, newBalance)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmSession) Deal(account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Deal(&_Vm.TransactOpts, account, newBalance)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmTransactorSession) Deal(account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Deal(&_Vm.TransactOpts, account, newBalance)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) DeleteSnapshot(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteSnapshot", snapshotId)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) DeleteSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) DeleteSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmTransactor) DeleteSnapshots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteSnapshots")
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmSession) DeleteSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshots(&_Vm.TransactOpts)
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmTransactorSession) DeleteSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshots(&_Vm.TransactOpts)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmTransactor) Difficulty(opts *bind.TransactOpts, newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "difficulty", newDifficulty)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmSession) Difficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Difficulty(&_Vm.TransactOpts, newDifficulty)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmTransactorSession) Difficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Difficulty(&_Vm.TransactOpts, newDifficulty)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmTransactor) DumpState(opts *bind.TransactOpts, pathToStateJson string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "dumpState", pathToStateJson)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmSession) DumpState(pathToStateJson string) (*types.Transaction, error) {
	return _Vm.Contract.DumpState(&_Vm.TransactOpts, pathToStateJson)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmTransactorSession) DumpState(pathToStateJson string) (*types.Transaction, error) {
	return _Vm.Contract.DumpState(&_Vm.TransactOpts, pathToStateJson)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmTransactor) Etch(opts *bind.TransactOpts, target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "etch", target, newRuntimeBytecode)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmSession) Etch(target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.Contract.Etch(&_Vm.TransactOpts, target, newRuntimeBytecode)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmTransactorSession) Etch(target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.Contract.Etch(&_Vm.TransactOpts, target, newRuntimeBytecode)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmTransactor) EthGetLogs(opts *bind.TransactOpts, fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "eth_getLogs", fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.EthGetLogs(&_Vm.TransactOpts, fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmTransactorSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.EthGetLogs(&_Vm.TransactOpts, fromBlock, toBlock, target, topics)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_Vm *VmTransactor) Exists(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "exists", path)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_Vm *VmSession) Exists(path string) (*types.Transaction, error) {
	return _Vm.Contract.Exists(&_Vm.TransactOpts, path)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_Vm *VmTransactorSession) Exists(path string) (*types.Transaction, error) {
	return _Vm.Contract.Exists(&_Vm.TransactOpts, path)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall", callee, msgValue, gas, data)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmSession) ExpectCall(callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall(&_Vm.TransactOpts, callee, msgValue, gas, data)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall(callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall(&_Vm.TransactOpts, callee, msgValue, gas, data)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall0", callee, msgValue, gas, data, count)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall0(callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall0(&_Vm.TransactOpts, callee, msgValue, gas, data, count)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall0(callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall0(&_Vm.TransactOpts, callee, msgValue, gas, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall1(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall1", callee, msgValue, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall1(callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall1(&_Vm.TransactOpts, callee, msgValue, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall1(callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall1(&_Vm.TransactOpts, callee, msgValue, data, count)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall2(opts *bind.TransactOpts, callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall2", callee, data)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmSession) ExpectCall2(callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall2(&_Vm.TransactOpts, callee, data)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall2(callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall2(&_Vm.TransactOpts, callee, data)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall3(opts *bind.TransactOpts, callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall3", callee, data, count)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall3(callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall3(&_Vm.TransactOpts, callee, data, count)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall3(callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall3(&_Vm.TransactOpts, callee, data, count)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall4(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall4", callee, msgValue, data)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmSession) ExpectCall4(callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall4(&_Vm.TransactOpts, callee, msgValue, data)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall4(callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall4(&_Vm.TransactOpts, callee, msgValue, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmTransactor) ExpectCallMinGas(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCallMinGas", callee, msgValue, minGas, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmSession) ExpectCallMinGas(callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas(&_Vm.TransactOpts, callee, msgValue, minGas, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCallMinGas(callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas(&_Vm.TransactOpts, callee, msgValue, minGas, data)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCallMinGas0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCallMinGas0", callee, msgValue, minGas, data, count)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCallMinGas0(callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas0(&_Vm.TransactOpts, callee, msgValue, minGas, data, count)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCallMinGas0(callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas0(&_Vm.TransactOpts, callee, msgValue, minGas, data, count)
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmTransactor) ExpectEmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit")
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmSession) ExpectEmit() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit(&_Vm.TransactOpts)
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmTransactorSession) ExpectEmit() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit(&_Vm.TransactOpts)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactor) ExpectEmit0(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit0", checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmSession) ExpectEmit0(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit0(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactorSession) ExpectEmit0(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit0(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactor) ExpectEmit1(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit1", checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmSession) ExpectEmit1(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit1(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmit1(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit1(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmTransactor) ExpectEmit2(opts *bind.TransactOpts, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit2", emitter)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmSession) ExpectEmit2(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit2(&_Vm.TransactOpts, emitter)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmit2(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit2(&_Vm.TransactOpts, emitter)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmTransactor) ExpectRevert(opts *bind.TransactOpts, revertData [4]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert", revertData)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmSession) ExpectRevert(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert(&_Vm.TransactOpts, revertData)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmTransactorSession) ExpectRevert(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert(&_Vm.TransactOpts, revertData)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmTransactor) ExpectRevert0(opts *bind.TransactOpts, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert0", revertData)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmSession) ExpectRevert0(revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert0(&_Vm.TransactOpts, revertData)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmTransactorSession) ExpectRevert0(revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert0(&_Vm.TransactOpts, revertData)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmTransactor) ExpectRevert1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert1")
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmSession) ExpectRevert1() (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert1(&_Vm.TransactOpts)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmTransactorSession) ExpectRevert1() (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert1(&_Vm.TransactOpts)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmTransactor) ExpectSafeMemory(opts *bind.TransactOpts, min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectSafeMemory", min, max)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmSession) ExpectSafeMemory(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemory(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmTransactorSession) ExpectSafeMemory(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemory(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmTransactor) ExpectSafeMemoryCall(opts *bind.TransactOpts, min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectSafeMemoryCall", min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmSession) ExpectSafeMemoryCall(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemoryCall(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmTransactorSession) ExpectSafeMemoryCall(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemoryCall(&_Vm.TransactOpts, min, max)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmTransactor) Fee(opts *bind.TransactOpts, newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "fee", newBasefee)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmSession) Fee(newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Fee(&_Vm.TransactOpts, newBasefee)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmTransactorSession) Fee(newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Fee(&_Vm.TransactOpts, newBasefee)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmTransactor) Ffi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "ffi", commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.Ffi(&_Vm.TransactOpts, commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmTransactorSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.Ffi(&_Vm.TransactOpts, commandInput)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmTransactor) GetMappingKeyAndParentOf(opts *bind.TransactOpts, target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingKeyAndParentOf", target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingKeyAndParentOf(&_Vm.TransactOpts, target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmTransactorSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingKeyAndParentOf(&_Vm.TransactOpts, target, elementSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmTransactor) GetMappingLength(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingLength", target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingLength(&_Vm.TransactOpts, target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmTransactorSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingLength(&_Vm.TransactOpts, target, mappingSlot)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmTransactor) GetMappingSlotAt(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingSlotAt", target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingSlotAt(&_Vm.TransactOpts, target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmTransactorSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingSlotAt(&_Vm.TransactOpts, target, mappingSlot, idx)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmTransactor) GetNonce0(opts *bind.TransactOpts, wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getNonce0", wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.Contract.GetNonce0(&_Vm.TransactOpts, wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmTransactorSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.Contract.GetNonce0(&_Vm.TransactOpts, wallet)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmTransactor) GetRecordedLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getRecordedLogs")
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmSession) GetRecordedLogs() (*types.Transaction, error) {
	return _Vm.Contract.GetRecordedLogs(&_Vm.TransactOpts)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmTransactorSession) GetRecordedLogs() (*types.Transaction, error) {
	return _Vm.Contract.GetRecordedLogs(&_Vm.TransactOpts)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_Vm *VmTransactor) IsDir(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "isDir", path)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_Vm *VmSession) IsDir(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsDir(&_Vm.TransactOpts, path)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_Vm *VmTransactorSession) IsDir(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsDir(&_Vm.TransactOpts, path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_Vm *VmTransactor) IsFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "isFile", path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_Vm *VmSession) IsFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsFile(&_Vm.TransactOpts, path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_Vm *VmTransactorSession) IsFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.IsFile(&_Vm.TransactOpts, path)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmTransactor) Label(opts *bind.TransactOpts, account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "label", account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.Contract.Label(&_Vm.TransactOpts, account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmTransactorSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.Contract.Label(&_Vm.TransactOpts, account, newLabel)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmTransactor) LoadAllocs(opts *bind.TransactOpts, pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "loadAllocs", pathToAllocsJson)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmSession) LoadAllocs(pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.Contract.LoadAllocs(&_Vm.TransactOpts, pathToAllocsJson)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmTransactorSession) LoadAllocs(pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.Contract.LoadAllocs(&_Vm.TransactOpts, pathToAllocsJson)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmTransactor) MakePersistent(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent", accounts)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmSession) MakePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent(&_Vm.TransactOpts, accounts)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmTransactorSession) MakePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent(&_Vm.TransactOpts, accounts)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmTransactor) MakePersistent0(opts *bind.TransactOpts, account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent0", account0, account1)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmSession) MakePersistent0(account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent0(&_Vm.TransactOpts, account0, account1)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmTransactorSession) MakePersistent0(account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent0(&_Vm.TransactOpts, account0, account1)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmTransactor) MakePersistent1(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent1", account)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmSession) MakePersistent1(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent1(&_Vm.TransactOpts, account)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmTransactorSession) MakePersistent1(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent1(&_Vm.TransactOpts, account)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmTransactor) MakePersistent2(opts *bind.TransactOpts, account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent2", account0, account1, account2)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmSession) MakePersistent2(account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent2(&_Vm.TransactOpts, account0, account1, account2)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmTransactorSession) MakePersistent2(account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent2(&_Vm.TransactOpts, account0, account1, account2)
}

// MockCall is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall", callee, msgValue, data, returnData)
}

// MockCall is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmSession) MockCall(callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall(callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall0(opts *bind.TransactOpts, callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall0", callee, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmSession) MockCall0(callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall0(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall0(callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall0(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert", callee, msgValue, data, revertData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert(callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert(callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert0(opts *bind.TransactOpts, callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert0", callee, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert0(callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert0(&_Vm.TransactOpts, callee, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert0(callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert0(&_Vm.TransactOpts, callee, data, revertData)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmTransactor) PauseGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "pauseGasMetering")
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmSession) PauseGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.PauseGasMetering(&_Vm.TransactOpts)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmTransactorSession) PauseGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.PauseGasMetering(&_Vm.TransactOpts)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactor) Prank(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank", msgSender, txOrigin)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmSession) Prank(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank(&_Vm.TransactOpts, msgSender, txOrigin)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactorSession) Prank(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank(&_Vm.TransactOpts, msgSender, txOrigin)
}

// Prank0 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmTransactor) Prank0(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank0", msgSender)
}

// Prank0 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmSession) Prank0(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank0(&_Vm.TransactOpts, msgSender)
}

// Prank0 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmTransactorSession) Prank0(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank0(&_Vm.TransactOpts, msgSender)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmTransactor) Prevrandao(opts *bind.TransactOpts, newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prevrandao", newPrevrandao)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmSession) Prevrandao(newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmTransactorSession) Prevrandao(newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao(&_Vm.TransactOpts, newPrevrandao)
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmTransactor) ReadCallers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "readCallers")
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmSession) ReadCallers() (*types.Transaction, error) {
	return _Vm.Contract.ReadCallers(&_Vm.TransactOpts)
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmTransactorSession) ReadCallers() (*types.Transaction, error) {
	return _Vm.Contract.ReadCallers(&_Vm.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmTransactor) Record(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "record")
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmSession) Record() (*types.Transaction, error) {
	return _Vm.Contract.Record(&_Vm.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmTransactorSession) Record() (*types.Transaction, error) {
	return _Vm.Contract.Record(&_Vm.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmTransactor) RecordLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "recordLogs")
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmSession) RecordLogs() (*types.Transaction, error) {
	return _Vm.Contract.RecordLogs(&_Vm.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmTransactorSession) RecordLogs() (*types.Transaction, error) {
	return _Vm.Contract.RecordLogs(&_Vm.TransactOpts)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmTransactor) RememberKey(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rememberKey", privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RememberKey(&_Vm.TransactOpts, privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmTransactorSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RememberKey(&_Vm.TransactOpts, privateKey)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmTransactor) RemoveDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "removeDir", path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.RemoveDir(&_Vm.TransactOpts, path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmTransactorSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.RemoveDir(&_Vm.TransactOpts, path, recursive)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmTransactor) RemoveFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "removeFile", path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmSession) RemoveFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.RemoveFile(&_Vm.TransactOpts, path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmTransactorSession) RemoveFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.RemoveFile(&_Vm.TransactOpts, path)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmTransactor) ResetNonce(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resetNonce", account)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmSession) ResetNonce(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ResetNonce(&_Vm.TransactOpts, account)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmTransactorSession) ResetNonce(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ResetNonce(&_Vm.TransactOpts, account)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmTransactor) ResumeGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resumeGasMetering")
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmSession) ResumeGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResumeGasMetering(&_Vm.TransactOpts)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmTransactorSession) ResumeGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResumeGasMetering(&_Vm.TransactOpts)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertTo(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertTo", snapshotId)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertTo(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertTo(&_Vm.TransactOpts, snapshotId)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertTo(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertTo(&_Vm.TransactOpts, snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertToAndDelete(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertToAndDelete", snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertToAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertToAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmTransactor) RevokePersistent(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revokePersistent", accounts)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmSession) RevokePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent(&_Vm.TransactOpts, accounts)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmTransactorSession) RevokePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent(&_Vm.TransactOpts, accounts)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmTransactor) RevokePersistent0(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revokePersistent0", account)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmSession) RevokePersistent0(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent0(&_Vm.TransactOpts, account)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmTransactorSession) RevokePersistent0(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent0(&_Vm.TransactOpts, account)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmTransactor) Roll(opts *bind.TransactOpts, newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "roll", newHeight)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmSession) Roll(newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Roll(&_Vm.TransactOpts, newHeight)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmTransactorSession) Roll(newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Roll(&_Vm.TransactOpts, newHeight)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmTransactor) RollFork(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork", txHash)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmSession) RollFork(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork(&_Vm.TransactOpts, txHash)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmTransactorSession) RollFork(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork(&_Vm.TransactOpts, txHash)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmTransactor) RollFork0(opts *bind.TransactOpts, forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork0", forkId, blockNumber)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmSession) RollFork0(forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork0(&_Vm.TransactOpts, forkId, blockNumber)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmTransactorSession) RollFork0(forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork0(&_Vm.TransactOpts, forkId, blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmTransactor) RollFork1(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork1", blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmSession) RollFork1(blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork1(&_Vm.TransactOpts, blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmTransactorSession) RollFork1(blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork1(&_Vm.TransactOpts, blockNumber)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactor) RollFork2(opts *bind.TransactOpts, forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork2", forkId, txHash)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmSession) RollFork2(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork2(&_Vm.TransactOpts, forkId, txHash)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactorSession) RollFork2(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork2(&_Vm.TransactOpts, forkId, txHash)
}

// Rpc is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmTransactor) Rpc(opts *bind.TransactOpts, method string, params string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rpc", method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmSession) Rpc(method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc(&_Vm.TransactOpts, method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmTransactorSession) Rpc(method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc(&_Vm.TransactOpts, method, params)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmTransactor) SelectFork(opts *bind.TransactOpts, forkId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "selectFork", forkId)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmSession) SelectFork(forkId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SelectFork(&_Vm.TransactOpts, forkId)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmTransactorSession) SelectFork(forkId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SelectFork(&_Vm.TransactOpts, forkId)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmTransactor) SerializeAddress(opts *bind.TransactOpts, objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeAddress", objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmTransactor) SerializeAddress0(opts *bind.TransactOpts, objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeAddress0", objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmTransactorSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBool(opts *bind.TransactOpts, objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBool", objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmTransactor) SerializeBool0(opts *bind.TransactOpts, objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBool0", objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBytes(opts *bind.TransactOpts, objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes", objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmTransactor) SerializeBytes0(opts *bind.TransactOpts, objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes0", objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBytes32(opts *bind.TransactOpts, objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes32", objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes32(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes32(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmTransactor) SerializeBytes320(opts *bind.TransactOpts, objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes320", objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes320(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes320(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmTransactor) SerializeInt(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeInt", objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmTransactor) SerializeInt0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeInt0", objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmTransactor) SerializeJson(opts *bind.TransactOpts, objectKey string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeJson", objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJson(&_Vm.TransactOpts, objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmTransactorSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJson(&_Vm.TransactOpts, objectKey, value)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmTransactor) SerializeString(opts *bind.TransactOpts, objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeString", objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmTransactor) SerializeString0(opts *bind.TransactOpts, objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeString0", objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmTransactorSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactor) SerializeUint(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUint", objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmTransactor) SerializeUint0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUint0", objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmTransactor) SetEnv(opts *bind.TransactOpts, name string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setEnv", name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SetEnv(&_Vm.TransactOpts, name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmTransactorSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SetEnv(&_Vm.TransactOpts, name, value)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmTransactor) SetNonce(opts *bind.TransactOpts, account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setNonce", account, newNonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmSession) SetNonce(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonce(&_Vm.TransactOpts, account, newNonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmTransactorSession) SetNonce(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonce(&_Vm.TransactOpts, account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmTransactor) SetNonceUnsafe(opts *bind.TransactOpts, account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setNonceUnsafe", account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmSession) SetNonceUnsafe(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonceUnsafe(&_Vm.TransactOpts, account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmTransactorSession) SetNonceUnsafe(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonceUnsafe(&_Vm.TransactOpts, account, newNonce)
}

// Sign is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmTransactor) Sign(opts *bind.TransactOpts, wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sign", wallet, digest)
}

// Sign is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Sign(&_Vm.TransactOpts, wallet, digest)
}

// Sign is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmTransactorSession) Sign(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Sign(&_Vm.TransactOpts, wallet, digest)
}

// Skip is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmTransactor) Skip(opts *bind.TransactOpts, skipTest bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "skip", skipTest)
}

// Skip is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmSession) Skip(skipTest bool) (*types.Transaction, error) {
	return _Vm.Contract.Skip(&_Vm.TransactOpts, skipTest)
}

// Skip is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmTransactorSession) Skip(skipTest bool) (*types.Transaction, error) {
	return _Vm.Contract.Skip(&_Vm.TransactOpts, skipTest)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmTransactor) Sleep(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sleep", duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sleep(&_Vm.TransactOpts, duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmTransactorSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sleep(&_Vm.TransactOpts, duration)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmSession) Snapshot() (*types.Transaction, error) {
	return _Vm.Contract.Snapshot(&_Vm.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmTransactorSession) Snapshot() (*types.Transaction, error) {
	return _Vm.Contract.Snapshot(&_Vm.TransactOpts)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmTransactor) StartBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast")
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmSession) StartBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast(&_Vm.TransactOpts)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmTransactorSession) StartBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast(&_Vm.TransactOpts)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmTransactor) StartBroadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast0", signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast0(&_Vm.TransactOpts, signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmTransactorSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast0(&_Vm.TransactOpts, signer)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmTransactor) StartBroadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast1", privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast1(&_Vm.TransactOpts, privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmTransactorSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast1(&_Vm.TransactOpts, privateKey)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmTransactor) StartMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startMappingRecording")
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmSession) StartMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartMappingRecording(&_Vm.TransactOpts)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmTransactorSession) StartMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartMappingRecording(&_Vm.TransactOpts)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmTransactor) StartPrank(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank", msgSender)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmSession) StartPrank(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank(&_Vm.TransactOpts, msgSender)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmTransactorSession) StartPrank(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank(&_Vm.TransactOpts, msgSender)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactor) StartPrank0(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank0", msgSender, txOrigin)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmSession) StartPrank0(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank0(&_Vm.TransactOpts, msgSender, txOrigin)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactorSession) StartPrank0(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank0(&_Vm.TransactOpts, msgSender, txOrigin)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmTransactor) StartStateDiffRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startStateDiffRecording")
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartStateDiffRecording(&_Vm.TransactOpts)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmTransactorSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartStateDiffRecording(&_Vm.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[])[] accountAccesses)
func (_Vm *VmTransactor) StopAndReturnStateDiff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopAndReturnStateDiff")
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[])[] accountAccesses)
func (_Vm *VmSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnStateDiff(&_Vm.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[])[] accountAccesses)
func (_Vm *VmTransactorSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnStateDiff(&_Vm.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmTransactor) StopBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopBroadcast")
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmSession) StopBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StopBroadcast(&_Vm.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmTransactorSession) StopBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StopBroadcast(&_Vm.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmTransactor) StopMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopMappingRecording")
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmSession) StopMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopMappingRecording(&_Vm.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmTransactorSession) StopMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopMappingRecording(&_Vm.TransactOpts)
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmTransactor) StopPrank(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopPrank")
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmSession) StopPrank() (*types.Transaction, error) {
	return _Vm.Contract.StopPrank(&_Vm.TransactOpts)
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmTransactorSession) StopPrank() (*types.Transaction, error) {
	return _Vm.Contract.StopPrank(&_Vm.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmTransactor) Store(opts *bind.TransactOpts, target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "store", target, slot, value)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmSession) Store(target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Store(&_Vm.TransactOpts, target, slot, value)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmTransactorSession) Store(target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Store(&_Vm.TransactOpts, target, slot, value)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactor) Transact(opts *bind.TransactOpts, forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "transact", forkId, txHash)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmSession) Transact(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact(&_Vm.TransactOpts, forkId, txHash)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactorSession) Transact(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact(&_Vm.TransactOpts, forkId, txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmTransactor) Transact0(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "transact0", txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmSession) Transact0(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact0(&_Vm.TransactOpts, txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmTransactorSession) Transact0(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact0(&_Vm.TransactOpts, txHash)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmTransactor) TryFfi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "tryFfi", commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.TryFfi(&_Vm.TransactOpts, commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmTransactorSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.TryFfi(&_Vm.TransactOpts, commandInput)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmTransactor) TxGasPrice(opts *bind.TransactOpts, newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "txGasPrice", newGasPrice)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmSession) TxGasPrice(newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.TxGasPrice(&_Vm.TransactOpts, newGasPrice)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmTransactorSession) TxGasPrice(newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.TxGasPrice(&_Vm.TransactOpts, newGasPrice)
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_Vm *VmTransactor) UnixTime(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "unixTime")
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_Vm *VmSession) UnixTime() (*types.Transaction, error) {
	return _Vm.Contract.UnixTime(&_Vm.TransactOpts)
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_Vm *VmTransactorSession) UnixTime() (*types.Transaction, error) {
	return _Vm.Contract.UnixTime(&_Vm.TransactOpts)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmTransactor) Warp(opts *bind.TransactOpts, newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "warp", newTimestamp)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmSession) Warp(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Warp(&_Vm.TransactOpts, newTimestamp)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmTransactorSession) Warp(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Warp(&_Vm.TransactOpts, newTimestamp)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmTransactor) WriteFile(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeFile", path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteFile(&_Vm.TransactOpts, path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmTransactorSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteFile(&_Vm.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmTransactor) WriteFileBinary(opts *bind.TransactOpts, path string, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeFileBinary", path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.WriteFileBinary(&_Vm.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmTransactorSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.WriteFileBinary(&_Vm.TransactOpts, path, data)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmTransactor) WriteJson(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeJson", json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmTransactorSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmTransactor) WriteJson0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeJson0", json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson0(&_Vm.TransactOpts, json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmTransactorSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson0(&_Vm.TransactOpts, json, path)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmTransactor) WriteLine(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeLine", path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteLine(&_Vm.TransactOpts, path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmTransactorSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteLine(&_Vm.TransactOpts, path, data)
}
