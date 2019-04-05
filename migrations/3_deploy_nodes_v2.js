const NodesV2 = artifacts.require('./NodesV2.sol');

module.exports = function deploy(deployer) {
  deployer.deploy(NodesV2, 240);
};
