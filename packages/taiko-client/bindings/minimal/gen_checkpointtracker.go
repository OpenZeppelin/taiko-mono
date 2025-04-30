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

// ICheckpointTrackerCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ICheckpointTrackerCheckpoint struct {
	PublicationId *big.Int
	Commitment    [32]byte
}

// ICheckpointTrackerMetaData contains all meta data concerning the ICheckpointTracker contract.
var ICheckpointTrackerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getProvenCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveTransition\",\"inputs\":[{\"name\":\"start\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"end\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"numPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CheckpointUpdated\",\"inputs\":[{\"name\":\"latestCheckpoint\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false}]",
}

// ICheckpointTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use ICheckpointTrackerMetaData.ABI instead.
var ICheckpointTrackerABI = ICheckpointTrackerMetaData.ABI

// ICheckpointTracker is an auto generated Go binding around an Ethereum contract.
type ICheckpointTracker struct {
	ICheckpointTrackerCaller     // Read-only binding to the contract
	ICheckpointTrackerTransactor // Write-only binding to the contract
	ICheckpointTrackerFilterer   // Log filterer for contract events
}

// ICheckpointTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckpointTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckpointTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckpointTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICheckpointTrackerSession struct {
	Contract     *ICheckpointTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICheckpointTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICheckpointTrackerCallerSession struct {
	Contract *ICheckpointTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ICheckpointTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICheckpointTrackerTransactorSession struct {
	Contract     *ICheckpointTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ICheckpointTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICheckpointTrackerRaw struct {
	Contract *ICheckpointTracker // Generic contract binding to access the raw methods on
}

// ICheckpointTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICheckpointTrackerCallerRaw struct {
	Contract *ICheckpointTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// ICheckpointTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICheckpointTrackerTransactorRaw struct {
	Contract *ICheckpointTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICheckpointTracker creates a new instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTracker(address common.Address, backend bind.ContractBackend) (*ICheckpointTracker, error) {
	contract, err := bindICheckpointTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTracker{ICheckpointTrackerCaller: ICheckpointTrackerCaller{contract: contract}, ICheckpointTrackerTransactor: ICheckpointTrackerTransactor{contract: contract}, ICheckpointTrackerFilterer: ICheckpointTrackerFilterer{contract: contract}}, nil
}

// NewICheckpointTrackerCaller creates a new read-only instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTrackerCaller(address common.Address, caller bind.ContractCaller) (*ICheckpointTrackerCaller, error) {
	contract, err := bindICheckpointTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerCaller{contract: contract}, nil
}

// NewICheckpointTrackerTransactor creates a new write-only instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*ICheckpointTrackerTransactor, error) {
	contract, err := bindICheckpointTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerTransactor{contract: contract}, nil
}

// NewICheckpointTrackerFilterer creates a new log filterer instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*ICheckpointTrackerFilterer, error) {
	contract, err := bindICheckpointTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerFilterer{contract: contract}, nil
}

// bindICheckpointTracker binds a generic wrapper to an already deployed contract.
func bindICheckpointTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICheckpointTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointTracker *ICheckpointTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICheckpointTracker.Contract.ICheckpointTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointTracker *ICheckpointTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ICheckpointTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointTracker *ICheckpointTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ICheckpointTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointTracker *ICheckpointTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICheckpointTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointTracker *ICheckpointTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointTracker *ICheckpointTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.contract.Transact(opts, method, params...)
}

// GetProvenCheckpoint is a free data retrieval call binding the contract method 0xfa70b52c.
//
// Solidity: function getProvenCheckpoint() view returns((uint256,bytes32))
func (_ICheckpointTracker *ICheckpointTrackerCaller) GetProvenCheckpoint(opts *bind.CallOpts) (ICheckpointTrackerCheckpoint, error) {
	var out []interface{}
	err := _ICheckpointTracker.contract.Call(opts, &out, "getProvenCheckpoint")

	if err != nil {
		return *new(ICheckpointTrackerCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(ICheckpointTrackerCheckpoint)).(*ICheckpointTrackerCheckpoint)

	return out0, err

}

// GetProvenCheckpoint is a free data retrieval call binding the contract method 0xfa70b52c.
//
// Solidity: function getProvenCheckpoint() view returns((uint256,bytes32))
func (_ICheckpointTracker *ICheckpointTrackerSession) GetProvenCheckpoint() (ICheckpointTrackerCheckpoint, error) {
	return _ICheckpointTracker.Contract.GetProvenCheckpoint(&_ICheckpointTracker.CallOpts)
}

// GetProvenCheckpoint is a free data retrieval call binding the contract method 0xfa70b52c.
//
// Solidity: function getProvenCheckpoint() view returns((uint256,bytes32))
func (_ICheckpointTracker *ICheckpointTrackerCallerSession) GetProvenCheckpoint() (ICheckpointTrackerCheckpoint, error) {
	return _ICheckpointTracker.Contract.GetProvenCheckpoint(&_ICheckpointTracker.CallOpts)
}

// ProveTransition is a paid mutator transaction binding the contract method 0x628a48ed.
//
// Solidity: function proveTransition((uint256,bytes32) start, (uint256,bytes32) end, uint256 numPublications, uint256 numDelayedPublications, bytes proof) returns()
func (_ICheckpointTracker *ICheckpointTrackerTransactor) ProveTransition(opts *bind.TransactOpts, start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, numPublications *big.Int, numDelayedPublications *big.Int, proof []byte) (*types.Transaction, error) {
	return _ICheckpointTracker.contract.Transact(opts, "proveTransition", start, end, numPublications, numDelayedPublications, proof)
}

// ProveTransition is a paid mutator transaction binding the contract method 0x628a48ed.
//
// Solidity: function proveTransition((uint256,bytes32) start, (uint256,bytes32) end, uint256 numPublications, uint256 numDelayedPublications, bytes proof) returns()
func (_ICheckpointTracker *ICheckpointTrackerSession) ProveTransition(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, numPublications *big.Int, numDelayedPublications *big.Int, proof []byte) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ProveTransition(&_ICheckpointTracker.TransactOpts, start, end, numPublications, numDelayedPublications, proof)
}

// ProveTransition is a paid mutator transaction binding the contract method 0x628a48ed.
//
// Solidity: function proveTransition((uint256,bytes32) start, (uint256,bytes32) end, uint256 numPublications, uint256 numDelayedPublications, bytes proof) returns()
func (_ICheckpointTracker *ICheckpointTrackerTransactorSession) ProveTransition(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, numPublications *big.Int, numDelayedPublications *big.Int, proof []byte) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ProveTransition(&_ICheckpointTracker.TransactOpts, start, end, numPublications, numDelayedPublications, proof)
}

// ICheckpointTrackerCheckpointUpdatedIterator is returned from FilterCheckpointUpdated and is used to iterate over the raw logs and unpacked data for CheckpointUpdated events raised by the ICheckpointTracker contract.
type ICheckpointTrackerCheckpointUpdatedIterator struct {
	Event *ICheckpointTrackerCheckpointUpdated // Event containing the contract specifics and raw log

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
func (it *ICheckpointTrackerCheckpointUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICheckpointTrackerCheckpointUpdated)
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
		it.Event = new(ICheckpointTrackerCheckpointUpdated)
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
func (it *ICheckpointTrackerCheckpointUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICheckpointTrackerCheckpointUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICheckpointTrackerCheckpointUpdated represents a CheckpointUpdated event raised by the ICheckpointTracker contract.
type ICheckpointTrackerCheckpointUpdated struct {
	LatestCheckpoint ICheckpointTrackerCheckpoint
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCheckpointUpdated is a free log retrieval operation binding the contract event 0x9a00cdcfd3a2f2a65df16ddf7b3ed3a0c0649669f2395fdb5bd55f0a2fd697be.
//
// Solidity: event CheckpointUpdated((uint256,bytes32) latestCheckpoint)
func (_ICheckpointTracker *ICheckpointTrackerFilterer) FilterCheckpointUpdated(opts *bind.FilterOpts) (*ICheckpointTrackerCheckpointUpdatedIterator, error) {

	logs, sub, err := _ICheckpointTracker.contract.FilterLogs(opts, "CheckpointUpdated")
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerCheckpointUpdatedIterator{contract: _ICheckpointTracker.contract, event: "CheckpointUpdated", logs: logs, sub: sub}, nil
}

// WatchCheckpointUpdated is a free log subscription operation binding the contract event 0x9a00cdcfd3a2f2a65df16ddf7b3ed3a0c0649669f2395fdb5bd55f0a2fd697be.
//
// Solidity: event CheckpointUpdated((uint256,bytes32) latestCheckpoint)
func (_ICheckpointTracker *ICheckpointTrackerFilterer) WatchCheckpointUpdated(opts *bind.WatchOpts, sink chan<- *ICheckpointTrackerCheckpointUpdated) (event.Subscription, error) {

	logs, sub, err := _ICheckpointTracker.contract.WatchLogs(opts, "CheckpointUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICheckpointTrackerCheckpointUpdated)
				if err := _ICheckpointTracker.contract.UnpackLog(event, "CheckpointUpdated", log); err != nil {
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

// ParseCheckpointUpdated is a log parse operation binding the contract event 0x9a00cdcfd3a2f2a65df16ddf7b3ed3a0c0649669f2395fdb5bd55f0a2fd697be.
//
// Solidity: event CheckpointUpdated((uint256,bytes32) latestCheckpoint)
func (_ICheckpointTracker *ICheckpointTrackerFilterer) ParseCheckpointUpdated(log types.Log) (*ICheckpointTrackerCheckpointUpdated, error) {
	event := new(ICheckpointTrackerCheckpointUpdated)
	if err := _ICheckpointTracker.contract.UnpackLog(event, "CheckpointUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
