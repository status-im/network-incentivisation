
function advanceBlock() {
  const id = Date.now();

  return new Promise((resolve, reject) => {
    web3.currentProvider.send(
      {
        jsonrpc: '2.0',
        method: 'evm_mine',
        id,
      },
      (err, res) => {
        if (err) return reject(err);
        return resolve(res);
      },
    );
  });
}

const advanceBlocks = async (n) => {
  const promises = [];
  for (let i = 0; i < n; i += 1) {
    promises.push(advanceBlock());
  }

  await Promise.all(promises);
};

module.exports = {
  advanceBlocks,
};
