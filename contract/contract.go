// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5061055b8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80634909b29b14610043578063693ec85e1461006c578063962939b81461007f575b5f80fd5b6100566100513660046102fc565b610094565b6040516100639190610336565b60405180910390f35b61005661007a3660046102fc565b610136565b61009261008d36600461036b565b6101e3565b005b80516020818301810180515f82529282019190930120915280546100b7906103cb565b80601f01602080910402602001604051908101604052809291908181526020018280546100e3906103cb565b801561012e5780601f106101055761010080835404028352916020019161012e565b820191905f5260205f20905b81548152906001019060200180831161011157829003601f168201915b505050505081565b60605f826040516101479190610403565b90815260200160405180910390208054610160906103cb565b80601f016020809104026020016040519081016040528092919081815260200182805461018c906103cb565b80156101d75780601f106101ae576101008083540402835291602001916101d7565b820191905f5260205f20905b8154815290600101906020018083116101ba57829003601f168201915b50505050509050919050565b805f836040516101f39190610403565b9081526020016040518091039020908161020d9190610465565b508160405161021c9190610403565b60405180910390207f523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3826040516102539190610336565b60405180910390a25050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610282575f80fd5b813567ffffffffffffffff8082111561029d5761029d61025f565b604051601f8301601f19908116603f011681019082821181831017156102c5576102c561025f565b816040528381528660208588010111156102dd575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f6020828403121561030c575f80fd5b813567ffffffffffffffff811115610322575f80fd5b61032e84828501610273565b949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f806040838503121561037c575f80fd5b823567ffffffffffffffff80821115610393575f80fd5b61039f86838701610273565b935060208501359150808211156103b4575f80fd5b506103c185828601610273565b9150509250929050565b600181811c908216806103df57607f821691505b6020821081036103fd57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f82518060208501845e5f920191825250919050565b601f82111561046057805f5260205f20601f840160051c8101602085101561043e5750805b601f840160051c820191505b8181101561045d575f815560010161044a565b50505b505050565b815167ffffffffffffffff81111561047f5761047f61025f565b6104938161048d84546103cb565b84610419565b602080601f8311600181146104c6575f84156104af5750858301515b5f19600386901b1c1916600185901b17855561051d565b5f85815260208120601f198616915b828110156104f4578886015182559484019460019091019084016104d5565b508582101561051157878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212205059cd7168c2f4befc791432f9992ced24685d36947523b20ef70ac91c0aba2164736f6c63430008190033",
}

// FileRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FileRegistryMetaData.ABI instead.
var FileRegistryABI = FileRegistryMetaData.ABI

// FileRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FileRegistryMetaData.Bin instead.
var FileRegistryBin = FileRegistryMetaData.Bin

// DeployFileRegistry deploys a new Ethereum contract, binding an instance of FileRegistry to it.
func DeployFileRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FileRegistry, error) {
	parsed, err := FileRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FileRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileRegistry{FileRegistryCaller: FileRegistryCaller{contract: contract}, FileRegistryTransactor: FileRegistryTransactor{contract: contract}, FileRegistryFilterer: FileRegistryFilterer{contract: contract}}, nil
}

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

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_FileRegistry *FileRegistryCaller) Get(opts *bind.CallOpts, filePath string) (string, error) {
	var out []interface{}
	err := _FileRegistry.contract.Call(opts, &out, "get", filePath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_FileRegistry *FileRegistrySession) Get(filePath string) (string, error) {
	return _FileRegistry.Contract.Get(&_FileRegistry.CallOpts, filePath)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_FileRegistry *FileRegistryCallerSession) Get(filePath string) (string, error) {
	return _FileRegistry.Contract.Get(&_FileRegistry.CallOpts, filePath)
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
