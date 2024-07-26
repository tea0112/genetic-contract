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

// ControllerMovie is an auto generated low-level Go binding around an user-defined struct.
type ControllerMovie struct {
	Id     *big.Int
	Title  string
	Length *big.Int
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"movieId\",\"type\":\"uint256\"}],\"name\":\"CreatingMovieEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"createNewMovie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteMovie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMovie\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"internalType\":\"structController.Movie\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"movies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// GetMovie is a free data retrieval call binding the contract method 0x179ffe76.
//
// Solidity: function getMovie(uint256 _id) view returns((uint256,string,uint256))
func (_Controller *ControllerCaller) GetMovie(opts *bind.CallOpts, _id *big.Int) (ControllerMovie, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getMovie", _id)

	if err != nil {
		return *new(ControllerMovie), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerMovie)).(*ControllerMovie)

	return out0, err

}

// GetMovie is a free data retrieval call binding the contract method 0x179ffe76.
//
// Solidity: function getMovie(uint256 _id) view returns((uint256,string,uint256))
func (_Controller *ControllerSession) GetMovie(_id *big.Int) (ControllerMovie, error) {
	return _Controller.Contract.GetMovie(&_Controller.CallOpts, _id)
}

// GetMovie is a free data retrieval call binding the contract method 0x179ffe76.
//
// Solidity: function getMovie(uint256 _id) view returns((uint256,string,uint256))
func (_Controller *ControllerCallerSession) GetMovie(_id *big.Int) (ControllerMovie, error) {
	return _Controller.Contract.GetMovie(&_Controller.CallOpts, _id)
}

// Movies is a free data retrieval call binding the contract method 0x404ebf6c.
//
// Solidity: function movies(uint256 ) view returns(uint256 id, string title, uint256 length)
func (_Controller *ControllerCaller) Movies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Title  string
	Length *big.Int
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "movies", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Title  string
		Length *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Length = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Movies is a free data retrieval call binding the contract method 0x404ebf6c.
//
// Solidity: function movies(uint256 ) view returns(uint256 id, string title, uint256 length)
func (_Controller *ControllerSession) Movies(arg0 *big.Int) (struct {
	Id     *big.Int
	Title  string
	Length *big.Int
}, error) {
	return _Controller.Contract.Movies(&_Controller.CallOpts, arg0)
}

// Movies is a free data retrieval call binding the contract method 0x404ebf6c.
//
// Solidity: function movies(uint256 ) view returns(uint256 id, string title, uint256 length)
func (_Controller *ControllerCallerSession) Movies(arg0 *big.Int) (struct {
	Id     *big.Int
	Title  string
	Length *big.Int
}, error) {
	return _Controller.Contract.Movies(&_Controller.CallOpts, arg0)
}

// CreateNewMovie is a paid mutator transaction binding the contract method 0x90416b68.
//
// Solidity: function createNewMovie(uint256 _id, string _title, uint256 _length) returns()
func (_Controller *ControllerTransactor) CreateNewMovie(opts *bind.TransactOpts, _id *big.Int, _title string, _length *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "createNewMovie", _id, _title, _length)
}

// CreateNewMovie is a paid mutator transaction binding the contract method 0x90416b68.
//
// Solidity: function createNewMovie(uint256 _id, string _title, uint256 _length) returns()
func (_Controller *ControllerSession) CreateNewMovie(_id *big.Int, _title string, _length *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.CreateNewMovie(&_Controller.TransactOpts, _id, _title, _length)
}

// CreateNewMovie is a paid mutator transaction binding the contract method 0x90416b68.
//
// Solidity: function createNewMovie(uint256 _id, string _title, uint256 _length) returns()
func (_Controller *ControllerTransactorSession) CreateNewMovie(_id *big.Int, _title string, _length *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.CreateNewMovie(&_Controller.TransactOpts, _id, _title, _length)
}

// DeleteMovie is a paid mutator transaction binding the contract method 0x6b4f6cb7.
//
// Solidity: function deleteMovie(uint256 _id) returns()
func (_Controller *ControllerTransactor) DeleteMovie(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "deleteMovie", _id)
}

// DeleteMovie is a paid mutator transaction binding the contract method 0x6b4f6cb7.
//
// Solidity: function deleteMovie(uint256 _id) returns()
func (_Controller *ControllerSession) DeleteMovie(_id *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.DeleteMovie(&_Controller.TransactOpts, _id)
}

// DeleteMovie is a paid mutator transaction binding the contract method 0x6b4f6cb7.
//
// Solidity: function deleteMovie(uint256 _id) returns()
func (_Controller *ControllerTransactorSession) DeleteMovie(_id *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.DeleteMovie(&_Controller.TransactOpts, _id)
}

// ControllerCreatingMovieEventIterator is returned from FilterCreatingMovieEvent and is used to iterate over the raw logs and unpacked data for CreatingMovieEvent events raised by the Controller contract.
type ControllerCreatingMovieEventIterator struct {
	Event *ControllerCreatingMovieEvent // Event containing the contract specifics and raw log

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
func (it *ControllerCreatingMovieEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerCreatingMovieEvent)
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
		it.Event = new(ControllerCreatingMovieEvent)
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
func (it *ControllerCreatingMovieEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerCreatingMovieEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerCreatingMovieEvent represents a CreatingMovieEvent event raised by the Controller contract.
type ControllerCreatingMovieEvent struct {
	Message string
	MovieId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatingMovieEvent is a free log retrieval operation binding the contract event 0xf0ceadded3dfba7a0089b8fd585b7d118e821ecdfda2f54dff3ae8d666afc602.
//
// Solidity: event CreatingMovieEvent(string message, uint256 movieId)
func (_Controller *ControllerFilterer) FilterCreatingMovieEvent(opts *bind.FilterOpts) (*ControllerCreatingMovieEventIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "CreatingMovieEvent")
	if err != nil {
		return nil, err
	}
	return &ControllerCreatingMovieEventIterator{contract: _Controller.contract, event: "CreatingMovieEvent", logs: logs, sub: sub}, nil
}

// WatchCreatingMovieEvent is a free log subscription operation binding the contract event 0xf0ceadded3dfba7a0089b8fd585b7d118e821ecdfda2f54dff3ae8d666afc602.
//
// Solidity: event CreatingMovieEvent(string message, uint256 movieId)
func (_Controller *ControllerFilterer) WatchCreatingMovieEvent(opts *bind.WatchOpts, sink chan<- *ControllerCreatingMovieEvent) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "CreatingMovieEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerCreatingMovieEvent)
				if err := _Controller.contract.UnpackLog(event, "CreatingMovieEvent", log); err != nil {
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

// ParseCreatingMovieEvent is a log parse operation binding the contract event 0xf0ceadded3dfba7a0089b8fd585b7d118e821ecdfda2f54dff3ae8d666afc602.
//
// Solidity: event CreatingMovieEvent(string message, uint256 movieId)
func (_Controller *ControllerFilterer) ParseCreatingMovieEvent(log types.Log) (*ControllerCreatingMovieEvent, error) {
	event := new(ControllerCreatingMovieEvent)
	if err := _Controller.contract.UnpackLog(event, "CreatingMovieEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
