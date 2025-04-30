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

// IPublicationFeedPublicationHeader is an auto generated low-level Go binding around an user-defined struct.
type IPublicationFeedPublicationHeader struct {
	Id             *big.Int
	PrevHash       [32]byte
	Publisher      common.Address
	Timestamp      *big.Int
	BlockNumber    *big.Int
	AttributesHash [32]byte
}

// IPublicationFeedMetaData contains all meta data concerning the IPublicationFeed contract.
var IPublicationFeedMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNextPublicationId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublicationHash\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publish\",\"inputs\":[{\"name\":\"attributes\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIPublicationFeed.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateHeader\",\"inputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIPublicationFeed.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Published\",\"inputs\":[{\"name\":\"pubHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"header\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPublicationFeed.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"attributes\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false}]",
}

// IPublicationFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use IPublicationFeedMetaData.ABI instead.
var IPublicationFeedABI = IPublicationFeedMetaData.ABI

// IPublicationFeed is an auto generated Go binding around an Ethereum contract.
type IPublicationFeed struct {
	IPublicationFeedCaller     // Read-only binding to the contract
	IPublicationFeedTransactor // Write-only binding to the contract
	IPublicationFeedFilterer   // Log filterer for contract events
}

// IPublicationFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPublicationFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPublicationFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPublicationFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPublicationFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPublicationFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPublicationFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPublicationFeedSession struct {
	Contract     *IPublicationFeed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPublicationFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPublicationFeedCallerSession struct {
	Contract *IPublicationFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPublicationFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPublicationFeedTransactorSession struct {
	Contract     *IPublicationFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPublicationFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPublicationFeedRaw struct {
	Contract *IPublicationFeed // Generic contract binding to access the raw methods on
}

// IPublicationFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPublicationFeedCallerRaw struct {
	Contract *IPublicationFeedCaller // Generic read-only contract binding to access the raw methods on
}

// IPublicationFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPublicationFeedTransactorRaw struct {
	Contract *IPublicationFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPublicationFeed creates a new instance of IPublicationFeed, bound to a specific deployed contract.
func NewIPublicationFeed(address common.Address, backend bind.ContractBackend) (*IPublicationFeed, error) {
	contract, err := bindIPublicationFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPublicationFeed{IPublicationFeedCaller: IPublicationFeedCaller{contract: contract}, IPublicationFeedTransactor: IPublicationFeedTransactor{contract: contract}, IPublicationFeedFilterer: IPublicationFeedFilterer{contract: contract}}, nil
}

// NewIPublicationFeedCaller creates a new read-only instance of IPublicationFeed, bound to a specific deployed contract.
func NewIPublicationFeedCaller(address common.Address, caller bind.ContractCaller) (*IPublicationFeedCaller, error) {
	contract, err := bindIPublicationFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPublicationFeedCaller{contract: contract}, nil
}

// NewIPublicationFeedTransactor creates a new write-only instance of IPublicationFeed, bound to a specific deployed contract.
func NewIPublicationFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*IPublicationFeedTransactor, error) {
	contract, err := bindIPublicationFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPublicationFeedTransactor{contract: contract}, nil
}

// NewIPublicationFeedFilterer creates a new log filterer instance of IPublicationFeed, bound to a specific deployed contract.
func NewIPublicationFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*IPublicationFeedFilterer, error) {
	contract, err := bindIPublicationFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPublicationFeedFilterer{contract: contract}, nil
}

// bindIPublicationFeed binds a generic wrapper to an already deployed contract.
func bindIPublicationFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPublicationFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPublicationFeed *IPublicationFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPublicationFeed.Contract.IPublicationFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPublicationFeed *IPublicationFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPublicationFeed.Contract.IPublicationFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPublicationFeed *IPublicationFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPublicationFeed.Contract.IPublicationFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPublicationFeed *IPublicationFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPublicationFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPublicationFeed *IPublicationFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPublicationFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPublicationFeed *IPublicationFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPublicationFeed.Contract.contract.Transact(opts, method, params...)
}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_IPublicationFeed *IPublicationFeedCaller) GetNextPublicationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPublicationFeed.contract.Call(opts, &out, "getNextPublicationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_IPublicationFeed *IPublicationFeedSession) GetNextPublicationId() (*big.Int, error) {
	return _IPublicationFeed.Contract.GetNextPublicationId(&_IPublicationFeed.CallOpts)
}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_IPublicationFeed *IPublicationFeedCallerSession) GetNextPublicationId() (*big.Int, error) {
	return _IPublicationFeed.Contract.GetNextPublicationId(&_IPublicationFeed.CallOpts)
}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_IPublicationFeed *IPublicationFeedCaller) GetPublicationHash(opts *bind.CallOpts, idx *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IPublicationFeed.contract.Call(opts, &out, "getPublicationHash", idx)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_IPublicationFeed *IPublicationFeedSession) GetPublicationHash(idx *big.Int) ([32]byte, error) {
	return _IPublicationFeed.Contract.GetPublicationHash(&_IPublicationFeed.CallOpts, idx)
}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_IPublicationFeed *IPublicationFeedCallerSession) GetPublicationHash(idx *big.Int) ([32]byte, error) {
	return _IPublicationFeed.Contract.GetPublicationHash(&_IPublicationFeed.CallOpts, idx)
}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_IPublicationFeed *IPublicationFeedCaller) ValidateHeader(opts *bind.CallOpts, header IPublicationFeedPublicationHeader) (bool, error) {
	var out []interface{}
	err := _IPublicationFeed.contract.Call(opts, &out, "validateHeader", header)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_IPublicationFeed *IPublicationFeedSession) ValidateHeader(header IPublicationFeedPublicationHeader) (bool, error) {
	return _IPublicationFeed.Contract.ValidateHeader(&_IPublicationFeed.CallOpts, header)
}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_IPublicationFeed *IPublicationFeedCallerSession) ValidateHeader(header IPublicationFeedPublicationHeader) (bool, error) {
	return _IPublicationFeed.Contract.ValidateHeader(&_IPublicationFeed.CallOpts, header)
}

// Publish is a paid mutator transaction binding the contract method 0x59c06c89.
//
// Solidity: function publish(bytes[] attributes) returns((uint256,bytes32,address,uint256,uint256,bytes32) header)
func (_IPublicationFeed *IPublicationFeedTransactor) Publish(opts *bind.TransactOpts, attributes [][]byte) (*types.Transaction, error) {
	return _IPublicationFeed.contract.Transact(opts, "publish", attributes)
}

// Publish is a paid mutator transaction binding the contract method 0x59c06c89.
//
// Solidity: function publish(bytes[] attributes) returns((uint256,bytes32,address,uint256,uint256,bytes32) header)
func (_IPublicationFeed *IPublicationFeedSession) Publish(attributes [][]byte) (*types.Transaction, error) {
	return _IPublicationFeed.Contract.Publish(&_IPublicationFeed.TransactOpts, attributes)
}

// Publish is a paid mutator transaction binding the contract method 0x59c06c89.
//
// Solidity: function publish(bytes[] attributes) returns((uint256,bytes32,address,uint256,uint256,bytes32) header)
func (_IPublicationFeed *IPublicationFeedTransactorSession) Publish(attributes [][]byte) (*types.Transaction, error) {
	return _IPublicationFeed.Contract.Publish(&_IPublicationFeed.TransactOpts, attributes)
}

// IPublicationFeedPublishedIterator is returned from FilterPublished and is used to iterate over the raw logs and unpacked data for Published events raised by the IPublicationFeed contract.
type IPublicationFeedPublishedIterator struct {
	Event *IPublicationFeedPublished // Event containing the contract specifics and raw log

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
func (it *IPublicationFeedPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicationFeedPublished)
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
		it.Event = new(IPublicationFeedPublished)
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
func (it *IPublicationFeedPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicationFeedPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicationFeedPublished represents a Published event raised by the IPublicationFeed contract.
type IPublicationFeedPublished struct {
	PubHash    [32]byte
	Header     IPublicationFeedPublicationHeader
	Attributes [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPublished is a free log retrieval operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_IPublicationFeed *IPublicationFeedFilterer) FilterPublished(opts *bind.FilterOpts, pubHash [][32]byte) (*IPublicationFeedPublishedIterator, error) {

	var pubHashRule []interface{}
	for _, pubHashItem := range pubHash {
		pubHashRule = append(pubHashRule, pubHashItem)
	}

	logs, sub, err := _IPublicationFeed.contract.FilterLogs(opts, "Published", pubHashRule)
	if err != nil {
		return nil, err
	}
	return &IPublicationFeedPublishedIterator{contract: _IPublicationFeed.contract, event: "Published", logs: logs, sub: sub}, nil
}

// WatchPublished is a free log subscription operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_IPublicationFeed *IPublicationFeedFilterer) WatchPublished(opts *bind.WatchOpts, sink chan<- *IPublicationFeedPublished, pubHash [][32]byte) (event.Subscription, error) {

	var pubHashRule []interface{}
	for _, pubHashItem := range pubHash {
		pubHashRule = append(pubHashRule, pubHashItem)
	}

	logs, sub, err := _IPublicationFeed.contract.WatchLogs(opts, "Published", pubHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicationFeedPublished)
				if err := _IPublicationFeed.contract.UnpackLog(event, "Published", log); err != nil {
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

// ParsePublished is a log parse operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_IPublicationFeed *IPublicationFeedFilterer) ParsePublished(log types.Log) (*IPublicationFeedPublished, error) {
	event := new(IPublicationFeedPublished)
	if err := _IPublicationFeed.contract.UnpackLog(event, "Published", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
