// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package minimal

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

// ITaikoAnchorBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type ITaikoAnchorBlockHeader struct {
	ParentHash            [32]byte
	OmnersHash            [32]byte
	Coinbase              common.Address
	StateRoot             [32]byte
	TransactionsRoot      [32]byte
	ReceiptsRoot          [32]byte
	LogsBloom             []byte
	Difficulty            *big.Int
	Number                *big.Int
	GasLimit              uint64
	GasUsed               uint64
	Timestamp             uint64
	ExtraData             []byte
	MixedHash             [32]byte
	Nonce                 uint64
	BaseFeePerGas         *big.Int
	WithdrawalsRoot       [32]byte
	BlobGasUsed           uint64
	ExcessBlobGas         uint64
	ParentBeaconBlockRoot [32]byte
	RequestsHash          [32]byte
}

// ITaikoAnchorMetaData contains all meta data concerning the ITaikoAnchor contract.
var ITaikoAnchorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"anchor\",\"inputs\":[{\"name\":\"_publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_anchorBlockId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_anchorBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_anchorBlockHeader\",\"type\":\"tuple\",\"internalType\":\"structITaikoAnchor.BlockHeader\",\"components\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"omnersHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"coinbase\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiptsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logsBloom\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mixedHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawalsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blobGasUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excessBlobGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parentBeaconBlockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_parentGasUsed\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBaseFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPermissionedSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1BlockHashes\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Anchor\",\"inputs\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"anchorBlockId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"anchorBlockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"parentGasUsed\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
}

// ITaikoAnchorABI is the input ABI used to generate the binding from.
// Deprecated: Use ITaikoAnchorMetaData.ABI instead.
var ITaikoAnchorABI = ITaikoAnchorMetaData.ABI

// ITaikoAnchor is an auto generated Go binding around an Ethereum contract.
type ITaikoAnchor struct {
	ITaikoAnchorCaller     // Read-only binding to the contract
	ITaikoAnchorTransactor // Write-only binding to the contract
	ITaikoAnchorFilterer   // Log filterer for contract events
}

// ITaikoAnchorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITaikoAnchorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITaikoAnchorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITaikoAnchorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITaikoAnchorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITaikoAnchorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITaikoAnchorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITaikoAnchorSession struct {
	Contract     *ITaikoAnchor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITaikoAnchorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITaikoAnchorCallerSession struct {
	Contract *ITaikoAnchorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITaikoAnchorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITaikoAnchorTransactorSession struct {
	Contract     *ITaikoAnchorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITaikoAnchorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITaikoAnchorRaw struct {
	Contract *ITaikoAnchor // Generic contract binding to access the raw methods on
}

// ITaikoAnchorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITaikoAnchorCallerRaw struct {
	Contract *ITaikoAnchorCaller // Generic read-only contract binding to access the raw methods on
}

// ITaikoAnchorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITaikoAnchorTransactorRaw struct {
	Contract *ITaikoAnchorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITaikoAnchor creates a new instance of ITaikoAnchor, bound to a specific deployed contract.
func NewITaikoAnchor(address common.Address, backend bind.ContractBackend) (*ITaikoAnchor, error) {
	contract, err := bindITaikoAnchor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITaikoAnchor{ITaikoAnchorCaller: ITaikoAnchorCaller{contract: contract}, ITaikoAnchorTransactor: ITaikoAnchorTransactor{contract: contract}, ITaikoAnchorFilterer: ITaikoAnchorFilterer{contract: contract}}, nil
}

// NewITaikoAnchorCaller creates a new read-only instance of ITaikoAnchor, bound to a specific deployed contract.
func NewITaikoAnchorCaller(address common.Address, caller bind.ContractCaller) (*ITaikoAnchorCaller, error) {
	contract, err := bindITaikoAnchor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITaikoAnchorCaller{contract: contract}, nil
}

// NewITaikoAnchorTransactor creates a new write-only instance of ITaikoAnchor, bound to a specific deployed contract.
func NewITaikoAnchorTransactor(address common.Address, transactor bind.ContractTransactor) (*ITaikoAnchorTransactor, error) {
	contract, err := bindITaikoAnchor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITaikoAnchorTransactor{contract: contract}, nil
}

// NewITaikoAnchorFilterer creates a new log filterer instance of ITaikoAnchor, bound to a specific deployed contract.
func NewITaikoAnchorFilterer(address common.Address, filterer bind.ContractFilterer) (*ITaikoAnchorFilterer, error) {
	contract, err := bindITaikoAnchor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITaikoAnchorFilterer{contract: contract}, nil
}

// bindITaikoAnchor binds a generic wrapper to an already deployed contract.
func bindITaikoAnchor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITaikoAnchorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITaikoAnchor *ITaikoAnchorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITaikoAnchor.Contract.ITaikoAnchorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITaikoAnchor *ITaikoAnchorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITaikoAnchor.Contract.ITaikoAnchorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITaikoAnchor *ITaikoAnchorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITaikoAnchor.Contract.ITaikoAnchorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITaikoAnchor *ITaikoAnchorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITaikoAnchor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITaikoAnchor *ITaikoAnchorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITaikoAnchor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITaikoAnchor *ITaikoAnchorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITaikoAnchor.Contract.contract.Transact(opts, method, params...)
}

// GetBaseFee is a free data retrieval call binding the contract method 0x15e812ad.
//
// Solidity: function getBaseFee() view returns(uint256)
func (_ITaikoAnchor *ITaikoAnchorCaller) GetBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ITaikoAnchor.contract.Call(opts, &out, "getBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseFee is a free data retrieval call binding the contract method 0x15e812ad.
//
// Solidity: function getBaseFee() view returns(uint256)
func (_ITaikoAnchor *ITaikoAnchorSession) GetBaseFee() (*big.Int, error) {
	return _ITaikoAnchor.Contract.GetBaseFee(&_ITaikoAnchor.CallOpts)
}

// GetBaseFee is a free data retrieval call binding the contract method 0x15e812ad.
//
// Solidity: function getBaseFee() view returns(uint256)
func (_ITaikoAnchor *ITaikoAnchorCallerSession) GetBaseFee() (*big.Int, error) {
	return _ITaikoAnchor.Contract.GetBaseFee(&_ITaikoAnchor.CallOpts)
}

// GetPermissionedSender is a free data retrieval call binding the contract method 0x77c4ee59.
//
// Solidity: function getPermissionedSender() view returns(address)
func (_ITaikoAnchor *ITaikoAnchorCaller) GetPermissionedSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITaikoAnchor.contract.Call(opts, &out, "getPermissionedSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionedSender is a free data retrieval call binding the contract method 0x77c4ee59.
//
// Solidity: function getPermissionedSender() view returns(address)
func (_ITaikoAnchor *ITaikoAnchorSession) GetPermissionedSender() (common.Address, error) {
	return _ITaikoAnchor.Contract.GetPermissionedSender(&_ITaikoAnchor.CallOpts)
}

// GetPermissionedSender is a free data retrieval call binding the contract method 0x77c4ee59.
//
// Solidity: function getPermissionedSender() view returns(address)
func (_ITaikoAnchor *ITaikoAnchorCallerSession) GetPermissionedSender() (common.Address, error) {
	return _ITaikoAnchor.Contract.GetPermissionedSender(&_ITaikoAnchor.CallOpts)
}

// L1BlockHashes is a free data retrieval call binding the contract method 0x8d7a1d5f.
//
// Solidity: function l1BlockHashes(uint256 blockId) view returns(bytes32 blockHash)
func (_ITaikoAnchor *ITaikoAnchorCaller) L1BlockHashes(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ITaikoAnchor.contract.Call(opts, &out, "l1BlockHashes", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1BlockHashes is a free data retrieval call binding the contract method 0x8d7a1d5f.
//
// Solidity: function l1BlockHashes(uint256 blockId) view returns(bytes32 blockHash)
func (_ITaikoAnchor *ITaikoAnchorSession) L1BlockHashes(blockId *big.Int) ([32]byte, error) {
	return _ITaikoAnchor.Contract.L1BlockHashes(&_ITaikoAnchor.CallOpts, blockId)
}

// L1BlockHashes is a free data retrieval call binding the contract method 0x8d7a1d5f.
//
// Solidity: function l1BlockHashes(uint256 blockId) view returns(bytes32 blockHash)
func (_ITaikoAnchor *ITaikoAnchorCallerSession) L1BlockHashes(blockId *big.Int) ([32]byte, error) {
	return _ITaikoAnchor.Contract.L1BlockHashes(&_ITaikoAnchor.CallOpts, blockId)
}

// Anchor is a paid mutator transaction binding the contract method 0x38b89131.
//
// Solidity: function anchor(uint256 _publicationId, uint256 _anchorBlockId, bytes32 _anchorBlockHash, (bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint64,uint64,uint64,bytes,bytes32,uint64,uint256,bytes32,uint64,uint64,bytes32,bytes32) _anchorBlockHeader, uint32 _parentGasUsed) returns()
func (_ITaikoAnchor *ITaikoAnchorTransactor) Anchor(opts *bind.TransactOpts, _publicationId *big.Int, _anchorBlockId *big.Int, _anchorBlockHash [32]byte, _anchorBlockHeader ITaikoAnchorBlockHeader, _parentGasUsed uint32) (*types.Transaction, error) {
	return _ITaikoAnchor.contract.Transact(opts, "anchor", _publicationId, _anchorBlockId, _anchorBlockHash, _anchorBlockHeader, _parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0x38b89131.
//
// Solidity: function anchor(uint256 _publicationId, uint256 _anchorBlockId, bytes32 _anchorBlockHash, (bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint64,uint64,uint64,bytes,bytes32,uint64,uint256,bytes32,uint64,uint64,bytes32,bytes32) _anchorBlockHeader, uint32 _parentGasUsed) returns()
func (_ITaikoAnchor *ITaikoAnchorSession) Anchor(_publicationId *big.Int, _anchorBlockId *big.Int, _anchorBlockHash [32]byte, _anchorBlockHeader ITaikoAnchorBlockHeader, _parentGasUsed uint32) (*types.Transaction, error) {
	return _ITaikoAnchor.Contract.Anchor(&_ITaikoAnchor.TransactOpts, _publicationId, _anchorBlockId, _anchorBlockHash, _anchorBlockHeader, _parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0x38b89131.
//
// Solidity: function anchor(uint256 _publicationId, uint256 _anchorBlockId, bytes32 _anchorBlockHash, (bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint64,uint64,uint64,bytes,bytes32,uint64,uint256,bytes32,uint64,uint64,bytes32,bytes32) _anchorBlockHeader, uint32 _parentGasUsed) returns()
func (_ITaikoAnchor *ITaikoAnchorTransactorSession) Anchor(_publicationId *big.Int, _anchorBlockId *big.Int, _anchorBlockHash [32]byte, _anchorBlockHeader ITaikoAnchorBlockHeader, _parentGasUsed uint32) (*types.Transaction, error) {
	return _ITaikoAnchor.Contract.Anchor(&_ITaikoAnchor.TransactOpts, _publicationId, _anchorBlockId, _anchorBlockHash, _anchorBlockHeader, _parentGasUsed)
}

// ITaikoAnchorAnchorIterator is returned from FilterAnchor and is used to iterate over the raw logs and unpacked data for Anchor events raised by the ITaikoAnchor contract.
type ITaikoAnchorAnchorIterator struct {
	Event *ITaikoAnchorAnchor // Event containing the contract specifics and raw log

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
func (it *ITaikoAnchorAnchorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITaikoAnchorAnchor)
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
		it.Event = new(ITaikoAnchorAnchor)
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
func (it *ITaikoAnchorAnchorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITaikoAnchorAnchorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITaikoAnchorAnchor represents a Anchor event raised by the ITaikoAnchor contract.
type ITaikoAnchorAnchor struct {
	PublicationId   *big.Int
	AnchorBlockId   *big.Int
	AnchorBlockHash [32]byte
	ParentGasUsed   uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAnchor is a free log retrieval operation binding the contract event 0x9949542201b26d80d3033327db6b6c5da2a0de65e04797a9a94745f6502902c0.
//
// Solidity: event Anchor(uint256 publicationId, uint256 anchorBlockId, bytes32 anchorBlockHash, uint32 parentGasUsed)
func (_ITaikoAnchor *ITaikoAnchorFilterer) FilterAnchor(opts *bind.FilterOpts) (*ITaikoAnchorAnchorIterator, error) {

	logs, sub, err := _ITaikoAnchor.contract.FilterLogs(opts, "Anchor")
	if err != nil {
		return nil, err
	}
	return &ITaikoAnchorAnchorIterator{contract: _ITaikoAnchor.contract, event: "Anchor", logs: logs, sub: sub}, nil
}

// WatchAnchor is a free log subscription operation binding the contract event 0x9949542201b26d80d3033327db6b6c5da2a0de65e04797a9a94745f6502902c0.
//
// Solidity: event Anchor(uint256 publicationId, uint256 anchorBlockId, bytes32 anchorBlockHash, uint32 parentGasUsed)
func (_ITaikoAnchor *ITaikoAnchorFilterer) WatchAnchor(opts *bind.WatchOpts, sink chan<- *ITaikoAnchorAnchor) (event.Subscription, error) {

	logs, sub, err := _ITaikoAnchor.contract.WatchLogs(opts, "Anchor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITaikoAnchorAnchor)
				if err := _ITaikoAnchor.contract.UnpackLog(event, "Anchor", log); err != nil {
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

// ParseAnchor is a log parse operation binding the contract event 0x9949542201b26d80d3033327db6b6c5da2a0de65e04797a9a94745f6502902c0.
//
// Solidity: event Anchor(uint256 publicationId, uint256 anchorBlockId, bytes32 anchorBlockHash, uint32 parentGasUsed)
func (_ITaikoAnchor *ITaikoAnchorFilterer) ParseAnchor(log types.Log) (*ITaikoAnchorAnchor, error) {
	event := new(ITaikoAnchorAnchor)
	if err := _ITaikoAnchor.contract.UnpackLog(event, "Anchor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
