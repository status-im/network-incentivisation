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
const NodesV2ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setOverrideBlockPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"publicKeyToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inactiveNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inactiveNodes\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"joinVotes\",\"type\":\"uint8\"},{\"name\":\"removeVotes\",\"type\":\"uint8\"},{\"name\":\"lastTimeHasVoted\",\"type\":\"uint256\"},{\"name\":\"lastTimeHasBeenVoted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getInactiveNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodes\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"},{\"name\":\"joinVotes\",\"type\":\"uint8\"},{\"name\":\"removeVotes\",\"type\":\"uint8\"},{\"name\":\"lastTimeHasVoted\",\"type\":\"uint256\"},{\"name\":\"lastTimeHasBeenVoted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"joinNodes\",\"type\":\"address[]\"},{\"name\":\"removeNodes\",\"type\":\"address[]\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"ip\",\"type\":\"uint32\"},{\"name\":\"port\",\"type\":\"uint16\"}],\"name\":\"addActiveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_blockPeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// NodesV2Bin is the compiled bytecode used for deploying new contracts.
const NodesV2Bin = `0x608060405234801561001057600080fd5b50604051602080611a388339810180604052602081101561003057600080fd5b5051600080546001600160a01b03191633179055436006556005556119de8061005a6000396000f3fe60806040526004361061009c5760003560e01c806372460fa81161006457806372460fa81461034d578063753408151461043957806393696e1a1461044e57806396f9d98314610478578063a19e39e8146104a2578063dad7bcee146105d25761009c565b80631c19b3e3146100a957806343ae656c146100d55780634f0f4aa9146101a257806363cd6e18146102665780636d1c76c214610326575b36156100a757600080fd5b005b3480156100b557600080fd5b506100a7600480360360208110156100cc57600080fd5b50351515610692565b3480156100e157600080fd5b50610186600480360360208110156100f857600080fd5b810190602081018135600160201b81111561011257600080fd5b82018360208201111561012457600080fd5b803590602001918460018302840111600160201b8311171561014557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106bc945050505050565b604080516001600160a01b039092168252519081900360200190f35b3480156101ae57600080fd5b506101cc600480360360208110156101c557600080fd5b50356106ce565b60405180806020018463ffffffff1663ffffffff1681526020018361ffff1661ffff168152602001828103825285818151815260200191508051906020019080838360005b83811015610229578181015183820152602001610211565b50505050905090810190601f1680156102565780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561027257600080fd5b506100a76004803603606081101561028957600080fd5b810190602081018135600160201b8111156102a357600080fd5b8201836020820111156102b557600080fd5b803590602001918460018302840111600160201b831117156102d657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813563ffffffff169250506020013561ffff1661080a565b34801561033257600080fd5b5061033b6109b0565b60408051918252519081900360200190f35b34801561035957600080fd5b506103776004803603602081101561037057600080fd5b50356109b7565b60405180806020018863ffffffff1663ffffffff1681526020018761ffff1661ffff1681526020018660ff1660ff1681526020018560ff1660ff168152602001848152602001838152602001828103825289818151815260200191508051906020019080838360005b838110156103f85781810151838201526020016103e0565b50505050905090810190601f1680156104255780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561044557600080fd5b5061033b610aa1565b34801561045a57600080fd5b506101cc6004803603602081101561047157600080fd5b5035610aa7565b34801561048457600080fd5b506103776004803603602081101561049b57600080fd5b5035610ac1565b3480156104ae57600080fd5b506100a7600480360360408110156104c557600080fd5b810190602081018135600160201b8111156104df57600080fd5b8201836020820111156104f157600080fd5b803590602001918460208302840111600160201b8311171561051257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561056157600080fd5b82018360208201111561057357600080fd5b803590602001918460208302840111600160201b8311171561059457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610ace945050505050565b3480156105de57600080fd5b506100a7600480360360608110156105f557600080fd5b810190602081018135600160201b81111561060f57600080fd5b82018360208201111561062157600080fd5b803590602001918460018302840111600160201b8311171561064257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813563ffffffff169250506020013561ffff16610f2e565b6000546001600160a01b031633146106a957600080fd5b6003805460ff1916911515919091179055565b8051602090910120606090811b901c90565b60606000806106db611822565b600185815481106106e857fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e083018481529293909284929091849184018282801561078a5780601f1061075f5761010080835404028352916020019161078a565b820191906000526020600020905b81548152906001019060200180831161076d57829003601f168201915b5050509183525050600182015463ffffffff811660208084019190915261ffff600160201b83041660408085019190915260ff600160301b840481166060860152600160381b9093049092166080840152600284015460a084015260039093015460c0909201919091528251918301519201519097919650945092505050565b610813836106bc565b6001600160a01b0316336001600160a01b03161461083057600080fd5b3360009081526007602052604090205415801561085a575033600090815260086020526040902054155b61086357600080fd5b61086b611822565b506040805160e08101825284815263ffffffff841660208083019190915261ffff8416928201929092526000606082018190526080820181905260a0820181905260c08201819052600280546001810180835591909252825180519394919385936004027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace019261090092849291019061185e565b506020828101516001830180546040808701516060880151608089015163ffffffff1990941663ffffffff9096169590951765ffff000000001916600160201b61ffff909216919091021766ff0000000000001916600160301b60ff958616021767ff000000000000001916600160381b949092169390930217905560a084015160028085019190915560c090940151600390930192909255915433600090815260089093529120555050505050565b6002545b90565b600281815481106109c457fe5b60009182526020918290206004919091020180546040805160026001841615610100026000190190931692909204601f810185900485028301850190915280825291935091839190830182828015610a5d5780601f10610a3257610100808354040283529160200191610a5d565b820191906000526020600020905b815481529060010190602001808311610a4057829003601f168201915b5050505060018301546002840154600390940154929363ffffffff82169361ffff600160201b840416935060ff600160301b8404811693600160381b900416919087565b60015490565b6060600080610ab4611822565b600285815481106106e857fe5b600181815481106109c457fe5b610ad6611009565b15610ae357610ae3611016565b3360009081526007602052604090205480610afd57600080fd5b600654600180830381548110610b0f57fe5b9060005260206000209060040201600201541415610b2c57600080fd5b600654600180830381548110610b3e57fe5b600091825260208220600260049092020101919091555b8251811015610c5457600060076000858481518110610b7057fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600014610c4b576000600180830381548110610bb457fe5b90600052602060002090600402019050600654816003015414610bf757600654600382015560018101805467ff000000000000001916600160381b179055610c25565b6001808201805460ff600160381b80830482169094011690920267ff00000000000000199092169190911790555b6004546001820154600160381b900460ff161415610c4957610c496001830361102b565b505b50600101610b55565b5060005b8351811015610f28576000848281518110610c6f57fe5b6020908102919091018101516001600160a01b038116600090815260089092526040909120549091508015610f1e57600060026001830381548110610cb057fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152919350610d5692849190830182828015610d4c5780601f10610d2157610100808354040283529160200191610d4c565b820191906000526020600020905b815481529060010190602001808311610d2f57829003601f168201915b50505050506106bc565b6001600160a01b0316836001600160a01b031614610d7357600080fd5b600654816003015414610da557600654600382015560018101805466ff0000000000001916600160301b179055610dd2565b6001808201805460ff600160301b80830482169094011690920266ff000000000000199092169190911790555b6004546001820154600160301b900460ff161415610f1c57610df2611822565b60026001840381548110610e0257fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e0830184815292939092849290918491840182828015610ea45780601f10610e7957610100808354040283529160200191610ea4565b820191906000526020600020905b815481529060010190602001808311610e8757829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c0909101529050610f1060001984016113b3565b610f1a8482611715565b505b505b5050600101610c58565b50505050565b6000546001600160a01b03163314610f4557600080fd5b6000610f50846106bc565b6001600160a01b038116600090815260076020526040902054909150158015610f8f57506001600160a01b038116600090815260086020526040902054155b610f9857600080fd5b610fa0611822565b6040518060e001604052808681526020018563ffffffff1681526020018461ffff168152602001600060ff168152602001600060ff1681526020016000815260200160008152509050610ff38282611715565b600154610fff90611818565b6004555050505050565b6005546006540143101590565b60015461102290611818565b60045543600655565b600154811061103957600080fd5b611041611822565b6001828154811061104e57fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e08301848152929390928492909184918401828280156110f05780601f106110c5576101008083540402835291602001916110f0565b820191906000526020600020905b8154815290600101906020018083116110d357829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c09091015280519091506000906007908290611163906106bc565b6001600160a01b0316815260208101919091526040016000205560018054600019810190811061118f57fe5b9060005260206000209060040201600183815481106111aa57fe5b9060005260206000209060040201600082018160000190805460018160011615610100020316600290046111df9291906118dc565b5060018281018054838301805463ffffffff191663ffffffff90921691909117808255825461ffff600160201b91829004160265ffff000000001990911617808255825460ff600160301b91829004811690910266ff00000000000019909216919091178083559254600160381b908190049091160267ff000000000000001990921691909117905560028084015490830155600392830154929091019190915580548061128957fe5b600082815260208120600019909201916004830201906112a98282611951565b506001818101805467ffffffffffffffff1916905560006002830181905560039092019190915591555482146113af576060600183815481106112e857fe5b6000918252602091829020600490910201805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561137b5780601f106113505761010080835404028352916020019161137b565b820191906000526020600020905b81548152906001019060200180831161135e57829003601f168201915b505050505090508260010160076000611393846106bc565b6001600160a01b03168152602081019190915260400160002055505b5050565b60025481106113c157600080fd5b6113c9611822565b600282815481106113d657fe5b6000918252602091829020604080516004939093029091018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e08301848152929390928492909184918401828280156114785780601f1061144d57610100808354040283529160200191611478565b820191906000526020600020905b81548152906001019060200180831161145b57829003601f168201915b5050509183525050600182015463ffffffff8116602083015261ffff600160201b820416604083015260ff600160301b820481166060840152600160381b909104166080820152600282015460a082015260039091015460c090910152805190915060009060089082906114eb906106bc565b6001600160a01b0316815260208101919091526040016000205560028054600019810190811061151757fe5b90600052602060002090600402016002838154811061153257fe5b9060005260206000209060040201600082018160000190805460018160011615610100020316600290046115679291906118dc565b5060018281018054918301805463ffffffff191663ffffffff90931692909217808355815461ffff600160201b91829004160265ffff000000001990911617808355815460ff600160301b91829004811690910266ff00000000000019909216919091178084559154600160381b908190049091160267ff000000000000001990911617905560028083015481830155600392830154929091019190915580548061160e57fe5b6000828152602081206000199092019160048302019061162e8282611951565b5060018101805467ffffffffffffffff191690556000600280830182905560039092015591555482146113af5760606002838154811061166a57fe5b6000918252602091829020600490910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156116fd5780601f106116d2576101008083540402835291602001916116fd565b820191906000526020600020905b8154815290600101906020018083116116e057829003601f168201915b505050505090508260010160086000611393846106bc565b600180548082018083556000929092528251805184926004027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601916117609183916020019061185e565b5060208281015160018381018054604080880151606089015160808a015163ffffffff1990941663ffffffff9097169690961765ffff000000001916600160201b61ffff909216919091021766ff0000000000001916600160301b60ff968716021767ff000000000000001916600160381b959092169490940217905560a0850151600285015560c09094015160039093019290925591546001600160a01b0395909516600090815260079092529020929092555050565b6002900460010190565b6040805160e081018252606080825260006020830181905292820183905281018290526080810182905260a0810182905260c081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061189f57805160ff19168380011785556118cc565b828001600101855582156118cc579182015b828111156118cc5782518255916020019190600101906118b1565b506118d8929150611998565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061191557805485556118cc565b828001600101855582156118cc57600052602060002091601f016020900482015b828111156118cc578254825591600101919060010190611936565b50805460018160011615610100020316600290046000825580601f106119775750611995565b601f0160209004906000526020600020908101906119959190611998565b50565b6109b491905b808211156118d8576000815560010161199e56fea165627a7a72305820269c443b88a1ed34af058b67bb94d97e38dd2e63df42e221799e92b6e9718e560029`

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
