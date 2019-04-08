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

// NodesV2ABI is the input ABI used to generate the binding from.
const NodesV2ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setOverrideBlockPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"publicKeyToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inactiveNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inactiveNodes\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"joinVotes\",\"type\":\"uint8\"},{\"name\":\"removeVotes\",\"type\":\"uint8\"},{\"name\":\"lastTimeHasVoted\",\"type\":\"uint256\"},{\"name\":\"lastTimeHasBeenVoted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getInactiveNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodes\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"joinVotes\",\"type\":\"uint8\"},{\"name\":\"removeVotes\",\"type\":\"uint8\"},{\"name\":\"lastTimeHasVoted\",\"type\":\"uint256\"},{\"name\":\"lastTimeHasBeenVoted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"joinNodes\",\"type\":\"address[]\"},{\"name\":\"removeNodes\",\"type\":\"address[]\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"addActiveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_blockPeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// NodesV2Bin is the compiled bytecode used for deploying new contracts.
const NodesV2Bin = `0x608060405234801561001057600080fd5b50604051602080611b5b8339810180604052602081101561003057600080fd5b5051600080546001600160a01b0319163317905543600655600555611b018061005a6000396000f3fe6080604052600436106100a75760003560e01c806372460fa81161006457806372460fa81461041d578063753408151461050957806393696e1a1461051e57806396f9d98314610548578063a19e39e814610572578063dad7bcee146106a2576100a7565b80631c19b3e3146100b457806343ae656c146100e05780634f0f4aa9146101ad5780635aca952e1461027157806363cd6e18146103365780636d1c76c2146103f6575b36156100b257600080fd5b005b3480156100c057600080fd5b506100b2600480360360208110156100d757600080fd5b50351515610762565b3480156100ec57600080fd5b506101916004803603602081101561010357600080fd5b810190602081018135600160201b81111561011d57600080fd5b82018360208201111561012f57600080fd5b803590602001918460018302840111600160201b8311171561015057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061078c945050505050565b604080516001600160a01b039092168252519081900360200190f35b3480156101b957600080fd5b506101d7600480360360208110156101d057600080fd5b503561079e565b60405180806020018463ffffffff1663ffffffff1681526020018361ffff1661ffff168152602001828103825285818151815260200191508051906020019080838360005b8381101561023457818101518382015260200161021c565b50505050905090810190601f1680156102615780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561027d57600080fd5b506103226004803603602081101561029457600080fd5b810190602081018135600160201b8111156102ae57600080fd5b8201836020820111156102c057600080fd5b803590602001918460018302840111600160201b831117156102e157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108da945050505050565b604080519115158252519081900360200190f35b34801561034257600080fd5b506100b26004803603606081101561035957600080fd5b810190602081018135600160201b81111561037357600080fd5b82018360208201111561038557600080fd5b803590602001918460018302840111600160201b831117156103a657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813563ffffffff169250506020013561ffff1661092d565b34801561040257600080fd5b5061040b610ad3565b60408051918252519081900360200190f35b34801561042957600080fd5b506104476004803603602081101561044057600080fd5b5035610ada565b60405180806020018863ffffffff1663ffffffff1681526020018761ffff1661ffff1681526020018660ff1660ff1681526020018560ff1660ff168152602001848152602001838152602001828103825289818151815260200191508051906020019080838360005b838110156104c85781810151838201526020016104b0565b50505050905090810190601f1680156104f55780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561051557600080fd5b5061040b610bc4565b34801561052a57600080fd5b506101d76004803603602081101561054157600080fd5b5035610bca565b34801561055457600080fd5b506104476004803603602081101561056b57600080fd5b5035610be4565b34801561057e57600080fd5b506100b26004803603604081101561059557600080fd5b810190602081018135600160201b8111156105af57600080fd5b8201836020820111156105c157600080fd5b803590602001918460208302840111600160201b831117156105e257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561063157600080fd5b82018360208201111561064357600080fd5b803590602001918460208302840111600160201b8311171561066457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610bf1945050505050565b3480156106ae57600080fd5b506100b2600480360360608110156106c557600080fd5b810190602081018135600160201b8111156106df57600080fd5b8201836020820111156106f157600080fd5b803590602001918460018302840111600160201b8311171561071257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813563ffffffff169250506020013561ffff16611051565b6000546001600160a01b0316331461077957600080fd5b6003805460ff1916911515919091179055565b8051602090910120606090811b901c90565b60606000806107ab611945565b600185815481106107b857fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e083018481529293909284929091849184018282801561085a5780601f1061082f5761010080835404028352916020019161085a565b820191906000526020600020905b81548152906001019060200180831161083d57829003601f168201915b5050509183525050600182015463ffffffff811660208084019190915261ffff600160201b83041660408085019190915260ff600160301b840481166060860152600160381b9093049092166080840152600284015460a084015260039093015460c0909201919091528251918301519201519097919650945092505050565b6000806108e68361078c565b6001600160a01b03811660009081526007602052604090205490915015801561092557506001600160a01b038116600090815260086020526040902054155b159392505050565b6109368361078c565b6001600160a01b0316336001600160a01b03161461095357600080fd5b3360009081526007602052604090205415801561097d575033600090815260086020526040902054155b61098657600080fd5b61098e611945565b506040805160e08101825284815263ffffffff841660208083019190915261ffff8416928201929092526000606082018190526080820181905260a0820181905260c08201819052600280546001810180835591909252825180519394919385936004027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0192610a23928492910190611981565b506020828101516001830180546040808701516060880151608089015163ffffffff1990941663ffffffff9096169590951765ffff000000001916600160201b61ffff909216919091021766ff0000000000001916600160301b60ff958616021767ff000000000000001916600160381b949092169390930217905560a084015160028085019190915560c090940151600390930192909255915433600090815260089093529120555050505050565b6002545b90565b60028181548110610ae757fe5b60009182526020918290206004919091020180546040805160026001841615610100026000190190931692909204601f810185900485028301850190915280825291935091839190830182828015610b805780601f10610b5557610100808354040283529160200191610b80565b820191906000526020600020905b815481529060010190602001808311610b6357829003601f168201915b5050505060018301546002840154600390940154929363ffffffff82169361ffff600160201b840416935060ff600160301b8404811693600160381b900416919087565b60015490565b6060600080610bd7611945565b600285815481106107b857fe5b60018181548110610ae757fe5b610bf961112c565b15610c0657610c06611139565b3360009081526007602052604090205480610c2057600080fd5b600654600180830381548110610c3257fe5b9060005260206000209060040201600201541415610c4f57600080fd5b600654600180830381548110610c6157fe5b600091825260208220600260049092020101919091555b8251811015610d7757600060076000858481518110610c9357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600014610d6e576000600180830381548110610cd757fe5b90600052602060002090600402019050600654816003015414610d1a57600654600382015560018101805467ff000000000000001916600160381b179055610d48565b6001808201805460ff600160381b80830482169094011690920267ff00000000000000199092169190911790555b6004546001820154600160381b900460ff161415610d6c57610d6c6001830361114e565b505b50600101610c78565b5060005b835181101561104b576000848281518110610d9257fe5b6020908102919091018101516001600160a01b03811660009081526008909252604090912054909150801561104157600060026001830381548110610dd357fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152919350610e7992849190830182828015610e6f5780601f10610e4457610100808354040283529160200191610e6f565b820191906000526020600020905b815481529060010190602001808311610e5257829003601f168201915b505050505061078c565b6001600160a01b0316836001600160a01b031614610e9657600080fd5b600654816003015414610ec857600654600382015560018101805466ff0000000000001916600160301b179055610ef5565b6001808201805460ff600160301b80830482169094011690920266ff000000000000199092169190911790555b6004546001820154600160301b900460ff16141561103f57610f15611945565b60026001840381548110610f2557fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e0830184815292939092849290918491840182828015610fc75780601f10610f9c57610100808354040283529160200191610fc7565b820191906000526020600020905b815481529060010190602001808311610faa57829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c090910152905061103360001984016114d6565b61103d8482611838565b505b505b5050600101610d7b565b50505050565b6000546001600160a01b0316331461106857600080fd5b60006110738461078c565b6001600160a01b0381166000908152600760205260409020549091501580156110b257506001600160a01b038116600090815260086020526040902054155b6110bb57600080fd5b6110c3611945565b6040518060e001604052808681526020018563ffffffff1681526020018461ffff168152602001600060ff168152602001600060ff16815260200160008152602001600081525090506111168282611838565b6001546111229061193b565b6004555050505050565b6005546006540143101590565b6001546111459061193b565b60045543600655565b600154811061115c57600080fd5b611164611945565b6001828154811061117157fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e08301848152929390928492909184918401828280156112135780601f106111e857610100808354040283529160200191611213565b820191906000526020600020905b8154815290600101906020018083116111f657829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c090910152805190915060009060079082906112869061078c565b6001600160a01b031681526020810191909152604001600020556001805460001981019081106112b257fe5b9060005260206000209060040201600183815481106112cd57fe5b9060005260206000209060040201600082018160000190805460018160011615610100020316600290046113029291906119ff565b5060018281018054838301805463ffffffff191663ffffffff90921691909117808255825461ffff600160201b91829004160265ffff000000001990911617808255825460ff600160301b91829004811690910266ff00000000000019909216919091178083559254600160381b908190049091160267ff00000000000000199092169190911790556002808401549083015560039283015492909101919091558054806113ac57fe5b600082815260208120600019909201916004830201906113cc8282611a74565b506001818101805467ffffffffffffffff1916905560006002830181905560039092019190915591555482146114d25760606001838154811061140b57fe5b6000918252602091829020600490910201805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561149e5780601f106114735761010080835404028352916020019161149e565b820191906000526020600020905b81548152906001019060200180831161148157829003601f168201915b5050505050905082600101600760006114b68461078c565b6001600160a01b03168152602081019190915260400160002055505b5050565b60025481106114e457600080fd5b6114ec611945565b600282815481106114f957fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e083018481529293909284929091849184018282801561159b5780601f106115705761010080835404028352916020019161159b565b820191906000526020600020905b81548152906001019060200180831161157e57829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c0909101528051909150600090600890829061160e9061078c565b6001600160a01b0316815260208101919091526040016000205560028054600019810190811061163a57fe5b90600052602060002090600402016002838154811061165557fe5b90600052602060002090600402016000820181600001908054600181600116156101000203166002900461168a9291906119ff565b5060018281018054918301805463ffffffff191663ffffffff90931692909217808355815461ffff600160201b91829004160265ffff000000001990911617808355815460ff600160301b91829004811690910266ff00000000000019909216919091178084559154600160381b908190049091160267ff000000000000001990911617905560028083015481830155600392830154929091019190915580548061173157fe5b600082815260208120600019909201916004830201906117518282611a74565b5060018101805467ffffffffffffffff191690556000600280830182905560039092015591555482146114d25760606002838154811061178d57fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156118205780601f106117f557610100808354040283529160200191611820565b820191906000526020600020905b81548152906001019060200180831161180357829003601f168201915b5050505050905082600101600860006114b68461078c565b600180548082018083556000929092528251805184926004027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019161188391839160200190611981565b5060208281015160018381018054604080880151606089015160808a015163ffffffff1990941663ffffffff9097169690961765ffff000000001916600160201b61ffff909216919091021766ff0000000000001916600160301b60ff968716021767ff000000000000001916600160381b959092169490940217905560a0850151600285015560c09094015160039093019290925591546001600160a01b0395909516600090815260079092529020929092555050565b6002900460010190565b6040805160e081018252606080825260006020830181905292820183905281018290526080810182905260a0810182905260c081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119c257805160ff19168380011785556119ef565b828001600101855582156119ef579182015b828111156119ef5782518255916020019190600101906119d4565b506119fb929150611abb565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a3857805485556119ef565b828001600101855582156119ef57600052602060002091601f016020900482015b828111156119ef578254825591600101919060010190611a59565b50805460018160011615610100020316600290046000825580601f10611a9a5750611ab8565b601f016020900490600052602060002090810190611ab89190611abb565b50565b610ad791905b808211156119fb5760008155600101611ac156fea165627a7a72305820266eed7f72c93bfb5f9e8d186ce1ca6b27bd0b82ebf40fc47e62a2e66df82f680029`

// DeployNodesV2 deploys a new Ethereum contract, binding an instance of NodesV2 to it.
func DeployNodesV2(auth *bind.TransactOpts, backend bind.ContractBackend, _blockPeriod *big.Int) (common.Address, *types.Transaction, *NodesV2, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodesV2Bin), backend, _blockPeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodesV2{NodesV2Caller: NodesV2Caller{contract: contract}, NodesV2Transactor: NodesV2Transactor{contract: contract}, NodesV2Filterer: NodesV2Filterer{contract: contract}}, nil
}

// NodesV2 is an auto generated Go binding around an Ethereum contract.
type NodesV2 struct {
	NodesV2Caller     // Read-only binding to the contract
	NodesV2Transactor // Write-only binding to the contract
	NodesV2Filterer   // Log filterer for contract events
}

// NodesV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type NodesV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesV2Session struct {
	Contract     *NodesV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesV2CallerSession struct {
	Contract *NodesV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NodesV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesV2TransactorSession struct {
	Contract     *NodesV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NodesV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type NodesV2Raw struct {
	Contract *NodesV2 // Generic contract binding to access the raw methods on
}

// NodesV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesV2CallerRaw struct {
	Contract *NodesV2Caller // Generic read-only contract binding to access the raw methods on
}

// NodesV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesV2TransactorRaw struct {
	Contract *NodesV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNodesV2 creates a new instance of NodesV2, bound to a specific deployed contract.
func NewNodesV2(address common.Address, backend bind.ContractBackend) (*NodesV2, error) {
	contract, err := bindNodesV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodesV2{NodesV2Caller: NodesV2Caller{contract: contract}, NodesV2Transactor: NodesV2Transactor{contract: contract}, NodesV2Filterer: NodesV2Filterer{contract: contract}}, nil
}

// NewNodesV2Caller creates a new read-only instance of NodesV2, bound to a specific deployed contract.
func NewNodesV2Caller(address common.Address, caller bind.ContractCaller) (*NodesV2Caller, error) {
	contract, err := bindNodesV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesV2Caller{contract: contract}, nil
}

// NewNodesV2Transactor creates a new write-only instance of NodesV2, bound to a specific deployed contract.
func NewNodesV2Transactor(address common.Address, transactor bind.ContractTransactor) (*NodesV2Transactor, error) {
	contract, err := bindNodesV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesV2Transactor{contract: contract}, nil
}

// NewNodesV2Filterer creates a new log filterer instance of NodesV2, bound to a specific deployed contract.
func NewNodesV2Filterer(address common.Address, filterer bind.ContractFilterer) (*NodesV2Filterer, error) {
	contract, err := bindNodesV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesV2Filterer{contract: contract}, nil
}

// bindNodesV2 binds a generic wrapper to an already deployed contract.
func bindNodesV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesV2 *NodesV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodesV2.Contract.NodesV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesV2 *NodesV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesV2.Contract.NodesV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesV2 *NodesV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesV2.Contract.NodesV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesV2 *NodesV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodesV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesV2 *NodesV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesV2 *NodesV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesV2.Contract.contract.Transact(opts, method, params...)
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() constant returns(uint256)
func (_NodesV2 *NodesV2Caller) ActiveNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodesV2.contract.Call(opts, out, "activeNodeCount")
	return *ret0, err
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() constant returns(uint256)
func (_NodesV2 *NodesV2Session) ActiveNodeCount() (*big.Int, error) {
	return _NodesV2.Contract.ActiveNodeCount(&_NodesV2.CallOpts)
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() constant returns(uint256)
func (_NodesV2 *NodesV2CallerSession) ActiveNodeCount() (*big.Int, error) {
	return _NodesV2.Contract.ActiveNodeCount(&_NodesV2.CallOpts)
}

// ActiveNodes is a free data retrieval call binding the contract method 0x96f9d983.
//
// Solidity: function activeNodes(uint256 ) constant returns(bytes publicKey, uint32 ip, uint16 port, uint8 joinVotes, uint8 removeVotes, uint256 lastTimeHasVoted, uint256 lastTimeHasBeenVoted)
func (_NodesV2 *NodesV2Caller) ActiveNodes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PublicKey            []byte
	Ip                   uint32
	Port                 uint16
	JoinVotes            uint8
	RemoveVotes          uint8
	LastTimeHasVoted     *big.Int
	LastTimeHasBeenVoted *big.Int
}, error) {
	ret := new(struct {
		PublicKey            []byte
		Ip                   uint32
		Port                 uint16
		JoinVotes            uint8
		RemoveVotes          uint8
		LastTimeHasVoted     *big.Int
		LastTimeHasBeenVoted *big.Int
	})
	out := ret
	err := _NodesV2.contract.Call(opts, out, "activeNodes", arg0)
	return *ret, err
}

// ActiveNodes is a free data retrieval call binding the contract method 0x96f9d983.
//
// Solidity: function activeNodes(uint256 ) constant returns(bytes publicKey, uint32 ip, uint16 port, uint8 joinVotes, uint8 removeVotes, uint256 lastTimeHasVoted, uint256 lastTimeHasBeenVoted)
func (_NodesV2 *NodesV2Session) ActiveNodes(arg0 *big.Int) (struct {
	PublicKey            []byte
	Ip                   uint32
	Port                 uint16
	JoinVotes            uint8
	RemoveVotes          uint8
	LastTimeHasVoted     *big.Int
	LastTimeHasBeenVoted *big.Int
}, error) {
	return _NodesV2.Contract.ActiveNodes(&_NodesV2.CallOpts, arg0)
}

// ActiveNodes is a free data retrieval call binding the contract method 0x96f9d983.
//
// Solidity: function activeNodes(uint256 ) constant returns(bytes publicKey, uint32 ip, uint16 port, uint8 joinVotes, uint8 removeVotes, uint256 lastTimeHasVoted, uint256 lastTimeHasBeenVoted)
func (_NodesV2 *NodesV2CallerSession) ActiveNodes(arg0 *big.Int) (struct {
	PublicKey            []byte
	Ip                   uint32
	Port                 uint16
	JoinVotes            uint8
	RemoveVotes          uint8
	LastTimeHasVoted     *big.Int
	LastTimeHasBeenVoted *big.Int
}, error) {
	return _NodesV2.Contract.ActiveNodes(&_NodesV2.CallOpts, arg0)
}

// GetInactiveNode is a free data retrieval call binding the contract method 0x93696e1a.
//
// Solidity: function getInactiveNode(uint256 index) constant returns(bytes, uint32, uint16)
func (_NodesV2 *NodesV2Caller) GetInactiveNode(opts *bind.CallOpts, index *big.Int) ([]byte, uint32, uint16, error) {
	var (
		ret0 = new([]byte)
		ret1 = new(uint32)
		ret2 = new(uint16)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _NodesV2.contract.Call(opts, out, "getInactiveNode", index)
	return *ret0, *ret1, *ret2, err
}

// GetInactiveNode is a free data retrieval call binding the contract method 0x93696e1a.
//
// Solidity: function getInactiveNode(uint256 index) constant returns(bytes, uint32, uint16)
func (_NodesV2 *NodesV2Session) GetInactiveNode(index *big.Int) ([]byte, uint32, uint16, error) {
	return _NodesV2.Contract.GetInactiveNode(&_NodesV2.CallOpts, index)
}

// GetInactiveNode is a free data retrieval call binding the contract method 0x93696e1a.
//
// Solidity: function getInactiveNode(uint256 index) constant returns(bytes, uint32, uint16)
func (_NodesV2 *NodesV2CallerSession) GetInactiveNode(index *big.Int) ([]byte, uint32, uint16, error) {
	return _NodesV2.Contract.GetInactiveNode(&_NodesV2.CallOpts, index)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 index) constant returns(bytes, uint32, uint16)
func (_NodesV2 *NodesV2Caller) GetNode(opts *bind.CallOpts, index *big.Int) ([]byte, uint32, uint16, error) {
	var (
		ret0 = new([]byte)
		ret1 = new(uint32)
		ret2 = new(uint16)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _NodesV2.contract.Call(opts, out, "getNode", index)
	return *ret0, *ret1, *ret2, err
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 index) constant returns(bytes, uint32, uint16)
func (_NodesV2 *NodesV2Session) GetNode(index *big.Int) ([]byte, uint32, uint16, error) {
	return _NodesV2.Contract.GetNode(&_NodesV2.CallOpts, index)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 index) constant returns(bytes, uint32, uint16)
func (_NodesV2 *NodesV2CallerSession) GetNode(index *big.Int) ([]byte, uint32, uint16, error) {
	return _NodesV2.Contract.GetNode(&_NodesV2.CallOpts, index)
}

// InactiveNodeCount is a free data retrieval call binding the contract method 0x6d1c76c2.
//
// Solidity: function inactiveNodeCount() constant returns(uint256)
func (_NodesV2 *NodesV2Caller) InactiveNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodesV2.contract.Call(opts, out, "inactiveNodeCount")
	return *ret0, err
}

// InactiveNodeCount is a free data retrieval call binding the contract method 0x6d1c76c2.
//
// Solidity: function inactiveNodeCount() constant returns(uint256)
func (_NodesV2 *NodesV2Session) InactiveNodeCount() (*big.Int, error) {
	return _NodesV2.Contract.InactiveNodeCount(&_NodesV2.CallOpts)
}

// InactiveNodeCount is a free data retrieval call binding the contract method 0x6d1c76c2.
//
// Solidity: function inactiveNodeCount() constant returns(uint256)
func (_NodesV2 *NodesV2CallerSession) InactiveNodeCount() (*big.Int, error) {
	return _NodesV2.Contract.InactiveNodeCount(&_NodesV2.CallOpts)
}

// InactiveNodes is a free data retrieval call binding the contract method 0x72460fa8.
//
// Solidity: function inactiveNodes(uint256 ) constant returns(bytes publicKey, uint32 ip, uint16 port, uint8 joinVotes, uint8 removeVotes, uint256 lastTimeHasVoted, uint256 lastTimeHasBeenVoted)
func (_NodesV2 *NodesV2Caller) InactiveNodes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PublicKey            []byte
	Ip                   uint32
	Port                 uint16
	JoinVotes            uint8
	RemoveVotes          uint8
	LastTimeHasVoted     *big.Int
	LastTimeHasBeenVoted *big.Int
}, error) {
	ret := new(struct {
		PublicKey            []byte
		Ip                   uint32
		Port                 uint16
		JoinVotes            uint8
		RemoveVotes          uint8
		LastTimeHasVoted     *big.Int
		LastTimeHasBeenVoted *big.Int
	})
	out := ret
	err := _NodesV2.contract.Call(opts, out, "inactiveNodes", arg0)
	return *ret, err
}

// InactiveNodes is a free data retrieval call binding the contract method 0x72460fa8.
//
// Solidity: function inactiveNodes(uint256 ) constant returns(bytes publicKey, uint32 ip, uint16 port, uint8 joinVotes, uint8 removeVotes, uint256 lastTimeHasVoted, uint256 lastTimeHasBeenVoted)
func (_NodesV2 *NodesV2Session) InactiveNodes(arg0 *big.Int) (struct {
	PublicKey            []byte
	Ip                   uint32
	Port                 uint16
	JoinVotes            uint8
	RemoveVotes          uint8
	LastTimeHasVoted     *big.Int
	LastTimeHasBeenVoted *big.Int
}, error) {
	return _NodesV2.Contract.InactiveNodes(&_NodesV2.CallOpts, arg0)
}

// InactiveNodes is a free data retrieval call binding the contract method 0x72460fa8.
//
// Solidity: function inactiveNodes(uint256 ) constant returns(bytes publicKey, uint32 ip, uint16 port, uint8 joinVotes, uint8 removeVotes, uint256 lastTimeHasVoted, uint256 lastTimeHasBeenVoted)
func (_NodesV2 *NodesV2CallerSession) InactiveNodes(arg0 *big.Int) (struct {
	PublicKey            []byte
	Ip                   uint32
	Port                 uint16
	JoinVotes            uint8
	RemoveVotes          uint8
	LastTimeHasVoted     *big.Int
	LastTimeHasBeenVoted *big.Int
}, error) {
	return _NodesV2.Contract.InactiveNodes(&_NodesV2.CallOpts, arg0)
}

// PublicKeyToAddress is a free data retrieval call binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) constant returns(address)
func (_NodesV2 *NodesV2Caller) PublicKeyToAddress(opts *bind.CallOpts, publicKey []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodesV2.contract.Call(opts, out, "publicKeyToAddress", publicKey)
	return *ret0, err
}

// PublicKeyToAddress is a free data retrieval call binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) constant returns(address)
func (_NodesV2 *NodesV2Session) PublicKeyToAddress(publicKey []byte) (common.Address, error) {
	return _NodesV2.Contract.PublicKeyToAddress(&_NodesV2.CallOpts, publicKey)
}

// PublicKeyToAddress is a free data retrieval call binding the contract method 0x43ae656c.
//
// Solidity: function publicKeyToAddress(bytes publicKey) constant returns(address)
func (_NodesV2 *NodesV2CallerSession) PublicKeyToAddress(publicKey []byte) (common.Address, error) {
	return _NodesV2.Contract.PublicKeyToAddress(&_NodesV2.CallOpts, publicKey)
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes publicKey) constant returns(bool)
func (_NodesV2 *NodesV2Caller) Registered(opts *bind.CallOpts, publicKey []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodesV2.contract.Call(opts, out, "registered", publicKey)
	return *ret0, err
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes publicKey) constant returns(bool)
func (_NodesV2 *NodesV2Session) Registered(publicKey []byte) (bool, error) {
	return _NodesV2.Contract.Registered(&_NodesV2.CallOpts, publicKey)
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes publicKey) constant returns(bool)
func (_NodesV2 *NodesV2CallerSession) Registered(publicKey []byte) (bool, error) {
	return _NodesV2.Contract.Registered(&_NodesV2.CallOpts, publicKey)
}

// AddActiveNode is a paid mutator transaction binding the contract method 0xdad7bcee.
//
// Solidity: function addActiveNode(bytes publicKey, uint32 ip, uint16 port) returns()
func (_NodesV2 *NodesV2Transactor) AddActiveNode(opts *bind.TransactOpts, publicKey []byte, ip uint32, port uint16) (*types.Transaction, error) {
	return _NodesV2.contract.Transact(opts, "addActiveNode", publicKey, ip, port)
}

// AddActiveNode is a paid mutator transaction binding the contract method 0xdad7bcee.
//
// Solidity: function addActiveNode(bytes publicKey, uint32 ip, uint16 port) returns()
func (_NodesV2 *NodesV2Session) AddActiveNode(publicKey []byte, ip uint32, port uint16) (*types.Transaction, error) {
	return _NodesV2.Contract.AddActiveNode(&_NodesV2.TransactOpts, publicKey, ip, port)
}

// AddActiveNode is a paid mutator transaction binding the contract method 0xdad7bcee.
//
// Solidity: function addActiveNode(bytes publicKey, uint32 ip, uint16 port) returns()
func (_NodesV2 *NodesV2TransactorSession) AddActiveNode(publicKey []byte, ip uint32, port uint16) (*types.Transaction, error) {
	return _NodesV2.Contract.AddActiveNode(&_NodesV2.TransactOpts, publicKey, ip, port)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x63cd6e18.
//
// Solidity: function registerNode(bytes publicKey, uint32 ip, uint16 port) returns()
func (_NodesV2 *NodesV2Transactor) RegisterNode(opts *bind.TransactOpts, publicKey []byte, ip uint32, port uint16) (*types.Transaction, error) {
	return _NodesV2.contract.Transact(opts, "registerNode", publicKey, ip, port)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x63cd6e18.
//
// Solidity: function registerNode(bytes publicKey, uint32 ip, uint16 port) returns()
func (_NodesV2 *NodesV2Session) RegisterNode(publicKey []byte, ip uint32, port uint16) (*types.Transaction, error) {
	return _NodesV2.Contract.RegisterNode(&_NodesV2.TransactOpts, publicKey, ip, port)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x63cd6e18.
//
// Solidity: function registerNode(bytes publicKey, uint32 ip, uint16 port) returns()
func (_NodesV2 *NodesV2TransactorSession) RegisterNode(publicKey []byte, ip uint32, port uint16) (*types.Transaction, error) {
	return _NodesV2.Contract.RegisterNode(&_NodesV2.TransactOpts, publicKey, ip, port)
}

// SetOverrideBlockPeriod is a paid mutator transaction binding the contract method 0x1c19b3e3.
//
// Solidity: function setOverrideBlockPeriod(bool val) returns()
func (_NodesV2 *NodesV2Transactor) SetOverrideBlockPeriod(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _NodesV2.contract.Transact(opts, "setOverrideBlockPeriod", val)
}

// SetOverrideBlockPeriod is a paid mutator transaction binding the contract method 0x1c19b3e3.
//
// Solidity: function setOverrideBlockPeriod(bool val) returns()
func (_NodesV2 *NodesV2Session) SetOverrideBlockPeriod(val bool) (*types.Transaction, error) {
	return _NodesV2.Contract.SetOverrideBlockPeriod(&_NodesV2.TransactOpts, val)
}

// SetOverrideBlockPeriod is a paid mutator transaction binding the contract method 0x1c19b3e3.
//
// Solidity: function setOverrideBlockPeriod(bool val) returns()
func (_NodesV2 *NodesV2TransactorSession) SetOverrideBlockPeriod(val bool) (*types.Transaction, error) {
	return _NodesV2.Contract.SetOverrideBlockPeriod(&_NodesV2.TransactOpts, val)
}

// Vote is a paid mutator transaction binding the contract method 0xa19e39e8.
//
// Solidity: function vote(address[] joinNodes, address[] removeNodes) returns()
func (_NodesV2 *NodesV2Transactor) Vote(opts *bind.TransactOpts, joinNodes []common.Address, removeNodes []common.Address) (*types.Transaction, error) {
	return _NodesV2.contract.Transact(opts, "vote", joinNodes, removeNodes)
}

// Vote is a paid mutator transaction binding the contract method 0xa19e39e8.
//
// Solidity: function vote(address[] joinNodes, address[] removeNodes) returns()
func (_NodesV2 *NodesV2Session) Vote(joinNodes []common.Address, removeNodes []common.Address) (*types.Transaction, error) {
	return _NodesV2.Contract.Vote(&_NodesV2.TransactOpts, joinNodes, removeNodes)
}

// Vote is a paid mutator transaction binding the contract method 0xa19e39e8.
//
// Solidity: function vote(address[] joinNodes, address[] removeNodes) returns()
func (_NodesV2 *NodesV2TransactorSession) Vote(joinNodes []common.Address, removeNodes []common.Address) (*types.Transaction, error) {
	return _NodesV2.Contract.Vote(&_NodesV2.TransactOpts, joinNodes, removeNodes)
}
