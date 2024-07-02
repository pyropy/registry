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
	Bin: "0x6080604052348015600e575f80fd5b5061087c8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80634909b29b14610043578063693ec85e14610073578063962939b8146100a3575b5f80fd5b61005d600480360381019061005891906103e7565b6100bf565b60405161006a919061048e565b60405180910390f35b61008d600480360381019061008891906103e7565b610171565b60405161009a919061048e565b60405180910390f35b6100bd60048036038101906100b891906104ae565b61021e565b005b5f818051602081018201805184825260208301602085012081835280955050505050505f9150905080546100f290610551565b80601f016020809104026020016040519081016040528092919081815260200182805461011e90610551565b80156101695780601f1061014057610100808354040283529160200191610169565b820191905f5260205f20905b81548152906001019060200180831161014c57829003601f168201915b505050505081565b60605f8260405161018291906105bb565b9081526020016040518091039020805461019b90610551565b80601f01602080910402602001604051908101604052809291908181526020018280546101c790610551565b80156102125780601f106101e957610100808354040283529160200191610212565b820191905f5260205f20905b8154815290600101906020018083116101f557829003601f168201915b50505050509050919050565b805f8360405161022e91906105bb565b908152602001604051809103902090816102489190610777565b508160405161025791906105bb565b60405180910390207f523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef38260405161028e919061048e565b60405180910390a25050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102f9826102b3565b810181811067ffffffffffffffff82111715610318576103176102c3565b5b80604052505050565b5f61032a61029a565b905061033682826102f0565b919050565b5f67ffffffffffffffff821115610355576103546102c3565b5b61035e826102b3565b9050602081019050919050565b828183375f83830152505050565b5f61038b6103868461033b565b610321565b9050828152602081018484840111156103a7576103a66102af565b5b6103b284828561036b565b509392505050565b5f82601f8301126103ce576103cd6102ab565b5b81356103de848260208601610379565b91505092915050565b5f602082840312156103fc576103fb6102a3565b5b5f82013567ffffffffffffffff811115610419576104186102a7565b5b610425848285016103ba565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6104608261042e565b61046a8185610438565b935061047a818560208601610448565b610483816102b3565b840191505092915050565b5f6020820190508181035f8301526104a68184610456565b905092915050565b5f80604083850312156104c4576104c36102a3565b5b5f83013567ffffffffffffffff8111156104e1576104e06102a7565b5b6104ed858286016103ba565b925050602083013567ffffffffffffffff81111561050e5761050d6102a7565b5b61051a858286016103ba565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061056857607f821691505b60208210810361057b5761057a610524565b5b50919050565b5f81905092915050565b5f6105958261042e565b61059f8185610581565b93506105af818560208601610448565b80840191505092915050565b5f6105c6828461058b565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261062d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826105f2565b61063786836105f2565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61067b6106766106718461064f565b610658565b61064f565b9050919050565b5f819050919050565b61069483610661565b6106a86106a082610682565b8484546105fe565b825550505050565b5f90565b6106bc6106b0565b6106c781848461068b565b505050565b5b818110156106ea576106df5f826106b4565b6001810190506106cd565b5050565b601f82111561072f57610700816105d1565b610709846105e3565b81016020851015610718578190505b61072c610724856105e3565b8301826106cc565b50505b505050565b5f82821c905092915050565b5f61074f5f1984600802610734565b1980831691505092915050565b5f6107678383610740565b9150826002028217905092915050565b6107808261042e565b67ffffffffffffffff811115610799576107986102c3565b5b6107a38254610551565b6107ae8282856106ee565b5f60209050601f8311600181146107df575f84156107cd578287015190505b6107d7858261075c565b86555061083e565b601f1984166107ed866105d1565b5f5b82811015610814578489015182556001820191506020850194506020810190506107ef565b86831015610831578489015161082d601f891682610740565b8355505b6001600288020188555050505b50505050505056fea264697066735822122005fbf24b85bf4b7e6d7772e93e7192d7fa6e5d0bd003a83b9ff941be12ee0c4464736f6c63430008190033",
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
