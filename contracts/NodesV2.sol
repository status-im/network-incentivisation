pragma solidity >=0.4.21 <0.6.0;

contract NodesV2
{
  address owner;

  // Nodes that have passed the last check
  Enode[] public activeNodes;
  // Nodes that have not passed any check or failed previous check
  Enode[] public inactiveNodes;

  // For testing purposes
  bool overrideBlockPeriod;
  uint quorum;
  uint blockPeriod;
  uint currentBlock;

  struct Enode {
    bytes  publicKey;
    uint32 ip;
    uint16 port;
    uint8  joinVotes;
    uint8  removeVotes;
    // When was the last vote cast by this node
    uint lastTimeHasVoted;
    // For which voting round are the votes referring to
    uint lastTimeHasBeenVoted;
  }

  mapping(address => bool) registered;
  mapping(address => uint) activeNodeIndex;
  mapping(address => uint) inactiveNodeIndex;

  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  constructor(uint _blockPeriod)
  public
  {
    owner = msg.sender;
    currentBlock = block.number;
    blockPeriod = _blockPeriod;
  }

  // Register a node and adds it to the pending nodes
  function registerNode(bytes memory publicKey, uint32 ip, uint16 port)
  public
  {
    // Ensure it's registering their own node
    require(msg.sender == publicKeyToAddress(publicKey));
    // Make sure they haven't registered already
    require(!registered[msg.sender]);

    Enode memory _node = Enode(publicKey, ip, port, 0, 0, 0, 0);
    registered[msg.sender] = true;
    inactiveNodes.push(_node);
    inactiveNodeIndex[msg.sender] = inactiveNodes.length;
  }

 function activeNodeCount() public view returns (uint) {
    return activeNodes.length;
  }

  function inactiveNodeCount() public view returns (uint) {
    return inactiveNodes.length;
  }

  // When a new block starts:
  // - All pending nodes become inactive
  function performBlockPeriodOperation() internal {
    // Reset quorum for next vote
    quorum = calculateQuorum(activeNodes.length);
    currentBlock = block.number;
  }

  function vote(
    address[] memory joinNodes,
    address[] memory removeNodes
  )
  public
  {
    if (newBlockPeriod()) {
      performBlockPeriodOperation();
    }

    // Make sure is an active node
    uint enodeIndex = activeNodeIndex[msg.sender];
    require(enodeIndex != 0);

    // Make sure it has not voted already
    require(activeNodes[enodeIndex - 1].lastTimeHasVoted != currentBlock);
    activeNodes[enodeIndex - 1].lastTimeHasVoted = currentBlock;

    // Remove nodes that failed the vote
    for(uint i=0; i < removeNodes.length; i++) {
      uint index = activeNodeIndex[removeNodes[i]];
      if (index != 0) {
        Enode storage enode = activeNodes[index - 1];
        if (enode.lastTimeHasBeenVoted != currentBlock) {
          enode.lastTimeHasBeenVoted = currentBlock;
          enode.removeVotes = 1;
        } else {
            enode.removeVotes++;
        }

        if (enode.removeVotes == quorum) {
          _deleteActiveNode(index - 1);
        }
      }
    }

    // Promote nodes that passed the vote
    for(uint i=0; i < joinNodes.length; i++) {
      address enodeAddress = joinNodes[i];
      uint index = inactiveNodeIndex[joinNodes[i]];
      if (index != 0) {
        Enode storage enode = inactiveNodes[index - 1];
        if (enode.lastTimeHasBeenVoted != currentBlock) {
          enode.lastTimeHasBeenVoted = currentBlock;
          enode.joinVotes = 1;
        } else {
            enode.joinVotes++;
        }

        if (enode.joinVotes == quorum) {
          _deleteInactiveNode(index - 1);
          _addActiveNode(enodeAddress, enode);
        }
      }
    }
  }

  function newBlockPeriod() private view returns (bool) {
    return block.number >= currentBlock + blockPeriod;
  }

  function publicKeyToAddress (bytes memory publicKey) public pure returns (address) {
    bytes20 result =  bytes20 (uint160 (uint256 (keccak256 (publicKey))));
    return address(result);
  }

  function addActiveNode(bytes memory publicKey, uint32 ip, uint16 port)
  onlyOwner
  public
  {
    address enodeAddress = publicKeyToAddress(publicKey);
    // Make sure they haven't registered already
    require(!registered[enodeAddress]);

    Enode memory _node = Enode(publicKey, ip, port, 0, 0, 0, 0);
    registered[enodeAddress] = true;

    _addActiveNode(enodeAddress, _node);

    quorum = calculateQuorum(activeNodes.length);
  }

  function _addActiveNode(address enodeAddress, Enode memory enode)
  internal
  {
    activeNodes.push(enode);

    // Index starts with 1 so we can use to check existence
    activeNodeIndex[enodeAddress] = activeNodes.length;
  }


  function calculateQuorum(uint a) pure internal returns (uint ) {
    return a / 2 + 1;
  }

  function setOverrideBlockPeriod(bool val) public onlyOwner {
    overrideBlockPeriod = val;
  }

  function _deleteInactiveNode(uint index) internal {
    require(index < inactiveNodes.length);
    inactiveNodes[index] = inactiveNodes[inactiveNodes.length-1];
    delete inactiveNodes[inactiveNodes.length-1];
    inactiveNodes.length--;
    bytes memory publicKey = inactiveNodes[index].publicKey;
    inactiveNodeIndex[publicKeyToAddress(publicKey)] = index + 1;
  }

  function _deleteActiveNode(uint index) internal {
    require(index < activeNodes.length);
    activeNodes[index] = activeNodes[activeNodes.length-1];
    delete activeNodes[activeNodes.length-1];
    activeNodes.length--;
    bytes memory publicKey = activeNodes[index].publicKey;
    activeNodeIndex[publicKeyToAddress(publicKey)] = index + 1;
  }



  function () external payable {
    require(msg.data.length == 0);
  }
}
