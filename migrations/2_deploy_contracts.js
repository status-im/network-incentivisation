const Nodes = artifacts.require('./Nodes.sol');

module.exports = function deploy(deployer) {
  deployer.deploy(Nodes);
};
