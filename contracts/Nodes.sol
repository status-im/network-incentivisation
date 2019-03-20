pragma solidity >=0.4.21 <0.6.0;

contract Nodes
{
  address owner;
  string[] public nodes;
  uint public nodeCount;

  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  constructor()
  public
  {
    owner = msg.sender;
  }

  function addNode(string memory _node)
  public
  onlyOwner
  {
    nodes.push(_node);
    nodeCount++;
  }

  function deleteAll()
  public
  onlyOwner
  {
    delete nodes;
    nodeCount = 0;
  }

  function () external payable {
    require(msg.data.length == 0);
  }
}
