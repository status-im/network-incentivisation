pragma solidity >=0.4.21 <0.6.0;

contract NodesV2
{
  address owner;

  // Nodes that have passed the last check
  Enode[] public activeNodes;
  // Nodes that have not passed any check or failed previous check
  Enode[] public inactiveNodes;
  // Nodes that have been registered but haven't passed any checks
  Enode[] public pendingNodes;

  struct Enode {
    bytes  publicKey;
    uint32 ip;
    uint16 port;
  }

  mapping(address => bool) registered;
  mapping(address => bool) voted;
  mapping(string => uint) votes;

  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  constructor()
  public
  {
    owner = msg.sender;
  }

  // Register a node and adds it to the pending nodes
  function registerNode(bytes memory publicKey, uint32 ip, uint16 port)
  public
  {
    require(msg.sender == publicKeyToAddress(publicKey));
    //TODO: ensure that enodeID is the same as the address of the sender
    Enode memory _node = Enode(publicKey, ip, port);
    require(!registered[msg.sender]);
    registered[msg.sender] = true;
    pendingNodes.push(_node);
  }

  function pendingNodeCount() public view returns (uint) {
    return pendingNodes.length;
  }

  /*function vote(string[] memory misbehavingNodes)
  public
  {
    string memory enode = enodeIndex[msg.sender];
    // It needs to have registered
    //require(enode != "");
    // Only active nodes can vote
    require(activeNodeIndex[enode] != 0);
    // It does not have already voted
    require(!voted[msg.sender]);
    for(uint i=0; i < misbehavingNodes.length; i++) {
      votes[misbehavingNodes[i]]++;
    }
  }*/

  function publicKeyToAddress (bytes memory publicKey) public view returns (address) {
    bytes20 result =  bytes20 (uint160 (uint256 (keccak256 (publicKey))));
    //require(msg.sender == address(result));
    return address(result);
  }

  function () external payable {
    require(msg.data.length == 0);
  }
}
