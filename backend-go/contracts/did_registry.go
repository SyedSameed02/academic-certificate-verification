// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DIDRegistryMetaData contains all meta data concerning the DIDRegistry contract.
var DIDRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"IssuerActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"IssuerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"IssuerRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"activateIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"deactivateIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"getIssuerDID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"isValidIssuer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"registerIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DIDRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DIDRegistryMetaData.ABI instead.
var DIDRegistryABI = DIDRegistryMetaData.ABI

// DIDRegistry is an auto generated Go binding around an Ethereum contract.
type DIDRegistry struct {
	DIDRegistryCaller     // Read-only binding to the contract
	DIDRegistryTransactor // Write-only binding to the contract
	DIDRegistryFilterer   // Log filterer for contract events
}

// DIDRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDRegistrySession struct {
	Contract     *DIDRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDRegistryCallerSession struct {
	Contract *DIDRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DIDRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDRegistryTransactorSession struct {
	Contract     *DIDRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DIDRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDRegistryRaw struct {
	Contract *DIDRegistry // Generic contract binding to access the raw methods on
}

// DIDRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDRegistryCallerRaw struct {
	Contract *DIDRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DIDRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDRegistryTransactorRaw struct {
	Contract *DIDRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDRegistry creates a new instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistry(address common.Address, backend bind.ContractBackend) (*DIDRegistry, error) {
	contract, err := bindDIDRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDRegistry{DIDRegistryCaller: DIDRegistryCaller{contract: contract}, DIDRegistryTransactor: DIDRegistryTransactor{contract: contract}, DIDRegistryFilterer: DIDRegistryFilterer{contract: contract}}, nil
}

// NewDIDRegistryCaller creates a new read-only instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryCaller(address common.Address, caller bind.ContractCaller) (*DIDRegistryCaller, error) {
	contract, err := bindDIDRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryCaller{contract: contract}, nil
}

// NewDIDRegistryTransactor creates a new write-only instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDRegistryTransactor, error) {
	contract, err := bindDIDRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryTransactor{contract: contract}, nil
}

// NewDIDRegistryFilterer creates a new log filterer instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDRegistryFilterer, error) {
	contract, err := bindDIDRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryFilterer{contract: contract}, nil
}

// bindDIDRegistry binds a generic wrapper to an already deployed contract.
func bindDIDRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DIDRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDRegistry *DIDRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIDRegistry.Contract.DIDRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDRegistry *DIDRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DIDRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDRegistry *DIDRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DIDRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDRegistry *DIDRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIDRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDRegistry *DIDRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDRegistry *DIDRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetIssuerDID is a free data retrieval call binding the contract method 0xbda57cec.
//
// Solidity: function getIssuerDID(address issuer) view returns(string)
func (_DIDRegistry *DIDRegistryCaller) GetIssuerDID(opts *bind.CallOpts, issuer common.Address) (string, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "getIssuerDID", issuer)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIssuerDID is a free data retrieval call binding the contract method 0xbda57cec.
//
// Solidity: function getIssuerDID(address issuer) view returns(string)
func (_DIDRegistry *DIDRegistrySession) GetIssuerDID(issuer common.Address) (string, error) {
	return _DIDRegistry.Contract.GetIssuerDID(&_DIDRegistry.CallOpts, issuer)
}

// GetIssuerDID is a free data retrieval call binding the contract method 0xbda57cec.
//
// Solidity: function getIssuerDID(address issuer) view returns(string)
func (_DIDRegistry *DIDRegistryCallerSession) GetIssuerDID(issuer common.Address) (string, error) {
	return _DIDRegistry.Contract.GetIssuerDID(&_DIDRegistry.CallOpts, issuer)
}

// IsValidIssuer is a free data retrieval call binding the contract method 0x474f3c81.
//
// Solidity: function isValidIssuer(address issuer) view returns(bool)
func (_DIDRegistry *DIDRegistryCaller) IsValidIssuer(opts *bind.CallOpts, issuer common.Address) (bool, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "isValidIssuer", issuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidIssuer is a free data retrieval call binding the contract method 0x474f3c81.
//
// Solidity: function isValidIssuer(address issuer) view returns(bool)
func (_DIDRegistry *DIDRegistrySession) IsValidIssuer(issuer common.Address) (bool, error) {
	return _DIDRegistry.Contract.IsValidIssuer(&_DIDRegistry.CallOpts, issuer)
}

// IsValidIssuer is a free data retrieval call binding the contract method 0x474f3c81.
//
// Solidity: function isValidIssuer(address issuer) view returns(bool)
func (_DIDRegistry *DIDRegistryCallerSession) IsValidIssuer(issuer common.Address) (bool, error) {
	return _DIDRegistry.Contract.IsValidIssuer(&_DIDRegistry.CallOpts, issuer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDRegistry *DIDRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDRegistry *DIDRegistrySession) Owner() (common.Address, error) {
	return _DIDRegistry.Contract.Owner(&_DIDRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDRegistry *DIDRegistryCallerSession) Owner() (common.Address, error) {
	return _DIDRegistry.Contract.Owner(&_DIDRegistry.CallOpts)
}

// ActivateIssuer is a paid mutator transaction binding the contract method 0x4b86daab.
//
// Solidity: function activateIssuer(address issuer) returns()
func (_DIDRegistry *DIDRegistryTransactor) ActivateIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "activateIssuer", issuer)
}

// ActivateIssuer is a paid mutator transaction binding the contract method 0x4b86daab.
//
// Solidity: function activateIssuer(address issuer) returns()
func (_DIDRegistry *DIDRegistrySession) ActivateIssuer(issuer common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.ActivateIssuer(&_DIDRegistry.TransactOpts, issuer)
}

// ActivateIssuer is a paid mutator transaction binding the contract method 0x4b86daab.
//
// Solidity: function activateIssuer(address issuer) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) ActivateIssuer(issuer common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.ActivateIssuer(&_DIDRegistry.TransactOpts, issuer)
}

// DeactivateIssuer is a paid mutator transaction binding the contract method 0x7a009135.
//
// Solidity: function deactivateIssuer(address issuer) returns()
func (_DIDRegistry *DIDRegistryTransactor) DeactivateIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "deactivateIssuer", issuer)
}

// DeactivateIssuer is a paid mutator transaction binding the contract method 0x7a009135.
//
// Solidity: function deactivateIssuer(address issuer) returns()
func (_DIDRegistry *DIDRegistrySession) DeactivateIssuer(issuer common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DeactivateIssuer(&_DIDRegistry.TransactOpts, issuer)
}

// DeactivateIssuer is a paid mutator transaction binding the contract method 0x7a009135.
//
// Solidity: function deactivateIssuer(address issuer) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) DeactivateIssuer(issuer common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DeactivateIssuer(&_DIDRegistry.TransactOpts, issuer)
}

// RegisterIssuer is a paid mutator transaction binding the contract method 0x8dd0acb2.
//
// Solidity: function registerIssuer(address issuer, string did) returns()
func (_DIDRegistry *DIDRegistryTransactor) RegisterIssuer(opts *bind.TransactOpts, issuer common.Address, did string) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "registerIssuer", issuer, did)
}

// RegisterIssuer is a paid mutator transaction binding the contract method 0x8dd0acb2.
//
// Solidity: function registerIssuer(address issuer, string did) returns()
func (_DIDRegistry *DIDRegistrySession) RegisterIssuer(issuer common.Address, did string) (*types.Transaction, error) {
	return _DIDRegistry.Contract.RegisterIssuer(&_DIDRegistry.TransactOpts, issuer, did)
}

// RegisterIssuer is a paid mutator transaction binding the contract method 0x8dd0acb2.
//
// Solidity: function registerIssuer(address issuer, string did) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) RegisterIssuer(issuer common.Address, did string) (*types.Transaction, error) {
	return _DIDRegistry.Contract.RegisterIssuer(&_DIDRegistry.TransactOpts, issuer, did)
}

// DIDRegistryIssuerActivatedIterator is returned from FilterIssuerActivated and is used to iterate over the raw logs and unpacked data for IssuerActivated events raised by the DIDRegistry contract.
type DIDRegistryIssuerActivatedIterator struct {
	Event *DIDRegistryIssuerActivated // Event containing the contract specifics and raw log

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
func (it *DIDRegistryIssuerActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryIssuerActivated)
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
		it.Event = new(DIDRegistryIssuerActivated)
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
func (it *DIDRegistryIssuerActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryIssuerActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryIssuerActivated represents a IssuerActivated event raised by the DIDRegistry contract.
type DIDRegistryIssuerActivated struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerActivated is a free log retrieval operation binding the contract event 0xfeb937248dcd237f20a1c703b4ef9fab8a76046b455b8af63f0dc768201a5043.
//
// Solidity: event IssuerActivated(address indexed issuer)
func (_DIDRegistry *DIDRegistryFilterer) FilterIssuerActivated(opts *bind.FilterOpts, issuer []common.Address) (*DIDRegistryIssuerActivatedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "IssuerActivated", issuerRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryIssuerActivatedIterator{contract: _DIDRegistry.contract, event: "IssuerActivated", logs: logs, sub: sub}, nil
}

// WatchIssuerActivated is a free log subscription operation binding the contract event 0xfeb937248dcd237f20a1c703b4ef9fab8a76046b455b8af63f0dc768201a5043.
//
// Solidity: event IssuerActivated(address indexed issuer)
func (_DIDRegistry *DIDRegistryFilterer) WatchIssuerActivated(opts *bind.WatchOpts, sink chan<- *DIDRegistryIssuerActivated, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "IssuerActivated", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryIssuerActivated)
				if err := _DIDRegistry.contract.UnpackLog(event, "IssuerActivated", log); err != nil {
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

// ParseIssuerActivated is a log parse operation binding the contract event 0xfeb937248dcd237f20a1c703b4ef9fab8a76046b455b8af63f0dc768201a5043.
//
// Solidity: event IssuerActivated(address indexed issuer)
func (_DIDRegistry *DIDRegistryFilterer) ParseIssuerActivated(log types.Log) (*DIDRegistryIssuerActivated, error) {
	event := new(DIDRegistryIssuerActivated)
	if err := _DIDRegistry.contract.UnpackLog(event, "IssuerActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryIssuerDeactivatedIterator is returned from FilterIssuerDeactivated and is used to iterate over the raw logs and unpacked data for IssuerDeactivated events raised by the DIDRegistry contract.
type DIDRegistryIssuerDeactivatedIterator struct {
	Event *DIDRegistryIssuerDeactivated // Event containing the contract specifics and raw log

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
func (it *DIDRegistryIssuerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryIssuerDeactivated)
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
		it.Event = new(DIDRegistryIssuerDeactivated)
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
func (it *DIDRegistryIssuerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryIssuerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryIssuerDeactivated represents a IssuerDeactivated event raised by the DIDRegistry contract.
type DIDRegistryIssuerDeactivated struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerDeactivated is a free log retrieval operation binding the contract event 0xbd353cae5dce393e81582130369cd5e22da2ec7e943ffae3210a857b3a954fdf.
//
// Solidity: event IssuerDeactivated(address indexed issuer)
func (_DIDRegistry *DIDRegistryFilterer) FilterIssuerDeactivated(opts *bind.FilterOpts, issuer []common.Address) (*DIDRegistryIssuerDeactivatedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "IssuerDeactivated", issuerRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryIssuerDeactivatedIterator{contract: _DIDRegistry.contract, event: "IssuerDeactivated", logs: logs, sub: sub}, nil
}

// WatchIssuerDeactivated is a free log subscription operation binding the contract event 0xbd353cae5dce393e81582130369cd5e22da2ec7e943ffae3210a857b3a954fdf.
//
// Solidity: event IssuerDeactivated(address indexed issuer)
func (_DIDRegistry *DIDRegistryFilterer) WatchIssuerDeactivated(opts *bind.WatchOpts, sink chan<- *DIDRegistryIssuerDeactivated, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "IssuerDeactivated", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryIssuerDeactivated)
				if err := _DIDRegistry.contract.UnpackLog(event, "IssuerDeactivated", log); err != nil {
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

// ParseIssuerDeactivated is a log parse operation binding the contract event 0xbd353cae5dce393e81582130369cd5e22da2ec7e943ffae3210a857b3a954fdf.
//
// Solidity: event IssuerDeactivated(address indexed issuer)
func (_DIDRegistry *DIDRegistryFilterer) ParseIssuerDeactivated(log types.Log) (*DIDRegistryIssuerDeactivated, error) {
	event := new(DIDRegistryIssuerDeactivated)
	if err := _DIDRegistry.contract.UnpackLog(event, "IssuerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIDRegistryIssuerRegisteredIterator is returned from FilterIssuerRegistered and is used to iterate over the raw logs and unpacked data for IssuerRegistered events raised by the DIDRegistry contract.
type DIDRegistryIssuerRegisteredIterator struct {
	Event *DIDRegistryIssuerRegistered // Event containing the contract specifics and raw log

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
func (it *DIDRegistryIssuerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDRegistryIssuerRegistered)
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
		it.Event = new(DIDRegistryIssuerRegistered)
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
func (it *DIDRegistryIssuerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDRegistryIssuerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDRegistryIssuerRegistered represents a IssuerRegistered event raised by the DIDRegistry contract.
type DIDRegistryIssuerRegistered struct {
	Issuer common.Address
	Did    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerRegistered is a free log retrieval operation binding the contract event 0xc0e782e5ce31b3a8ed34a04353667fa795030402831742c8a4b938c2c96ef033.
//
// Solidity: event IssuerRegistered(address indexed issuer, string did)
func (_DIDRegistry *DIDRegistryFilterer) FilterIssuerRegistered(opts *bind.FilterOpts, issuer []common.Address) (*DIDRegistryIssuerRegisteredIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _DIDRegistry.contract.FilterLogs(opts, "IssuerRegistered", issuerRule)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryIssuerRegisteredIterator{contract: _DIDRegistry.contract, event: "IssuerRegistered", logs: logs, sub: sub}, nil
}

// WatchIssuerRegistered is a free log subscription operation binding the contract event 0xc0e782e5ce31b3a8ed34a04353667fa795030402831742c8a4b938c2c96ef033.
//
// Solidity: event IssuerRegistered(address indexed issuer, string did)
func (_DIDRegistry *DIDRegistryFilterer) WatchIssuerRegistered(opts *bind.WatchOpts, sink chan<- *DIDRegistryIssuerRegistered, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _DIDRegistry.contract.WatchLogs(opts, "IssuerRegistered", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDRegistryIssuerRegistered)
				if err := _DIDRegistry.contract.UnpackLog(event, "IssuerRegistered", log); err != nil {
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

// ParseIssuerRegistered is a log parse operation binding the contract event 0xc0e782e5ce31b3a8ed34a04353667fa795030402831742c8a4b938c2c96ef033.
//
// Solidity: event IssuerRegistered(address indexed issuer, string did)
func (_DIDRegistry *DIDRegistryFilterer) ParseIssuerRegistered(log types.Log) (*DIDRegistryIssuerRegistered, error) {
	event := new(DIDRegistryIssuerRegistered)
	if err := _DIDRegistry.contract.UnpackLog(event, "IssuerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
