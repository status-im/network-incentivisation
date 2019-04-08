// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodes

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NodesABI is the input ABI used to generate the binding from.
const NodesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"toggleRegistration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// NodesBin is the compiled bytecode used for deploying new contracts.
const NodesBin = `0x608060405234801561001057600080fd5b5060008054600160a01b60ff02196001600160a01b031990911633171690556107cd8061003e6000396000f3fe6080604052600436106100555760003560e01c80631c53c280146100625780634c164407146101015780636da49b83146101165780638994dd8e1461013d578063a0c15b77146101f0578063b2f10309146102a3575b361561006057600080fd5b005b34801561006e57600080fd5b5061008c6004803603602081101561008557600080fd5b50356102cf565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c65781810151838201526020016100ae565b50505050905090810190601f1680156100f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561010d57600080fd5b50610060610375565b34801561012257600080fd5b5061012b61039a565b60408051918252519081900360200190f35b34801561014957600080fd5b506100606004803603602081101561016057600080fd5b81019060208101813564010000000081111561017b57600080fd5b82018360208201111561018d57600080fd5b803590602001918460018302840111640100000000831117156101af57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103a1945050505050565b3480156101fc57600080fd5b506100606004803603602081101561021357600080fd5b81019060208101813564010000000081111561022e57600080fd5b82018360208201111561024057600080fd5b8035906020019184600183028401116401000000008311171561026257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610480945050505050565b3480156102af57600080fd5b50610060600480360360208110156102c657600080fd5b5035151561050a565b600181815481106102dc57fe5b600091825260209182902001805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529350909183018282801561036d5780601f106103425761010080835404028352916020019161036d565b820191906000526020600020905b81548152906001019060200180831161035057829003601f168201915b505050505081565b6000546001600160a01b0316331461038c57600080fd5b610398600160006105ec565b565b6001545b90565b600054600160a01b900460ff16806103c357506000546001600160a01b031633145b6103cc57600080fd5b60015460405182516002918491819060208401908083835b602083106104035780518252601f1990920191602091820191016103e4565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320939093555060018054808201808355600092909252845191935061047b927fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6909101919085019061060d565b505050565b6000546001600160a01b0316331461049757600080fd5b60006002826040518082805190602001908083835b602083106104cb5780518252601f1990920191602091820191016104ac565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205492506105069150829050610550565b5050565b6000546001600160a01b0316331461052157600080fd5b60008054911515600160a01b0274ff000000000000000000000000000000000000000019909216919091179055565b600154811061055e57600080fd5b60018054600019810190811061057057fe5b906000526020600020016001828154811061058757fe5b9060005260206000200190805460018160011615610100020316600290046105b092919061068b565b506001805460001981019081106105c357fe5b9060005260206000200160006105d99190610700565b6001805490610506906000198301610744565b508054600082559060005260206000209081019061060a9190610764565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061064e57805160ff191683800117855561067b565b8280016001018555821561067b579182015b8281111561067b578251825591602001919060010190610660565b50610687929150610787565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106c4578054855561067b565b8280016001018555821561067b57600052602060002091601f016020900482015b8281111561067b5782548255916001019190600101906106e5565b50805460018160011615610100020316600290046000825580601f10610726575061060a565b601f01602090049060005260206000209081019061060a9190610787565b81548183558181111561047b5760008381526020902061047b9181019083015b61039e91905b8082111561068757600061077e8282610700565b5060010161076a565b61039e91905b80821115610687576000815560010161078d56fea165627a7a723058206308c00ae9274cdce4d8800e10f29dd708bbb0f10e53f2286bc32e21d25315e10029`

// DeployNodes deploys a new Ethereum contract, binding an instance of Nodes to it.
func DeployNodes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nodes, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nodes{NodesCaller: NodesCaller{contract: contract}, NodesTransactor: NodesTransactor{contract: contract}, NodesFilterer: NodesFilterer{contract: contract}}, nil
}

// Nodes is an auto generated Go binding around an Ethereum contract.
type Nodes struct {
	NodesCaller     // Read-only binding to the contract
	NodesTransactor // Write-only binding to the contract
	NodesFilterer   // Log filterer for contract events
}

// NodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesSession struct {
	Contract     *Nodes            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesCallerSession struct {
	Contract *NodesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesTransactorSession struct {
	Contract     *NodesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodesRaw struct {
	Contract *Nodes // Generic contract binding to access the raw methods on
}

// NodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesCallerRaw struct {
	Contract *NodesCaller // Generic read-only contract binding to access the raw methods on
}

// NodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesTransactorRaw struct {
	Contract *NodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodes creates a new instance of Nodes, bound to a specific deployed contract.
func NewNodes(address common.Address, backend bind.ContractBackend) (*Nodes, error) {
	contract, err := bindNodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nodes{NodesCaller: NodesCaller{contract: contract}, NodesTransactor: NodesTransactor{contract: contract}, NodesFilterer: NodesFilterer{contract: contract}}, nil
}

// NewNodesCaller creates a new read-only instance of Nodes, bound to a specific deployed contract.
func NewNodesCaller(address common.Address, caller bind.ContractCaller) (*NodesCaller, error) {
	contract, err := bindNodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesCaller{contract: contract}, nil
}

// NewNodesTransactor creates a new write-only instance of Nodes, bound to a specific deployed contract.
func NewNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*NodesTransactor, error) {
	contract, err := bindNodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesTransactor{contract: contract}, nil
}

// NewNodesFilterer creates a new log filterer instance of Nodes, bound to a specific deployed contract.
func NewNodesFilterer(address common.Address, filterer bind.ContractFilterer) (*NodesFilterer, error) {
	contract, err := bindNodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesFilterer{contract: contract}, nil
}

// bindNodes binds a generic wrapper to an already deployed contract.
func bindNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodes *NodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nodes.Contract.NodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodes *NodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.Contract.NodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodes *NodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodes.Contract.NodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodes *NodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodes *NodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodes *NodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodes.Contract.contract.Transact(opts, method, params...)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() constant returns(uint256)
func (_Nodes *NodesCaller) NodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Nodes.contract.Call(opts, out, "nodeCount")
	return *ret0, err
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() constant returns(uint256)
func (_Nodes *NodesSession) NodeCount() (*big.Int, error) {
	return _Nodes.Contract.NodeCount(&_Nodes.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() constant returns(uint256)
func (_Nodes *NodesCallerSession) NodeCount() (*big.Int, error) {
	return _Nodes.Contract.NodeCount(&_Nodes.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(string)
func (_Nodes *NodesCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Nodes.contract.Call(opts, out, "nodes", arg0)
	return *ret0, err
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(string)
func (_Nodes *NodesSession) Nodes(arg0 *big.Int) (string, error) {
	return _Nodes.Contract.Nodes(&_Nodes.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(string)
func (_Nodes *NodesCallerSession) Nodes(arg0 *big.Int) (string, error) {
	return _Nodes.Contract.Nodes(&_Nodes.CallOpts, arg0)
}

// AddNode is a paid mutator transaction binding the contract method 0x8994dd8e.
//
// Solidity: function addNode(string _node) returns()
func (_Nodes *NodesTransactor) AddNode(opts *bind.TransactOpts, _node string) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "addNode", _node)
}

// AddNode is a paid mutator transaction binding the contract method 0x8994dd8e.
//
// Solidity: function addNode(string _node) returns()
func (_Nodes *NodesSession) AddNode(_node string) (*types.Transaction, error) {
	return _Nodes.Contract.AddNode(&_Nodes.TransactOpts, _node)
}

// AddNode is a paid mutator transaction binding the contract method 0x8994dd8e.
//
// Solidity: function addNode(string _node) returns()
func (_Nodes *NodesTransactorSession) AddNode(_node string) (*types.Transaction, error) {
	return _Nodes.Contract.AddNode(&_Nodes.TransactOpts, _node)
}

// DeleteAll is a paid mutator transaction binding the contract method 0x4c164407.
//
// Solidity: function deleteAll() returns()
func (_Nodes *NodesTransactor) DeleteAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "deleteAll")
}

// DeleteAll is a paid mutator transaction binding the contract method 0x4c164407.
//
// Solidity: function deleteAll() returns()
func (_Nodes *NodesSession) DeleteAll() (*types.Transaction, error) {
	return _Nodes.Contract.DeleteAll(&_Nodes.TransactOpts)
}

// DeleteAll is a paid mutator transaction binding the contract method 0x4c164407.
//
// Solidity: function deleteAll() returns()
func (_Nodes *NodesTransactorSession) DeleteAll() (*types.Transaction, error) {
	return _Nodes.Contract.DeleteAll(&_Nodes.TransactOpts)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(string _node) returns()
func (_Nodes *NodesTransactor) DeleteNode(opts *bind.TransactOpts, _node string) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "deleteNode", _node)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(string _node) returns()
func (_Nodes *NodesSession) DeleteNode(_node string) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteNode(&_Nodes.TransactOpts, _node)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(string _node) returns()
func (_Nodes *NodesTransactorSession) DeleteNode(_node string) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteNode(&_Nodes.TransactOpts, _node)
}

// ToggleRegistration is a paid mutator transaction binding the contract method 0xb2f10309.
//
// Solidity: function toggleRegistration(bool value) returns()
func (_Nodes *NodesTransactor) ToggleRegistration(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "toggleRegistration", value)
}

// ToggleRegistration is a paid mutator transaction binding the contract method 0xb2f10309.
//
// Solidity: function toggleRegistration(bool value) returns()
func (_Nodes *NodesSession) ToggleRegistration(value bool) (*types.Transaction, error) {
	return _Nodes.Contract.ToggleRegistration(&_Nodes.TransactOpts, value)
}

// ToggleRegistration is a paid mutator transaction binding the contract method 0xb2f10309.
//
// Solidity: function toggleRegistration(bool value) returns()
func (_Nodes *NodesTransactorSession) ToggleRegistration(value bool) (*types.Transaction, error) {
	return _Nodes.Contract.ToggleRegistration(&_Nodes.TransactOpts, value)
}
