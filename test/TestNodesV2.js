const Nodes = artifacts.require('NodesV2');
const helper = require('./helpers');

const blockPeriod = 10;

function ipToDecimal(dot) {
  const d = dot.split('.');
  return ((((((+d[0]) * 256) + (+d[1])) * 256) + (+d[2])) * 256) + (+d[3]);
}

function compareNodes(a, b) {
  if (a.pk < b.pk) {
    return -1;
  }
  if (a.pk > b.pk) {
    return 1;
  }
  return 0;
}

const node1 = {
  pk: '0xaf80b90d25145da28c583359beb47b21796b2fe1a23c1511e443e7a64dfdb27d7434c380f0aa4c500e220aa1a9d068514b1ff4d5019e624e7ba1efe82b340a59',
  ip: ipToDecimal('1.1.1.1'),
  port: 1,
  address: '0x627306090abab3a6e1400e9345bc60c78a8bef57',
};

const node2 = {
  pk: '0xce7edc292d7b747fab2f23584bbafaffde5c8ff17cf689969614441e0527b90015ea9fee96aed6d9c0fc2fbe0bd1883dee223b3200246ff1e21976bdbc9a0fc8',
  ip: ipToDecimal('1.1.1.2'),
  port: 2,
  address: '0xf17f52151ebef6c7334fad080c5704d77216b732',
};

const node3 = {
  pk: '0x785a891f323acd6cef0fc509bb14304410595914267c50467e51c87142acbb5e12ab6cc10390ee6e31985e52e7d5701bed4c265dfc899cac07bc9c608ab02a74',
  ip: ipToDecimal('1.1.1.3'),
  port: 3,
  address: '0xc5fdf4076b8f3a5357c5e395ab970b5b54098fef',
};

const node4 = {
  pk: '0x396c2c8a22ec28dbe02613027edea9a3b0c314294985e09c2f389818b29fee066a8384bb217f7fc35134d0e778ad8681b2d922ed630f86bea161d27f97536a3e',
  ip: ipToDecimal('1.1.1.4'),
  port: 4,
  address: '0x821aea9a577a9b44299b9c15c88cf3087f3b5544',
};

const node5 = {
  pk: '0xe67ceb1f0af0ab4668227984782b48d286b88e54dc91487143199728d4597c02e080ed81eda0d36242b0dadbb56f9cad0bdcf42b6c51d6e839d47681e787d9bd',
  ip: ipToDecimal('1.1.1.5'),
  port: 5,
  address: '0x0d1d4e623d10f9fba5db95830f7d3839406c6af2',
};

contract('Nodes', async (accounts) => {
  let instance;
  beforeEach(async () => {
    instance = await Nodes.new(blockPeriod);
  });

  describe('registerNode', async () => {
    beforeEach(async () => {
      await instance.registerNode(node1.pk, node1.ip, node1.port);
      await instance.registerNode(node2.pk, node2.ip, node2.port, { from: accounts[1] });
    });

    it('adds the nodes', async () => {
      const actualNodeCount = await instance.inactiveNodeCount();
      assert.equal(2, actualNodeCount);
    });

    describe('node is already in inactive', async () => {
      it('throws an exception', async () => {
        try {
          await instance.registerNode(node1.pk, node1.ip, node1.port);
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
    describe('node public key and sender do not match', async () => {
      it('throws an exception', async () => {
        try {
          await instance.registerNode(node1.pk, node1.ip, node1.port, { from: accounts[1] });
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });

  describe('addActiveNode', async () => {
    beforeEach(async () => {
      await instance.addActiveNode(node1.pk, node1.ip, node1.port);
    });

    it('adds the nodes', async () => {
      const actualNodeCount = await instance.activeNodeCount();
      assert.equal(1, actualNodeCount);
    });

    describe('node is already in active', async () => {
      it('throws an exception', async () => {
        try {
          await instance.addActiveNode(node1.pk, node1.ip, node1.port);
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
    describe('not the owner', async () => {
      it('throws an exception', async () => {
        try {
          await instance.addActiveNode(node1.pk, node1.ip, node1.port, { from: accounts[1] });
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });

  describe('registered', async () => {
    beforeEach(async () => {
      await instance.registerNode(node1.pk, node1.ip, node1.port);
    });

    it('it returns true when node is registered', async () => {
      const response = await instance.registered(node1.pk);
      assert.equal(true, response);
    });

    it('it returns true when node is registered', async () => {
      const response = await instance.registered(node2.pk);
      assert.equal(false, response);
    });
  });

  describe('vote', async () => {
    describe('first round of voting', async () => {
      // Node1, Node2 and Node3 are active, Node4 is inactive
      beforeEach(async () => {
        await instance.addActiveNode(node1.pk, node1.ip, node1.port);
        await instance.addActiveNode(node2.pk, node2.ip, node2.port);
        await instance.addActiveNode(node3.pk, node3.ip, node3.port);
        await instance.registerNode(node4.pk, node4.ip, node4.port, { from: accounts[3] });
        await instance.registerNode(node5.pk, node5.ip, node5.port, { from: accounts[4] });
        // Last block period finished
        await helper.advanceBlocks(blockPeriod);
        await instance.vote([node5.address], [node2.address]);
        await instance.vote([], [node1.address, node3.address], { from: accounts[1] });
        await instance.vote([node4.address, node5.address], [node2.address], { from: accounts[2] });
      });

      describe('currentBlock', async () => {
        it('returns 1', async () => {
          const actualCurrentBlock = await instance.getCurrentBlock();
          assert.equal(1, actualCurrentBlock);
        });
      });

      describe('double voting', async () => {
        it('throws an exception', async () => {
          try {
            await instance.vote([node4.address], [node2.address]);
          } catch (error) {
            return;
          }
          assert.fail('it should throw an exception');
        });
      });

      describe('voter is not active', async () => {
        it('throws an exception', async () => {
          try {
            await instance.vote([node5.address], [node2.address], { from: accounts[3] });
          } catch (error) {
            return;
          }
          assert.fail('it should throw an exception');
        });
      });

      xdescribe('voter is newly active', async () => {
        it('throws an exception', async () => {
          try {
            await instance.vote([node4.address], [node2.address], { from: accounts[3] });
          } catch (error) {
            return;
          }
          assert.fail('it should throw an exception');
        });
      });

      describe('second round of voting', async () => {
        beforeEach(async () => {
          await helper.advanceBlocks(blockPeriod);
          // We vote to trigger changes
          await instance.vote([], [node2.address]);
        });

        describe('currentBlock', async () => {
          it('returns 2', async () => {
            const actualCurrentBlock = await instance.getCurrentBlock();
            assert.equal(2, actualCurrentBlock);
          });
        });

        it('promotes the inactive nodes that have passed the checks', async () => {
          const actualNodeCount = await instance.inactiveNodeCount();
          assert.equal(1, actualNodeCount);
          const inactiveNode1 = await instance.getInactiveNode(0);

          assert.equal(inactiveNode1[0], node4.pk);
          assert.equal(inactiveNode1[1], node4.ip);
          assert.equal(inactiveNode1[2], node4.port);
        });

        it('removes nodes that have failed the checks and promotes those that have passed them', async () => {
          const actualNodeCount = await instance.activeNodeCount();
          assert.equal(3, actualNodeCount);

          const actualNode1 = await instance.getNode(0);
          const actualNode2 = await instance.getNode(1);
          const actualNode3 = await instance.getNode(2);

          const actualNodes = [
            {
              pk: actualNode1[0],
              ip: actualNode1[1],
              port: actualNode1[2],
            },
            {
              pk: actualNode2[0],
              ip: actualNode2[1],
              port: actualNode2[2],
            },
            {
              pk: actualNode3[0],
              ip: actualNode3[1],
              port: actualNode3[2],
            },
          ];

          actualNodes.sort(compareNodes);

          assert.equal(actualNodes[0].pk, node3.pk);
          assert.equal(actualNodes[0].port, node3.port);
          assert.equal(actualNodes[0].ip, node3.ip);

          assert.equal(actualNodes[1].pk, node1.pk);
          assert.equal(actualNodes[1].port, node1.port);
          assert.equal(actualNodes[1].ip, node1.ip);


          assert.equal(actualNodes[2].pk, node5.pk);
          assert.equal(actualNodes[2].port, node5.port);
          assert.equal(actualNodes[2].ip, node5.ip);
        });
      });
    });
  });
});
