pragma solidity >=0.4.21 <0.6.0;

contract NodesV2
{
  address owner;

  // Nodes that have passed the last check
  Enode[] public activeNodes;
  // Nodes that have not passed any check or failed previous check
  Enode[] public inactiveNodes;

  uint quorum;
  // How many blocks is a period
  uint blockPeriod;
  uint currentBlockStart;
  uint public currentBlock;

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
    currentBlock = 0;
    currentBlockStart = block.number;
    blockPeriod = _blockPeriod;
  }

  function getCurrentBlock() public view returns (uint) {
    if (newBlockPeriod()) {
      return currentBlock + 1;
    } else {
      return currentBlock;
    }
  }

  function getNode(uint index) public view returns (bytes memory, uint32, uint16) {
    Enode memory enode = activeNodes[index];

    return (enode.publicKey, enode.ip, enode.port);
  }

  function getInactiveNode(uint index) public view returns (bytes memory, uint32, uint16) {
    Enode memory enode = inactiveNodes[index];

    return (enode.publicKey, enode.ip, enode.port);
  }


  // Register a node and adds it to the pending nodes
  function registerNode(bytes memory publicKey, uint32 ip, uint16 port)
  public
  {
    // Ensure it's registering their own node
    require(msg.sender == publicKeyToAddress(publicKey));
    // Make sure they haven't registered already
    require(activeNodeIndex[msg.sender] == 0 && inactiveNodeIndex[msg.sender] == 0);

    Enode memory _node = Enode(publicKey, ip, port, 0, 0, 0, 0);
    inactiveNodes.push(_node);
    inactiveNodeIndex[msg.sender] = inactiveNodes.length;
  }

  function registered(bytes memory publicKey)
  public
  view
  returns (bool) {
    address node = publicKeyToAddress(publicKey);
    return !(activeNodeIndex[node] == 0 && inactiveNodeIndex[node] == 0);
  }

 function activeNodeCount() public view returns (uint) {
    return activeNodes.length;
  }

  function inactiveNodeCount() public view returns (uint) {
    return inactiveNodes.length;
  }

  function performBlockPeriodOperation() internal {
    // Reset quorum for next vote
    quorum = calculateQuorum(activeNodes.length);
    // Set current block
    currentBlock++;
    // Set start
    currentBlockStart = block.number;
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
      uint index = inactiveNodeIndex[enodeAddress];
      if (index != 0) {
        Enode storage enode = inactiveNodes[index - 1];
        require(enodeAddress == publicKeyToAddress(enode.publicKey));
        if (enode.lastTimeHasBeenVoted != currentBlock) {
          enode.lastTimeHasBeenVoted = currentBlock;
          enode.joinVotes = 1;
        } else {
            enode.joinVotes++;
        }

        if (enode.joinVotes == quorum) {
          Enode memory enodeCopy = inactiveNodes[index - 1];
          _deleteInactiveNode(index - 1);
          _addActiveNode(enodeAddress, enodeCopy);
        }
      }
    }
  }

  function newBlockPeriod() private view returns (bool) {
    return block.number >= currentBlockStart + blockPeriod;
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
    require(activeNodeIndex[enodeAddress] == 0 && inactiveNodeIndex[enodeAddress] == 0);

    Enode memory _node = Enode(publicKey, ip, port, 0, 0, 0, 0);

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

  function _deleteInactiveNode(uint index) internal {
    require(index < inactiveNodes.length);
    // Remove from index
    Enode memory oldNode = inactiveNodes[index];
    inactiveNodeIndex[publicKeyToAddress(oldNode.publicKey)] = 0;

    // Copy last element to a previous cell
    inactiveNodes[index] = inactiveNodes[inactiveNodes.length-1];

    // Shorten array
    inactiveNodes.pop();

    // Delete last element if necessary
    // Update the index
    if (index != inactiveNodes.length) {
      bytes memory publicKey = inactiveNodes[index].publicKey;
      inactiveNodeIndex[publicKeyToAddress(publicKey)] = index + 1;
    }
  }

  function _deleteActiveNode(uint index) internal {
    require(index < activeNodes.length);
    // Remove from index
    Enode memory oldNode = activeNodes[index];
    activeNodeIndex[publicKeyToAddress(oldNode.publicKey)] = 0;

    // Copy last element to a previous cell
    activeNodes[index] = activeNodes[activeNodes.length-1];

    // Shorten array
    activeNodes.pop();

    // Delete last element if necessary
    // Update the index
    if (index != activeNodes.length) {
      bytes memory publicKey = activeNodes[index].publicKey;
      activeNodeIndex[publicKeyToAddress(publicKey)] = index + 1;
    }
  }

  function () external payable {
    require(msg.data.length == 0);
  }
}
