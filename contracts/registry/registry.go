// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// FileRegistryMetaData contains all meta data concerning the FileRegistry contract.
var FileRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FileRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FileRegistryMetaData.ABI instead.
var FileRegistryABI = FileRegistryMetaData.ABI

// FileRegistry is an auto generated Go binding around an Ethereum contract.
type FileRegistry struct {
	FileRegistryCaller     // Read-only binding to the contract
	FileRegistryTransactor // Write-only binding to the contract
	FileRegistryFilterer   // Log filterer for contract events
}

// FileRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileRegistrySession struct {
	Contract     *FileRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileRegistryCallerSession struct {
	Contract *FileRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FileRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileRegistryTransactorSession struct {
	Contract     *FileRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FileRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileRegistryRaw struct {
	Contract *FileRegistry // Generic contract binding to access the raw methods on
}

// FileRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileRegistryCallerRaw struct {
	Contract *FileRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FileRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileRegistryTransactorRaw struct {
	Contract *FileRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileRegistry creates a new instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistry(address common.Address, backend bind.ContractBackend) (*FileRegistry, error) {
	contract, err := bindFileRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileRegistry{FileRegistryCaller: FileRegistryCaller{contract: contract}, FileRegistryTransactor: FileRegistryTransactor{contract: contract}, FileRegistryFilterer: FileRegistryFilterer{contract: contract}}, nil
}

// NewFileRegistryCaller creates a new read-only instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistryCaller(address common.Address, caller bind.ContractCaller) (*FileRegistryCaller, error) {
	contract, err := bindFileRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileRegistryCaller{contract: contract}, nil
}

// NewFileRegistryTransactor creates a new write-only instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FileRegistryTransactor, error) {
	contract, err := bindFileRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileRegistryTransactor{contract: contract}, nil
}

// NewFileRegistryFilterer creates a new log filterer instance of FileRegistry, bound to a specific deployed contract.
func NewFileRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FileRegistryFilterer, error) {
	contract, err := bindFileRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileRegistryFilterer{contract: contract}, nil
}

// bindFileRegistry binds a generic wrapper to an already deployed contract.
func bindFileRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FileRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileRegistry *FileRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileRegistry.Contract.FileRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileRegistry *FileRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileRegistry.Contract.FileRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileRegistry *FileRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileRegistry.Contract.FileRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileRegistry *FileRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileRegistry *FileRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileRegistry *FileRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileRegistry.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x4909b29b.
//
// Solidity: function items(string ) view returns(string)
func (_FileRegistry *FileRegistryCaller) Items(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _FileRegistry.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x4909b29b.
//
// Solidity: function items(string ) view returns(string)
func (_FileRegistry *FileRegistrySession) Items(arg0 string) (string, error) {
	return _FileRegistry.Contract.Items(&_FileRegistry.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x4909b29b.
//
// Solidity: function items(string ) view returns(string)
func (_FileRegistry *FileRegistryCallerSession) Items(arg0 string) (string, error) {
	return _FileRegistry.Contract.Items(&_FileRegistry.CallOpts, arg0)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_FileRegistry *FileRegistryTransactor) Save(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error) {
	return _FileRegistry.contract.Transact(opts, "save", filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_FileRegistry *FileRegistrySession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _FileRegistry.Contract.Save(&_FileRegistry.TransactOpts, filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_FileRegistry *FileRegistryTransactorSession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _FileRegistry.Contract.Save(&_FileRegistry.TransactOpts, filePath, cid)
}

// FileRegistryItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the FileRegistry contract.
type FileRegistryItemSetIterator struct {
	Event *FileRegistryItemSet // Event containing the contract specifics and raw log

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
func (it *FileRegistryItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileRegistryItemSet)
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
		it.Event = new(FileRegistryItemSet)
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
func (it *FileRegistryItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileRegistryItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileRegistryItemSet represents a ItemSet event raised by the FileRegistry contract.
type FileRegistryItemSet struct {
	FilePath common.Hash
	Cid      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3.
//
// Solidity: event ItemSet(string indexed filePath, string cid)
func (_FileRegistry *FileRegistryFilterer) FilterItemSet(opts *bind.FilterOpts, filePath []string) (*FileRegistryItemSetIterator, error) {

	var filePathRule []interface{}
	for _, filePathItem := range filePath {
		filePathRule = append(filePathRule, filePathItem)
	}

	logs, sub, err := _FileRegistry.contract.FilterLogs(opts, "ItemSet", filePathRule)
	if err != nil {
		return nil, err
	}
	return &FileRegistryItemSetIterator{contract: _FileRegistry.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3.
//
// Solidity: event ItemSet(string indexed filePath, string cid)
func (_FileRegistry *FileRegistryFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *FileRegistryItemSet, filePath []string) (event.Subscription, error) {

	var filePathRule []interface{}
	for _, filePathItem := range filePath {
		filePathRule = append(filePathRule, filePathItem)
	}

	logs, sub, err := _FileRegistry.contract.WatchLogs(opts, "ItemSet", filePathRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileRegistryItemSet)
				if err := _FileRegistry.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0x523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3.
//
// Solidity: event ItemSet(string indexed filePath, string cid)
func (_FileRegistry *FileRegistryFilterer) ParseItemSet(log types.Log) (*FileRegistryItemSet, error) {
	event := new(FileRegistryItemSet)
	if err := _FileRegistry.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
