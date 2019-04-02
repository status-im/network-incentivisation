const Nodes = artifacts.require('NodesV2');

const node1 = 'enode://da61e9eff86a56633b635f887d8b91e0ff5236bbc05b8169834292e92afb92929dcf6efdbf373a37903da8fe0384d5a0a8247e83f1ce211aa429200b6d28c548@47.91.156.93:443';
const node2 = 'enode://7de99e4cb1b3523bd26ca212369540646607c721ad4f3e5c821ed9148150ce6ce2e72631723002210fac1fd52dfa8bbdf3555e05379af79515e1179da37cc3db@35.188.19.210:443';
const node3 = 'enode://ebefab39b69bbbe64d8cd86be765b3be356d8c4b24660f65d493143a0c44f38c85a257300178f7845592a1b0332811542e9a58281c835babdd7535babb64efc1@35.202.99.224:443';

contract('Nodes', async (accounts) => {
  let instance;
  beforeEach(async () => {
    instance = await Nodes.deployed();
    await instance.deleteAll();
    await instance.toggleRegistration(false);
  });

  describe('addActiveNode', async () => {
    describe('called by the owner', async () => {
      beforeEach(async () => {
        await instance.addActiveNode(node1);
        await instance.addActiveNode(node2);
      });

      it('adds the first node', async () => {
        const actualNode1 = await instance.activeNodes(0);
        assert.equal(actualNode1, node1);
      });

      it('adds the second node', async () => {
        const actualNode2 = await instance.activeNodes(1);
        assert.equal(actualNode2, node2);
      });

      it('sets the count', async () => {
        const actualNodeCount = await instance.activeNodeCount();
        assert.equal(2, actualNodeCount);
      });
    });
    describe('called by someone else', async () => {
      it('throws an exception', async () => {
        try {
          await instance.addActiveNode(node1, { from: accounts[1] });
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });

  describe('registerNode', async () => {
    describe('registration is open', async () => {
      beforeEach(async () => {
        await instance.toggleRegistration(true);
        await instance.registerNode(node1);
        await instance.registerNode(node2, { from: accounts[1] });
      });

      it('adds the first node', async () => {
        const actualNode1 = await instance.pendingNodes(0);
        assert.equal(actualNode1, node1);
      });

      it('adds the second node', async () => {
        const actualNode2 = await instance.pendingNodes(1);
        assert.equal(actualNode2, node2);
      });

      it('sets the count', async () => {
        const actualNodeCount = await instance.pendingNodeCount();
        assert.equal(2, actualNodeCount);
      });

      describe('node is already in pending', async () => {
        it('throws an exception', async () => {
          try {
            await instance.registerNode(node1);
          } catch (error) {
            return;
          }
          assert.fail('it should throw an exception');
        });
      });

      describe('node is already in active', async () => {
        it('throws an exception', async () => {
          await instance.addActiveNode(node3);
          try {
            await instance.registerNode(node3);
          } catch (error) {
            return;
          }
          assert.fail('it should throw an exception');
        });
      });
    });
    describe('registration is closed', async () => {
      it('throws an exception', async () => {
        try {
          await instance.registerNode(node1);
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });

  describe('deleteActiveNode', async () => {
    describe('called by the owner', async () => {
      beforeEach(async () => {
        await instance.addActiveNode(node1);
        await instance.addActiveNode(node2);
        await instance.deleteActiveNode(node1);
      });

      it('removes the first node', async () => {
        const actualNode2 = await instance.activeNodes(0);
        assert.equal(actualNode2, node2);
      });

      it('sets the count', async () => {
        const actualNodeCount = await instance.activeNodeCount();
        assert.equal(1, actualNodeCount);
      });
    });
    describe('called by someone else', async () => {
      it('throws an exception', async () => {
        try {
          await instance.deleteActiveNode(node1, { from: accounts[1] });
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });

  describe('deleteAll', async () => {
    beforeEach(async () => {
      await instance.addActiveNode(node1);
      await instance.addActiveNode(node2);
      await instance.deleteAll();
    });

    it('empties the array', async () => {
      try {
        await instance.activeNodes(0);
      } catch (error) {
        return;
      }
      assert.fail('it should throw an exception');
    });

    it('sets the count', async () => {
      const actualNodeCount = await instance.activeNodeCount();
      assert.equal(0, actualNodeCount);
    });

    describe('called by someone else', async () => {
      it('throws an exception', async () => {
        try {
          await instance.deleteAll({ from: accounts[1] });
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });
});
