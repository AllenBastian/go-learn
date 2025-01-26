// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// HashStoreRevokeResult is an auto generated low-level Go binding around an user-defined struct.
type HashStoreRevokeResult struct {
	Hash string
	Stat bool
}

// HashStoreStoreResult is an auto generated low-level Go binding around an user-defined struct.
type HashStoreStoreResult struct {
	Hash   string
	Stored bool
}

// HashStoreVerificationResult is an auto generated low-level Go binding around an user-defined struct.
type HashStoreVerificationResult struct {
	VerifiedHashes      string
	VerificationResults bool
	Revoked             bool
	MerkleRoot          string
}

// HashStoreallMerkleRoot is an auto generated low-level Go binding around an user-defined struct.
type HashStoreallMerkleRoot struct {
	Roots string
}

// HashStoreMetaData contains all meta data concerning the HashStore contract.
var HashStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rev\",\"type\":\"bool\"}],\"name\":\"HashRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"HashStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"merkleroot\",\"type\":\"string\"}],\"name\":\"finaliseBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMerkleRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"roots\",\"type\":\"string\"}],\"internalType\":\"structHashStore.allMerkleRoot[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"getBatchById\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHashesForNextBatch\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"hashesToRevoke\",\"type\":\"string[]\"}],\"name\":\"revokeHashes\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"stat\",\"type\":\"bool\"}],\"internalType\":\"structHashStore.RevokeResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"hashes\",\"type\":\"string[]\"}],\"name\":\"storeHash\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"Hash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"stored\",\"type\":\"bool\"}],\"internalType\":\"structHashStore.StoreResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"hashesToVerify\",\"type\":\"string[]\"}],\"name\":\"verifyHashesByValue\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"verifiedHashes\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"verificationResults\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"merkleRoot\",\"type\":\"string\"}],\"internalType\":\"structHashStore.VerificationResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HashStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use HashStoreMetaData.ABI instead.
var HashStoreABI = HashStoreMetaData.ABI

// HashStore is an auto generated Go binding around an Ethereum contract.
type HashStore struct {
	HashStoreCaller     // Read-only binding to the contract
	HashStoreTransactor // Write-only binding to the contract
	HashStoreFilterer   // Log filterer for contract events
}

// HashStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashStoreSession struct {
	Contract     *HashStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashStoreCallerSession struct {
	Contract *HashStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HashStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashStoreTransactorSession struct {
	Contract     *HashStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HashStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashStoreRaw struct {
	Contract *HashStore // Generic contract binding to access the raw methods on
}

// HashStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashStoreCallerRaw struct {
	Contract *HashStoreCaller // Generic read-only contract binding to access the raw methods on
}

// HashStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashStoreTransactorRaw struct {
	Contract *HashStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashStore creates a new instance of HashStore, bound to a specific deployed contract.
func NewHashStore(address common.Address, backend bind.ContractBackend) (*HashStore, error) {
	contract, err := bindHashStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashStore{HashStoreCaller: HashStoreCaller{contract: contract}, HashStoreTransactor: HashStoreTransactor{contract: contract}, HashStoreFilterer: HashStoreFilterer{contract: contract}}, nil
}

// NewHashStoreCaller creates a new read-only instance of HashStore, bound to a specific deployed contract.
func NewHashStoreCaller(address common.Address, caller bind.ContractCaller) (*HashStoreCaller, error) {
	contract, err := bindHashStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashStoreCaller{contract: contract}, nil
}

// NewHashStoreTransactor creates a new write-only instance of HashStore, bound to a specific deployed contract.
func NewHashStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*HashStoreTransactor, error) {
	contract, err := bindHashStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashStoreTransactor{contract: contract}, nil
}

// NewHashStoreFilterer creates a new log filterer instance of HashStore, bound to a specific deployed contract.
func NewHashStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*HashStoreFilterer, error) {
	contract, err := bindHashStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashStoreFilterer{contract: contract}, nil
}

// bindHashStore binds a generic wrapper to an already deployed contract.
func bindHashStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashStore *HashStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashStore.Contract.HashStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashStore *HashStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashStore.Contract.HashStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashStore *HashStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashStore.Contract.HashStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashStore *HashStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashStore *HashStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashStore *HashStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashStore.Contract.contract.Transact(opts, method, params...)
}

// GetAllMerkleRoot is a free data retrieval call binding the contract method 0xf4c184d7.
//
// Solidity: function getAllMerkleRoot() view returns((string)[])
func (_HashStore *HashStoreCaller) GetAllMerkleRoot(opts *bind.CallOpts) ([]HashStoreallMerkleRoot, error) {
	var out []interface{}
	err := _HashStore.contract.Call(opts, &out, "getAllMerkleRoot")

	if err != nil {
		return *new([]HashStoreallMerkleRoot), err
	}

	out0 := *abi.ConvertType(out[0], new([]HashStoreallMerkleRoot)).(*[]HashStoreallMerkleRoot)

	return out0, err

}

// GetAllMerkleRoot is a free data retrieval call binding the contract method 0xf4c184d7.
//
// Solidity: function getAllMerkleRoot() view returns((string)[])
func (_HashStore *HashStoreSession) GetAllMerkleRoot() ([]HashStoreallMerkleRoot, error) {
	return _HashStore.Contract.GetAllMerkleRoot(&_HashStore.CallOpts)
}

// GetAllMerkleRoot is a free data retrieval call binding the contract method 0xf4c184d7.
//
// Solidity: function getAllMerkleRoot() view returns((string)[])
func (_HashStore *HashStoreCallerSession) GetAllMerkleRoot() ([]HashStoreallMerkleRoot, error) {
	return _HashStore.Contract.GetAllMerkleRoot(&_HashStore.CallOpts)
}

// GetBatchById is a free data retrieval call binding the contract method 0x5d5e1205.
//
// Solidity: function getBatchById(uint256 batchId) view returns(string[], string, uint256)
func (_HashStore *HashStoreCaller) GetBatchById(opts *bind.CallOpts, batchId *big.Int) ([]string, string, *big.Int, error) {
	var out []interface{}
	err := _HashStore.contract.Call(opts, &out, "getBatchById", batchId)

	if err != nil {
		return *new([]string), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetBatchById is a free data retrieval call binding the contract method 0x5d5e1205.
//
// Solidity: function getBatchById(uint256 batchId) view returns(string[], string, uint256)
func (_HashStore *HashStoreSession) GetBatchById(batchId *big.Int) ([]string, string, *big.Int, error) {
	return _HashStore.Contract.GetBatchById(&_HashStore.CallOpts, batchId)
}

// GetBatchById is a free data retrieval call binding the contract method 0x5d5e1205.
//
// Solidity: function getBatchById(uint256 batchId) view returns(string[], string, uint256)
func (_HashStore *HashStoreCallerSession) GetBatchById(batchId *big.Int) ([]string, string, *big.Int, error) {
	return _HashStore.Contract.GetBatchById(&_HashStore.CallOpts, batchId)
}

// GetContractDetails is a free data retrieval call binding the contract method 0x9dfc9117.
//
// Solidity: function getContractDetails() view returns(uint256, uint256, uint256)
func (_HashStore *HashStoreCaller) GetContractDetails(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _HashStore.contract.Call(opts, &out, "getContractDetails")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetContractDetails is a free data retrieval call binding the contract method 0x9dfc9117.
//
// Solidity: function getContractDetails() view returns(uint256, uint256, uint256)
func (_HashStore *HashStoreSession) GetContractDetails() (*big.Int, *big.Int, *big.Int, error) {
	return _HashStore.Contract.GetContractDetails(&_HashStore.CallOpts)
}

// GetContractDetails is a free data retrieval call binding the contract method 0x9dfc9117.
//
// Solidity: function getContractDetails() view returns(uint256, uint256, uint256)
func (_HashStore *HashStoreCallerSession) GetContractDetails() (*big.Int, *big.Int, *big.Int, error) {
	return _HashStore.Contract.GetContractDetails(&_HashStore.CallOpts)
}

// GetHashesForNextBatch is a free data retrieval call binding the contract method 0x2dfbc292.
//
// Solidity: function getHashesForNextBatch() view returns(string[])
func (_HashStore *HashStoreCaller) GetHashesForNextBatch(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _HashStore.contract.Call(opts, &out, "getHashesForNextBatch")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetHashesForNextBatch is a free data retrieval call binding the contract method 0x2dfbc292.
//
// Solidity: function getHashesForNextBatch() view returns(string[])
func (_HashStore *HashStoreSession) GetHashesForNextBatch() ([]string, error) {
	return _HashStore.Contract.GetHashesForNextBatch(&_HashStore.CallOpts)
}

// GetHashesForNextBatch is a free data retrieval call binding the contract method 0x2dfbc292.
//
// Solidity: function getHashesForNextBatch() view returns(string[])
func (_HashStore *HashStoreCallerSession) GetHashesForNextBatch() ([]string, error) {
	return _HashStore.Contract.GetHashesForNextBatch(&_HashStore.CallOpts)
}

// VerifyHashesByValue is a free data retrieval call binding the contract method 0x922f0f33.
//
// Solidity: function verifyHashesByValue(string[] hashesToVerify) view returns((string,bool,bool,string)[])
func (_HashStore *HashStoreCaller) VerifyHashesByValue(opts *bind.CallOpts, hashesToVerify []string) ([]HashStoreVerificationResult, error) {
	var out []interface{}
	err := _HashStore.contract.Call(opts, &out, "verifyHashesByValue", hashesToVerify)

	if err != nil {
		return *new([]HashStoreVerificationResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]HashStoreVerificationResult)).(*[]HashStoreVerificationResult)

	return out0, err

}

// VerifyHashesByValue is a free data retrieval call binding the contract method 0x922f0f33.
//
// Solidity: function verifyHashesByValue(string[] hashesToVerify) view returns((string,bool,bool,string)[])
func (_HashStore *HashStoreSession) VerifyHashesByValue(hashesToVerify []string) ([]HashStoreVerificationResult, error) {
	return _HashStore.Contract.VerifyHashesByValue(&_HashStore.CallOpts, hashesToVerify)
}

// VerifyHashesByValue is a free data retrieval call binding the contract method 0x922f0f33.
//
// Solidity: function verifyHashesByValue(string[] hashesToVerify) view returns((string,bool,bool,string)[])
func (_HashStore *HashStoreCallerSession) VerifyHashesByValue(hashesToVerify []string) ([]HashStoreVerificationResult, error) {
	return _HashStore.Contract.VerifyHashesByValue(&_HashStore.CallOpts, hashesToVerify)
}

// FinaliseBatch is a paid mutator transaction binding the contract method 0xefe05596.
//
// Solidity: function finaliseBatch(string merkleroot) returns(uint256, bool)
func (_HashStore *HashStoreTransactor) FinaliseBatch(opts *bind.TransactOpts, merkleroot string) (*types.Transaction, error) {
	return _HashStore.contract.Transact(opts, "finaliseBatch", merkleroot)
}

// FinaliseBatch is a paid mutator transaction binding the contract method 0xefe05596.
//
// Solidity: function finaliseBatch(string merkleroot) returns(uint256, bool)
func (_HashStore *HashStoreSession) FinaliseBatch(merkleroot string) (*types.Transaction, error) {
	return _HashStore.Contract.FinaliseBatch(&_HashStore.TransactOpts, merkleroot)
}

// FinaliseBatch is a paid mutator transaction binding the contract method 0xefe05596.
//
// Solidity: function finaliseBatch(string merkleroot) returns(uint256, bool)
func (_HashStore *HashStoreTransactorSession) FinaliseBatch(merkleroot string) (*types.Transaction, error) {
	return _HashStore.Contract.FinaliseBatch(&_HashStore.TransactOpts, merkleroot)
}

// RevokeHashes is a paid mutator transaction binding the contract method 0xfd40e7eb.
//
// Solidity: function revokeHashes(string[] hashesToRevoke) returns((string,bool)[])
func (_HashStore *HashStoreTransactor) RevokeHashes(opts *bind.TransactOpts, hashesToRevoke []string) (*types.Transaction, error) {
	return _HashStore.contract.Transact(opts, "revokeHashes", hashesToRevoke)
}

// RevokeHashes is a paid mutator transaction binding the contract method 0xfd40e7eb.
//
// Solidity: function revokeHashes(string[] hashesToRevoke) returns((string,bool)[])
func (_HashStore *HashStoreSession) RevokeHashes(hashesToRevoke []string) (*types.Transaction, error) {
	return _HashStore.Contract.RevokeHashes(&_HashStore.TransactOpts, hashesToRevoke)
}

// RevokeHashes is a paid mutator transaction binding the contract method 0xfd40e7eb.
//
// Solidity: function revokeHashes(string[] hashesToRevoke) returns((string,bool)[])
func (_HashStore *HashStoreTransactorSession) RevokeHashes(hashesToRevoke []string) (*types.Transaction, error) {
	return _HashStore.Contract.RevokeHashes(&_HashStore.TransactOpts, hashesToRevoke)
}

// StoreHash is a paid mutator transaction binding the contract method 0x95df4b40.
//
// Solidity: function storeHash(string[] hashes) returns((string,bool)[])
func (_HashStore *HashStoreTransactor) StoreHash(opts *bind.TransactOpts, hashes []string) (*types.Transaction, error) {
	return _HashStore.contract.Transact(opts, "storeHash", hashes)
}

// StoreHash is a paid mutator transaction binding the contract method 0x95df4b40.
//
// Solidity: function storeHash(string[] hashes) returns((string,bool)[])
func (_HashStore *HashStoreSession) StoreHash(hashes []string) (*types.Transaction, error) {
	return _HashStore.Contract.StoreHash(&_HashStore.TransactOpts, hashes)
}

// StoreHash is a paid mutator transaction binding the contract method 0x95df4b40.
//
// Solidity: function storeHash(string[] hashes) returns((string,bool)[])
func (_HashStore *HashStoreTransactorSession) StoreHash(hashes []string) (*types.Transaction, error) {
	return _HashStore.Contract.StoreHash(&_HashStore.TransactOpts, hashes)
}

// HashStoreHashRevokedIterator is returned from FilterHashRevoked and is used to iterate over the raw logs and unpacked data for HashRevoked events raised by the HashStore contract.
type HashStoreHashRevokedIterator struct {
	Event *HashStoreHashRevoked // Event containing the contract specifics and raw log

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
func (it *HashStoreHashRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashStoreHashRevoked)
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
		it.Event = new(HashStoreHashRevoked)
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
func (it *HashStoreHashRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashStoreHashRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashStoreHashRevoked represents a HashRevoked event raised by the HashStore contract.
type HashStoreHashRevoked struct {
	Hash string
	Rev  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHashRevoked is a free log retrieval operation binding the contract event 0x96182ebc17aea74313faef017a57a7922f36576ba2e39d4b0072938d2c024daf.
//
// Solidity: event HashRevoked(string hash, bool rev)
func (_HashStore *HashStoreFilterer) FilterHashRevoked(opts *bind.FilterOpts) (*HashStoreHashRevokedIterator, error) {

	logs, sub, err := _HashStore.contract.FilterLogs(opts, "HashRevoked")
	if err != nil {
		return nil, err
	}
	return &HashStoreHashRevokedIterator{contract: _HashStore.contract, event: "HashRevoked", logs: logs, sub: sub}, nil
}

// WatchHashRevoked is a free log subscription operation binding the contract event 0x96182ebc17aea74313faef017a57a7922f36576ba2e39d4b0072938d2c024daf.
//
// Solidity: event HashRevoked(string hash, bool rev)
func (_HashStore *HashStoreFilterer) WatchHashRevoked(opts *bind.WatchOpts, sink chan<- *HashStoreHashRevoked) (event.Subscription, error) {

	logs, sub, err := _HashStore.contract.WatchLogs(opts, "HashRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashStoreHashRevoked)
				if err := _HashStore.contract.UnpackLog(event, "HashRevoked", log); err != nil {
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

// ParseHashRevoked is a log parse operation binding the contract event 0x96182ebc17aea74313faef017a57a7922f36576ba2e39d4b0072938d2c024daf.
//
// Solidity: event HashRevoked(string hash, bool rev)
func (_HashStore *HashStoreFilterer) ParseHashRevoked(log types.Log) (*HashStoreHashRevoked, error) {
	event := new(HashStoreHashRevoked)
	if err := _HashStore.contract.UnpackLog(event, "HashRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashStoreHashStoredIterator is returned from FilterHashStored and is used to iterate over the raw logs and unpacked data for HashStored events raised by the HashStore contract.
type HashStoreHashStoredIterator struct {
	Event *HashStoreHashStored // Event containing the contract specifics and raw log

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
func (it *HashStoreHashStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashStoreHashStored)
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
		it.Event = new(HashStoreHashStored)
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
func (it *HashStoreHashStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashStoreHashStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashStoreHashStored represents a HashStored event raised by the HashStore contract.
type HashStoreHashStored struct {
	BatchId *big.Int
	Id      *big.Int
	Hash    string
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHashStored is a free log retrieval operation binding the contract event 0x6212de8bf38af332b5e4dbcdb6cfb305a09e58d2f7171cb78daff28612f7e067.
//
// Solidity: event HashStored(uint256 indexed batchId, uint256 indexed id, string hash, string message)
func (_HashStore *HashStoreFilterer) FilterHashStored(opts *bind.FilterOpts, batchId []*big.Int, id []*big.Int) (*HashStoreHashStoredIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _HashStore.contract.FilterLogs(opts, "HashStored", batchIdRule, idRule)
	if err != nil {
		return nil, err
	}
	return &HashStoreHashStoredIterator{contract: _HashStore.contract, event: "HashStored", logs: logs, sub: sub}, nil
}

// WatchHashStored is a free log subscription operation binding the contract event 0x6212de8bf38af332b5e4dbcdb6cfb305a09e58d2f7171cb78daff28612f7e067.
//
// Solidity: event HashStored(uint256 indexed batchId, uint256 indexed id, string hash, string message)
func (_HashStore *HashStoreFilterer) WatchHashStored(opts *bind.WatchOpts, sink chan<- *HashStoreHashStored, batchId []*big.Int, id []*big.Int) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _HashStore.contract.WatchLogs(opts, "HashStored", batchIdRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashStoreHashStored)
				if err := _HashStore.contract.UnpackLog(event, "HashStored", log); err != nil {
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

// ParseHashStored is a log parse operation binding the contract event 0x6212de8bf38af332b5e4dbcdb6cfb305a09e58d2f7171cb78daff28612f7e067.
//
// Solidity: event HashStored(uint256 indexed batchId, uint256 indexed id, string hash, string message)
func (_HashStore *HashStoreFilterer) ParseHashStored(log types.Log) (*HashStoreHashStored, error) {
	event := new(HashStoreHashStored)
	if err := _HashStore.contract.UnpackLog(event, "HashStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
