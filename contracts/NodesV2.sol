pragma solidity >=0.4.21 <0.6.0;

contract NodesV2
{
  address owner;
  bool allowRegistration;

  // Nodes that have passed the last check
  string[] public activeNodes;
  // Nodes that have not passed any check or failed previous check
  string[] public inactiveNodes;
  // Nodes that have been registered but haven't passed any checks
  string[] public pendingNodes;


  mapping(string => uint) activeNodeIndex;
  mapping(string => uint) pendingNodeIndex;

  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  modifier registrationOpen() {
    require(allowRegistration);
    _;
  }

  constructor()
  public
  {
    owner = msg.sender;
    allowRegistration = false;
  }

  function activeNodeCount()
  view
  public
  returns
  (uint)
  {
    return activeNodes.length;
  }

  function pendingNodeCount()
  view
  public
  returns
  (uint)
  {
    return pendingNodes.length;
  }

  // Register a node and adds it to the pending nodes
  function registerNode(string memory _node)
  public
  registrationOpen
  {
    require(pendingNodeIndex[_node] == 0);
    require(activeNodeIndex[_node] == 0);
    pendingNodeIndex[_node] = pendingNodes.length + 1;
    pendingNodes.push(_node);
  }

  // Add an active node to bootstrap the system, only used by the owner
  function addActiveNode(string memory _node)
  public
  onlyOwner
  {
    // We increment as 0 is the default value
    activeNodeIndex[_node] = activeNodes.length + 1;
    activeNodes.push(_node);
  }

  // Remove an active node from the list of nodes, only used by the owner
  function deleteActiveNode(string memory _node)
  public
  onlyOwner
  {
    uint index = activeNodeIndex[_node];

    require(index != 0);
    _deleteNode(activeNodes, index - 1);
    activeNodeIndex[_node] = 0;
  }

  // Remove a pending node
  function deletePendingNode(string memory _node)
  public
  onlyOwner
  {
    uint index = pendingNodeIndex[_node];

    require(index != 0);
    _deleteNode(pendingNodes, index - 1);
    pendingNodeIndex[_node] = 0;
  }

  // Remove all nodes, only used by the owner
  function deleteAll()
  public
  onlyOwner
  {
    for(uint i=0; i < activeNodes.length; i++) {
      activeNodeIndex[activeNodes[i]] = 0;
    }
    delete activeNodes;

    for(uint i=0; i < pendingNodes.length; i++) {
      pendingNodeIndex[pendingNodes[i]] = 0;
    }
    delete pendingNodes;
  }

  function toggleRegistration(bool value)
  public
  onlyOwner
  {
    allowRegistration = value;
  }

  function _deleteNode(string[] storage nodes, uint index) internal {
    require(index < nodes.length);
    nodes[index] = nodes[nodes.length-1];
    delete nodes[nodes.length-1];
    nodes.length--;
  }

  function () external payable {
    require(msg.data.length == 0);
  }
}
