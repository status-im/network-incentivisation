pragma solidity >=0.4.21 <0.6.0;

contract NodesV2
{
  address owner;

  // Nodes that have passed the last check
  Enode[] public activeNodes;
  // Nodes that have not passed any check or failed previous check
  Enode[] public inactiveNodes;

  uint16 quorum;
  uint16 maxNodes = 6000;
  // How many blocks is a session
  uint16 blockPerSession;
  // How many blocks per subscription period
  uint24 blockPerSubscription = 192960;
  uint currentSessionStart;
  // How much to send to user at each voting
  uint currentReward;
  uint32 public currentSession;

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
    // Session when the node joine
    uint32 joiningSession;
    // Session when became active
    uint32 activeSession;
  }

  mapping(address => uint) activeNodeIndex;
  mapping(address => uint) inactiveNodeIndex;

  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  constructor(uint16 _blockPerSession)
  public
  {
    owner = msg.sender;
    currentSession = 0;
    currentSessionStart = block.number;
    blockPerSession = _blockPerSession;
  }

  function getCurrentSession() public view returns (uint) {
    if (newSession()) {
      return currentSession + 1;
    } else {
      return currentSession;
    }
  }

  function getNode(uint index) public view returns (bytes memory, uint32, uint16, uint32, uint32) {
    Enode memory enode = activeNodes[index];

    return (enode.publicKey, enode.ip, enode.port, enode.joiningSession, enode.activeSession);
  }

  function getInactiveNode(uint index) public view returns (bytes memory, uint32, uint16, uint32, uint32) {
    Enode memory enode = inactiveNodes[index];

    return (enode.publicKey, enode.ip, enode.port, enode.joiningSession, enode.activeSession);
  }


  // Register a node and adds it to the pending nodes
  function registerNode(bytes memory publicKey, uint32 ip, uint16 port)
  public
  {
    // Ensure it's registering their own node
    require(msg.sender == publicKeyToAddress(publicKey));
    // Make sure they haven't registered already
    require(activeNodeIndex[msg.sender] == 0 && inactiveNodeIndex[msg.sender] == 0);
    // Make sure we don't hit the max nodes limit
    require(activeNodes.length + inactiveNodes.length < maxNodes);

    Enode memory _node = Enode(publicKey, ip, port, 0, 0, 0, 0, currentSession, 0);
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

  function overrideSession() onlyOwner public{
    performSessionOperation();
  }

  function performSessionOperation() internal {
    // Reset quorum for next vote
    quorum = calculateQuorum(activeNodes.length);
    // Set current session
    currentSession++;
    // Set start
    currentSessionStart = block.number;
    // Set reward
    currentReward = calculateReward();
  }

  function vote(
    address[] memory joinNodes,
    address[] memory removeNodes
  )
  public
  {
    if (newSession()) {
      performSessionOperation();
    }

    // Make sure is an active node
    uint enodeIndex = activeNodeIndex[msg.sender];
    require(enodeIndex != 0);

    // Make sure it has not voted already
    require(activeNodes[enodeIndex - 1].lastTimeHasVoted != currentSession);

    // Make sure it has not just became active, we allow 0 to ease bootstrapping
    require(activeNodes[enodeIndex - 1].activeSession == 0 || activeNodes[enodeIndex - 1].activeSession != currentSession);

    activeNodes[enodeIndex - 1].lastTimeHasVoted = currentSession;

    // Remove nodes that failed the vote
    for(uint i=0; i < removeNodes.length; i++) {
      uint index = activeNodeIndex[removeNodes[i]];
      if (index != 0) {
        Enode storage enode = activeNodes[index - 1];
        if (enode.lastTimeHasBeenVoted != currentSession) {
          enode.lastTimeHasBeenVoted = currentSession;
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
        if (enode.lastTimeHasBeenVoted != currentSession) {
          enode.lastTimeHasBeenVoted = currentSession;
          enode.joinVotes = 1;
        } else {
            enode.joinVotes++;
        }

        if (enode.joinVotes == quorum) {
          Enode memory enodeCopy = inactiveNodes[index - 1];
          enodeCopy.activeSession = currentSession;
          _deleteInactiveNode(index - 1);
          _addActiveNode(enodeAddress, enodeCopy);
        }
      }
    }

    msg.sender.transfer(currentReward);

  }

  function newSession() private view returns (bool) {
    return block.number >= currentSessionStart + blockPerSession;
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
    // Make sure we don't hit the max nodes limit
    require(activeNodes.length + inactiveNodes.length < maxNodes);

    Enode memory _node = Enode(publicKey, ip, port, 0, 0, 0, 0, currentSession, currentSession);

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


  function calculateQuorum(uint a) pure internal returns (uint16) {
    return uint16(a) / 2 + 1;
  }

  function calculateReward() view internal returns (uint) {
    return address(this).balance / (blockPerSubscription / blockPerSession) / activeNodes.length;
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
