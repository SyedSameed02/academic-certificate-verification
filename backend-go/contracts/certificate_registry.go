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

// CertificateRegistryMetaData contains all meta data concerning the CertificateRegistry contract.
var CertificateRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_didRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"CertificateIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"CertificateRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"didRegistry\",\"outputs\":[{\"internalType\":\"contractIDIDRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"getCertificate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"issueCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CertificateRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateRegistryMetaData.ABI instead.
var CertificateRegistryABI = CertificateRegistryMetaData.ABI

// CertificateRegistry is an auto generated Go binding around an Ethereum contract.
type CertificateRegistry struct {
	CertificateRegistryCaller     // Read-only binding to the contract
	CertificateRegistryTransactor // Write-only binding to the contract
	CertificateRegistryFilterer   // Log filterer for contract events
}

// CertificateRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateRegistrySession struct {
	Contract     *CertificateRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CertificateRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateRegistryCallerSession struct {
	Contract *CertificateRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CertificateRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateRegistryTransactorSession struct {
	Contract     *CertificateRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CertificateRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateRegistryRaw struct {
	Contract *CertificateRegistry // Generic contract binding to access the raw methods on
}

// CertificateRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateRegistryCallerRaw struct {
	Contract *CertificateRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateRegistryTransactorRaw struct {
	Contract *CertificateRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificateRegistry creates a new instance of CertificateRegistry, bound to a specific deployed contract.
func NewCertificateRegistry(address common.Address, backend bind.ContractBackend) (*CertificateRegistry, error) {
	contract, err := bindCertificateRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificateRegistry{CertificateRegistryCaller: CertificateRegistryCaller{contract: contract}, CertificateRegistryTransactor: CertificateRegistryTransactor{contract: contract}, CertificateRegistryFilterer: CertificateRegistryFilterer{contract: contract}}, nil
}

// NewCertificateRegistryCaller creates a new read-only instance of CertificateRegistry, bound to a specific deployed contract.
func NewCertificateRegistryCaller(address common.Address, caller bind.ContractCaller) (*CertificateRegistryCaller, error) {
	contract, err := bindCertificateRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateRegistryCaller{contract: contract}, nil
}

// NewCertificateRegistryTransactor creates a new write-only instance of CertificateRegistry, bound to a specific deployed contract.
func NewCertificateRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateRegistryTransactor, error) {
	contract, err := bindCertificateRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateRegistryTransactor{contract: contract}, nil
}

// NewCertificateRegistryFilterer creates a new log filterer instance of CertificateRegistry, bound to a specific deployed contract.
func NewCertificateRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateRegistryFilterer, error) {
	contract, err := bindCertificateRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateRegistryFilterer{contract: contract}, nil
}

// bindCertificateRegistry binds a generic wrapper to an already deployed contract.
func bindCertificateRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertificateRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateRegistry *CertificateRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateRegistry.Contract.CertificateRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateRegistry *CertificateRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.CertificateRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateRegistry *CertificateRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.CertificateRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateRegistry *CertificateRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateRegistry *CertificateRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateRegistry *CertificateRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.contract.Transact(opts, method, params...)
}

// DidRegistry is a free data retrieval call binding the contract method 0x577f9fb1.
//
// Solidity: function didRegistry() view returns(address)
func (_CertificateRegistry *CertificateRegistryCaller) DidRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CertificateRegistry.contract.Call(opts, &out, "didRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DidRegistry is a free data retrieval call binding the contract method 0x577f9fb1.
//
// Solidity: function didRegistry() view returns(address)
func (_CertificateRegistry *CertificateRegistrySession) DidRegistry() (common.Address, error) {
	return _CertificateRegistry.Contract.DidRegistry(&_CertificateRegistry.CallOpts)
}

// DidRegistry is a free data retrieval call binding the contract method 0x577f9fb1.
//
// Solidity: function didRegistry() view returns(address)
func (_CertificateRegistry *CertificateRegistryCallerSession) DidRegistry() (common.Address, error) {
	return _CertificateRegistry.Contract.DidRegistry(&_CertificateRegistry.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x38a699a4.
//
// Solidity: function exists(bytes32 certHash) view returns(bool)
func (_CertificateRegistry *CertificateRegistryCaller) Exists(opts *bind.CallOpts, certHash [32]byte) (bool, error) {
	var out []interface{}
	err := _CertificateRegistry.contract.Call(opts, &out, "exists", certHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x38a699a4.
//
// Solidity: function exists(bytes32 certHash) view returns(bool)
func (_CertificateRegistry *CertificateRegistrySession) Exists(certHash [32]byte) (bool, error) {
	return _CertificateRegistry.Contract.Exists(&_CertificateRegistry.CallOpts, certHash)
}

// Exists is a free data retrieval call binding the contract method 0x38a699a4.
//
// Solidity: function exists(bytes32 certHash) view returns(bool)
func (_CertificateRegistry *CertificateRegistryCallerSession) Exists(certHash [32]byte) (bool, error) {
	return _CertificateRegistry.Contract.Exists(&_CertificateRegistry.CallOpts, certHash)
}

// GetCertificate is a free data retrieval call binding the contract method 0xf333fe08.
//
// Solidity: function getCertificate(bytes32 certHash) view returns(address issuer, bool revoked, uint256 issuedAt)
func (_CertificateRegistry *CertificateRegistryCaller) GetCertificate(opts *bind.CallOpts, certHash [32]byte) (struct {
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}, error) {
	var out []interface{}
	err := _CertificateRegistry.contract.Call(opts, &out, "getCertificate", certHash)

	outstruct := new(struct {
		Issuer   common.Address
		Revoked  bool
		IssuedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Issuer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Revoked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.IssuedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCertificate is a free data retrieval call binding the contract method 0xf333fe08.
//
// Solidity: function getCertificate(bytes32 certHash) view returns(address issuer, bool revoked, uint256 issuedAt)
func (_CertificateRegistry *CertificateRegistrySession) GetCertificate(certHash [32]byte) (struct {
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}, error) {
	return _CertificateRegistry.Contract.GetCertificate(&_CertificateRegistry.CallOpts, certHash)
}

// GetCertificate is a free data retrieval call binding the contract method 0xf333fe08.
//
// Solidity: function getCertificate(bytes32 certHash) view returns(address issuer, bool revoked, uint256 issuedAt)
func (_CertificateRegistry *CertificateRegistryCallerSession) GetCertificate(certHash [32]byte) (struct {
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}, error) {
	return _CertificateRegistry.Contract.GetCertificate(&_CertificateRegistry.CallOpts, certHash)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xa1f81824.
//
// Solidity: function issueCertificate(bytes32 certHash) returns()
func (_CertificateRegistry *CertificateRegistryTransactor) IssueCertificate(opts *bind.TransactOpts, certHash [32]byte) (*types.Transaction, error) {
	return _CertificateRegistry.contract.Transact(opts, "issueCertificate", certHash)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xa1f81824.
//
// Solidity: function issueCertificate(bytes32 certHash) returns()
func (_CertificateRegistry *CertificateRegistrySession) IssueCertificate(certHash [32]byte) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.IssueCertificate(&_CertificateRegistry.TransactOpts, certHash)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xa1f81824.
//
// Solidity: function issueCertificate(bytes32 certHash) returns()
func (_CertificateRegistry *CertificateRegistryTransactorSession) IssueCertificate(certHash [32]byte) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.IssueCertificate(&_CertificateRegistry.TransactOpts, certHash)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 certHash) returns()
func (_CertificateRegistry *CertificateRegistryTransactor) RevokeCertificate(opts *bind.TransactOpts, certHash [32]byte) (*types.Transaction, error) {
	return _CertificateRegistry.contract.Transact(opts, "revokeCertificate", certHash)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 certHash) returns()
func (_CertificateRegistry *CertificateRegistrySession) RevokeCertificate(certHash [32]byte) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.RevokeCertificate(&_CertificateRegistry.TransactOpts, certHash)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 certHash) returns()
func (_CertificateRegistry *CertificateRegistryTransactorSession) RevokeCertificate(certHash [32]byte) (*types.Transaction, error) {
	return _CertificateRegistry.Contract.RevokeCertificate(&_CertificateRegistry.TransactOpts, certHash)
}

// CertificateRegistryCertificateIssuedIterator is returned from FilterCertificateIssued and is used to iterate over the raw logs and unpacked data for CertificateIssued events raised by the CertificateRegistry contract.
type CertificateRegistryCertificateIssuedIterator struct {
	Event *CertificateRegistryCertificateIssued // Event containing the contract specifics and raw log

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
func (it *CertificateRegistryCertificateIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateRegistryCertificateIssued)
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
		it.Event = new(CertificateRegistryCertificateIssued)
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
func (it *CertificateRegistryCertificateIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateRegistryCertificateIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateRegistryCertificateIssued represents a CertificateIssued event raised by the CertificateRegistry contract.
type CertificateRegistryCertificateIssued struct {
	CertHash [32]byte
	Issuer   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCertificateIssued is a free log retrieval operation binding the contract event 0x11fb9f283c08b0122fe14bf3b5692d38f07282003f545f4bbc8b766ddcbda832.
//
// Solidity: event CertificateIssued(bytes32 indexed certHash, address indexed issuer)
func (_CertificateRegistry *CertificateRegistryFilterer) FilterCertificateIssued(opts *bind.FilterOpts, certHash [][32]byte, issuer []common.Address) (*CertificateRegistryCertificateIssuedIterator, error) {

	var certHashRule []interface{}
	for _, certHashItem := range certHash {
		certHashRule = append(certHashRule, certHashItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _CertificateRegistry.contract.FilterLogs(opts, "CertificateIssued", certHashRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &CertificateRegistryCertificateIssuedIterator{contract: _CertificateRegistry.contract, event: "CertificateIssued", logs: logs, sub: sub}, nil
}

// WatchCertificateIssued is a free log subscription operation binding the contract event 0x11fb9f283c08b0122fe14bf3b5692d38f07282003f545f4bbc8b766ddcbda832.
//
// Solidity: event CertificateIssued(bytes32 indexed certHash, address indexed issuer)
func (_CertificateRegistry *CertificateRegistryFilterer) WatchCertificateIssued(opts *bind.WatchOpts, sink chan<- *CertificateRegistryCertificateIssued, certHash [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var certHashRule []interface{}
	for _, certHashItem := range certHash {
		certHashRule = append(certHashRule, certHashItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _CertificateRegistry.contract.WatchLogs(opts, "CertificateIssued", certHashRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateRegistryCertificateIssued)
				if err := _CertificateRegistry.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
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

// ParseCertificateIssued is a log parse operation binding the contract event 0x11fb9f283c08b0122fe14bf3b5692d38f07282003f545f4bbc8b766ddcbda832.
//
// Solidity: event CertificateIssued(bytes32 indexed certHash, address indexed issuer)
func (_CertificateRegistry *CertificateRegistryFilterer) ParseCertificateIssued(log types.Log) (*CertificateRegistryCertificateIssued, error) {
	event := new(CertificateRegistryCertificateIssued)
	if err := _CertificateRegistry.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateRegistryCertificateRevokedIterator is returned from FilterCertificateRevoked and is used to iterate over the raw logs and unpacked data for CertificateRevoked events raised by the CertificateRegistry contract.
type CertificateRegistryCertificateRevokedIterator struct {
	Event *CertificateRegistryCertificateRevoked // Event containing the contract specifics and raw log

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
func (it *CertificateRegistryCertificateRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateRegistryCertificateRevoked)
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
		it.Event = new(CertificateRegistryCertificateRevoked)
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
func (it *CertificateRegistryCertificateRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateRegistryCertificateRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateRegistryCertificateRevoked represents a CertificateRevoked event raised by the CertificateRegistry contract.
type CertificateRegistryCertificateRevoked struct {
	CertHash [32]byte
	Issuer   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCertificateRevoked is a free log retrieval operation binding the contract event 0x1dbd7f0554a6447214abb6405781d5ee119037a7024fcb7fffd542c3c12cd776.
//
// Solidity: event CertificateRevoked(bytes32 indexed certHash, address indexed issuer)
func (_CertificateRegistry *CertificateRegistryFilterer) FilterCertificateRevoked(opts *bind.FilterOpts, certHash [][32]byte, issuer []common.Address) (*CertificateRegistryCertificateRevokedIterator, error) {

	var certHashRule []interface{}
	for _, certHashItem := range certHash {
		certHashRule = append(certHashRule, certHashItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _CertificateRegistry.contract.FilterLogs(opts, "CertificateRevoked", certHashRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &CertificateRegistryCertificateRevokedIterator{contract: _CertificateRegistry.contract, event: "CertificateRevoked", logs: logs, sub: sub}, nil
}

// WatchCertificateRevoked is a free log subscription operation binding the contract event 0x1dbd7f0554a6447214abb6405781d5ee119037a7024fcb7fffd542c3c12cd776.
//
// Solidity: event CertificateRevoked(bytes32 indexed certHash, address indexed issuer)
func (_CertificateRegistry *CertificateRegistryFilterer) WatchCertificateRevoked(opts *bind.WatchOpts, sink chan<- *CertificateRegistryCertificateRevoked, certHash [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var certHashRule []interface{}
	for _, certHashItem := range certHash {
		certHashRule = append(certHashRule, certHashItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _CertificateRegistry.contract.WatchLogs(opts, "CertificateRevoked", certHashRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateRegistryCertificateRevoked)
				if err := _CertificateRegistry.contract.UnpackLog(event, "CertificateRevoked", log); err != nil {
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

// ParseCertificateRevoked is a log parse operation binding the contract event 0x1dbd7f0554a6447214abb6405781d5ee119037a7024fcb7fffd542c3c12cd776.
//
// Solidity: event CertificateRevoked(bytes32 indexed certHash, address indexed issuer)
func (_CertificateRegistry *CertificateRegistryFilterer) ParseCertificateRevoked(log types.Log) (*CertificateRegistryCertificateRevoked, error) {
	event := new(CertificateRegistryCertificateRevoked)
	if err := _CertificateRegistry.contract.UnpackLog(event, "CertificateRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
