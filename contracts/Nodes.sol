pragma solidity >=0.4.21 <0.6.0;

contract Nodes
{
  address owner;
  bool allowRegistration;
  string[] public nodes;
  mapping(string => uint) nodeIndex;

  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  constructor()
  public
  {
    owner = msg.sender;
    allowRegistration = false;
  }

  function nodeCount()
  view
  public
  returns
  (uint)
  {
    return nodes.length;
  }

  function addNode(string memory _node)
  public
  {
    require(allowRegistration || msg.sender == owner);
    nodeIndex[_node] = nodes.length;
    nodes.push(_node);
  }

  function deleteNode(string memory _node)
  public
  onlyOwner
  {
    uint index = nodeIndex[_node];
    _deleteNode(index);
  }

  function deleteAll()
  public
  onlyOwner
  {
    delete nodes;
  }

  function toggleRegistration(bool value)
  public
  onlyOwner
  {
    allowRegistration = value;
  }

  function _deleteNode(uint index) internal {
    require(index < nodes.length);
    nodes[index] = nodes[nodes.length-1];
    delete nodes[nodes.length-1];
    nodes.length--;
  }

  function () external payable {
    require(msg.data.length == 0);
  }
}
