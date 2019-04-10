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
const NodesV2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentSession\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"publicKeyToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inactiveNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inactiveNodes\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"joinVotes\",\"type\":\"uint8\"},{\"name\":\"removeVotes\",\"type\":\"uint8\"},{\"name\":\"lastTimeHasVoted\",\"type\":\"uint256\"},{\"name\":\"lastTimeHasBeenVoted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getInactiveNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodes\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"joinVotes\",\"type\":\"uint8\"},{\"name\":\"removeVotes\",\"type\":\"uint8\"},{\"name\":\"lastTimeHasVoted\",\"type\":\"uint256\"},{\"name\":\"lastTimeHasBeenVoted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"joinNodes\",\"type\":\"address[]\"},{\"name\":\"removeNodes\",\"type\":\"address[]\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSession\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"addActiveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_blockPerSession\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// NodesV2Bin is the compiled bytecode used for deploying new contracts.
const NodesV2Bin = `0x608060405234801561001057600080fd5b50604051602080611b7e8339810180604052602081101561003057600080fd5b5051600080546001600160a01b0319163317815560065543600555600455611b218061005d6000396000f3fe6080604052600436106100c25760003560e01c806372460fa81161007f57806396f9d9831161005957806396f9d9831461054c578063a19e39e814610576578063d4166763146106a6578063dad7bcee146106bb576100c2565b806372460fa814610421578063753408151461050d57806393696e1a14610522576100c2565b80631401795f146100cf57806343ae656c146100f65780634f0f4aa9146101c35780635aca952e1461028757806363cd6e181461034c5780636d1c76c21461040c575b36156100cd57600080fd5b005b3480156100db57600080fd5b506100e461077b565b60408051918252519081900360200190f35b34801561010257600080fd5b506101a76004803603602081101561011957600080fd5b810190602081018135600160201b81111561013357600080fd5b82018360208201111561014557600080fd5b803590602001918460018302840111600160201b8311171561016657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061079e945050505050565b604080516001600160a01b039092168252519081900360200190f35b3480156101cf57600080fd5b506101ed600480360360208110156101e657600080fd5b50356107b0565b60405180806020018463ffffffff1663ffffffff1681526020018361ffff1661ffff168152602001828103825285818151815260200191508051906020019080838360005b8381101561024a578181015183820152602001610232565b50505050905090810190601f1680156102775780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561029357600080fd5b50610338600480360360208110156102aa57600080fd5b810190602081018135600160201b8111156102c457600080fd5b8201836020820111156102d657600080fd5b803590602001918460018302840111600160201b831117156102f757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108ec945050505050565b604080519115158252519081900360200190f35b34801561035857600080fd5b506100cd6004803603606081101561036f57600080fd5b810190602081018135600160201b81111561038957600080fd5b82018360208201111561039b57600080fd5b803590602001918460018302840111600160201b831117156103bc57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813563ffffffff169250506020013561ffff1661093f565b34801561041857600080fd5b506100e4610ae5565b34801561042d57600080fd5b5061044b6004803603602081101561044457600080fd5b5035610aeb565b60405180806020018863ffffffff1663ffffffff1681526020018761ffff1661ffff1681526020018660ff1660ff1681526020018560ff1660ff168152602001848152602001838152602001828103825289818151815260200191508051906020019080838360005b838110156104cc5781810151838201526020016104b4565b50505050905090810190601f1680156104f95780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561051957600080fd5b506100e4610bd5565b34801561052e57600080fd5b506101ed6004803603602081101561054557600080fd5b5035610bdb565b34801561055857600080fd5b5061044b6004803603602081101561056f57600080fd5b5035610bf5565b34801561058257600080fd5b506100cd6004803603604081101561059957600080fd5b810190602081018135600160201b8111156105b357600080fd5b8201836020820111156105c557600080fd5b803590602001918460208302840111600160201b831117156105e657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561063557600080fd5b82018360208201111561064757600080fd5b803590602001918460208302840111600160201b8311171561066857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610c02945050505050565b3480156106b257600080fd5b506100e4611062565b3480156106c757600080fd5b506100cd600480360360608110156106de57600080fd5b810190602081018135600160201b8111156106f857600080fd5b82018360208201111561070a57600080fd5b803590602001918460018302840111600160201b8311171561072b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813563ffffffff169250506020013561ffff16611068565b6000610785611143565b15610796575060065460010161079b565b506006545b90565b8051602090910120606090811b901c90565b60606000806107bd611965565b600185815481106107ca57fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e083018481529293909284929091849184018282801561086c5780601f106108415761010080835404028352916020019161086c565b820191906000526020600020905b81548152906001019060200180831161084f57829003601f168201915b5050509183525050600182015463ffffffff811660208084019190915261ffff600160201b83041660408085019190915260ff600160301b840481166060860152600160381b9093049092166080840152600284015460a084015260039093015460c0909201919091528251918301519201519097919650945092505050565b6000806108f88361079e565b6001600160a01b03811660009081526007602052604090205490915015801561093757506001600160a01b038116600090815260086020526040902054155b159392505050565b6109488361079e565b6001600160a01b0316336001600160a01b03161461096557600080fd5b3360009081526007602052604090205415801561098f575033600090815260086020526040902054155b61099857600080fd5b6109a0611965565b506040805160e08101825284815263ffffffff841660208083019190915261ffff8416928201929092526000606082018190526080820181905260a0820181905260c08201819052600280546001810180835591909252825180519394919385936004027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0192610a359284929101906119a1565b506020828101516001830180546040808701516060880151608089015163ffffffff1990941663ffffffff9096169590951765ffff000000001916600160201b61ffff909216919091021766ff0000000000001916600160301b60ff958616021767ff000000000000001916600160381b949092169390930217905560a084015160028085019190915560c090940151600390930192909255915433600090815260089093529120555050505050565b60025490565b60028181548110610af857fe5b60009182526020918290206004919091020180546040805160026001841615610100026000190190931692909204601f810185900485028301850190915280825291935091839190830182828015610b915780601f10610b6657610100808354040283529160200191610b91565b820191906000526020600020905b815481529060010190602001808311610b7457829003601f168201915b5050505060018301546002840154600390940154929363ffffffff82169361ffff600160201b840416935060ff600160301b8404811693600160381b900416919087565b60015490565b6060600080610be8611965565b600285815481106107ca57fe5b60018181548110610af857fe5b610c0a611143565b15610c1757610c17611150565b3360009081526007602052604090205480610c3157600080fd5b600654600180830381548110610c4357fe5b9060005260206000209060040201600201541415610c6057600080fd5b600654600180830381548110610c7257fe5b600091825260208220600260049092020101919091555b8251811015610d8857600060076000858481518110610ca457fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600014610d7f576000600180830381548110610ce857fe5b90600052602060002090600402019050600654816003015414610d2b57600654600382015560018101805467ff000000000000001916600160381b179055610d59565b6001808201805460ff600160381b80830482169094011690920267ff00000000000000199092169190911790555b6003546001820154600160381b900460ff161415610d7d57610d7d6001830361116e565b505b50600101610c89565b5060005b835181101561105c576000848281518110610da357fe5b6020908102919091018101516001600160a01b03811660009081526008909252604090912054909150801561105257600060026001830381548110610de457fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152919350610e8a92849190830182828015610e805780601f10610e5557610100808354040283529160200191610e80565b820191906000526020600020905b815481529060010190602001808311610e6357829003601f168201915b505050505061079e565b6001600160a01b0316836001600160a01b031614610ea757600080fd5b600654816003015414610ed957600654600382015560018101805466ff0000000000001916600160301b179055610f06565b6001808201805460ff600160301b80830482169094011690920266ff000000000000199092169190911790555b6003546001820154600160301b900460ff16141561105057610f26611965565b60026001840381548110610f3657fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e0830184815292939092849290918491840182828015610fd85780601f10610fad57610100808354040283529160200191610fd8565b820191906000526020600020905b815481529060010190602001808311610fbb57829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c090910152905061104460001984016114f6565b61104e8482611858565b505b505b5050600101610d8c565b50505050565b60065481565b6000546001600160a01b0316331461107f57600080fd5b600061108a8461079e565b6001600160a01b0381166000908152600760205260409020549091501580156110c957506001600160a01b038116600090815260086020526040902054155b6110d257600080fd5b6110da611965565b6040518060e001604052808681526020018563ffffffff1681526020018461ffff168152602001600060ff168152602001600060ff168152602001600081526020016000815250905061112d8282611858565b6001546111399061195b565b6003555050505050565b6004546005540143101590565b60015461115c9061195b565b60035560068054600101905543600555565b600154811061117c57600080fd5b611184611965565b6001828154811061119157fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e08301848152929390928492909184918401828280156112335780601f1061120857610100808354040283529160200191611233565b820191906000526020600020905b81548152906001019060200180831161121657829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c090910152805190915060009060079082906112a69061079e565b6001600160a01b031681526020810191909152604001600020556001805460001981019081106112d257fe5b9060005260206000209060040201600183815481106112ed57fe5b906000526020600020906004020160008201816000019080546001816001161561010002031660029004611322929190611a1f565b5060018281018054838301805463ffffffff191663ffffffff90921691909117808255825461ffff600160201b91829004160265ffff000000001990911617808255825460ff600160301b91829004811690910266ff00000000000019909216919091178083559254600160381b908190049091160267ff00000000000000199092169190911790556002808401549083015560039283015492909101919091558054806113cc57fe5b600082815260208120600019909201916004830201906113ec8282611a94565b506001818101805467ffffffffffffffff1916905560006002830181905560039092019190915591555482146114f25760606001838154811061142b57fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156114be5780601f10611493576101008083540402835291602001916114be565b820191906000526020600020905b8154815290600101906020018083116114a157829003601f168201915b5050505050905082600101600760006114d68461079e565b6001600160a01b03168152602081019190915260400160002055505b5050565b600254811061150457600080fd5b61150c611965565b6002828154811061151957fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e08301848152929390928492909184918401828280156115bb5780601f10611590576101008083540402835291602001916115bb565b820191906000526020600020905b81548152906001019060200180831161159e57829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c0909101528051909150600090600890829061162e9061079e565b6001600160a01b0316815260208101919091526040016000205560028054600019810190811061165a57fe5b90600052602060002090600402016002838154811061167557fe5b9060005260206000209060040201600082018160000190805460018160011615610100020316600290046116aa929190611a1f565b5060018281018054918301805463ffffffff191663ffffffff90931692909217808355815461ffff600160201b91829004160265ffff000000001990911617808355815460ff600160301b91829004811690910266ff00000000000019909216919091178084559154600160381b908190049091160267ff000000000000001990911617905560028083015481830155600392830154929091019190915580548061175157fe5b600082815260208120600019909201916004830201906117718282611a94565b5060018101805467ffffffffffffffff191690556000600280830182905560039092015591555482146114f2576060600283815481106117ad57fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156118405780601f1061181557610100808354040283529160200191611840565b820191906000526020600020905b81548152906001019060200180831161182357829003601f168201915b5050505050905082600101600860006114d68461079e565b600180548082018083556000929092528251805184926004027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601916118a3918391602001906119a1565b5060208281015160018381018054604080880151606089015160808a015163ffffffff1990941663ffffffff9097169690961765ffff000000001916600160201b61ffff909216919091021766ff0000000000001916600160301b60ff968716021767ff000000000000001916600160381b959092169490940217905560a0850151600285015560c09094015160039093019290925591546001600160a01b0395909516600090815260079092529020929092555050565b6002900460010190565b6040805160e081018252606080825260006020830181905292820183905281018290526080810182905260a0810182905260c081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119e257805160ff1916838001178555611a0f565b82800160010185558215611a0f579182015b82811115611a0f5782518255916020019190600101906119f4565b50611a1b929150611adb565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a585780548555611a0f565b82800160010185558215611a0f57600052602060002091601f016020900482015b82811115611a0f578254825591600101919060010190611a79565b50805460018160011615610100020316600290046000825580601f10611aba5750611ad8565b601f016020900490600052602060002090810190611ad89190611adb565b50565b61079b91905b80821115611a1b5760008155600101611ae156fea165627a7a72305820c401d15488e9031dc72063821358c73d9ea26eb79694593a5364c6d0579696b10029`

// DeployNodesV2 deploys a new Ethereum contract, binding an instance of NodesV2 to it.
func DeployNodesV2(auth *bind.TransactOpts, backend bind.ContractBackend, _blockPerSession *big.Int) (common.Address, *types.Transaction, *NodesV2, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodesV2Bin), backend, _blockPerSession)
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

// CurrentSession is a free data retrieval call binding the contract method 0xd4166763.
//
// Solidity: function currentSession() constant returns(uint256)
func (_NodesV2 *NodesV2Caller) CurrentSession(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodesV2.contract.Call(opts, out, "currentSession")
	return *ret0, err
}

// CurrentSession is a free data retrieval call binding the contract method 0xd4166763.
//
// Solidity: function currentSession() constant returns(uint256)
func (_NodesV2 *NodesV2Session) CurrentSession() (*big.Int, error) {
	return _NodesV2.Contract.CurrentSession(&_NodesV2.CallOpts)
}

// CurrentSession is a free data retrieval call binding the contract method 0xd4166763.
//
// Solidity: function currentSession() constant returns(uint256)
func (_NodesV2 *NodesV2CallerSession) CurrentSession() (*big.Int, error) {
	return _NodesV2.Contract.CurrentSession(&_NodesV2.CallOpts)
}

// GetCurrentSession is a free data retrieval call binding the contract method 0x1401795f.
//
// Solidity: function getCurrentSession() constant returns(uint256)
func (_NodesV2 *NodesV2Caller) GetCurrentSession(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodesV2.contract.Call(opts, out, "getCurrentSession")
	return *ret0, err
}

// GetCurrentSession is a free data retrieval call binding the contract method 0x1401795f.
//
// Solidity: function getCurrentSession() constant returns(uint256)
func (_NodesV2 *NodesV2Session) GetCurrentSession() (*big.Int, error) {
	return _NodesV2.Contract.GetCurrentSession(&_NodesV2.CallOpts)
}

// GetCurrentSession is a free data retrieval call binding the contract method 0x1401795f.
//
// Solidity: function getCurrentSession() constant returns(uint256)
func (_NodesV2 *NodesV2CallerSession) GetCurrentSession() (*big.Int, error) {
	return _NodesV2.Contract.GetCurrentSession(&_NodesV2.CallOpts)
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
