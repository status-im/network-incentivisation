const Nodes = artifacts.require('Nodes');

const node1 = 'enode://da61e9eff86a56633b635f887d8b91e0ff5236bbc05b8169834292e92afb92929dcf6efdbf373a37903da8fe0384d5a0a8247e83f1ce211aa429200b6d28c548@47.91.156.93:443';

const node2 = 'enode://7de99e4cb1b3523bd26ca212369540646607c721ad4f3e5c821ed9148150ce6ce2e72631723002210fac1fd52dfa8bbdf3555e05379af79515e1179da37cc3db@35.188.19.210:443';

contract('Nodes', async (accounts) => {
  let instance;
  beforeEach(async () => {
    instance = await Nodes.deployed();
    await instance.deleteAll();
    await instance.toggleRegistration(false);
  });

  describe('addNode', async () => {
    describe('registering is enabled', async () => {
      beforeEach(async () => {
        await instance.toggleRegistration(true);
        await instance.addNode(node1);
        await instance.addNode(node2, { from: accounts[1] });
      });

      it('adds the first node from the owner', async () => {
        const actualNode1 = await instance.nodes(0);
        assert.equal(actualNode1, node1);
      });

      it('adds the second node from someone else', async () => {
        const actualNode2 = await instance.nodes(1);
        assert.equal(actualNode2, node2);
      });

      it('sets the count', async () => {
        const actualNodeCount = await instance.nodeCount();
        assert.equal(2, actualNodeCount);
      });
    });
    describe('registering is disabled', async () => {
      describe('called by the owner', async () => {
        beforeEach(async () => {
          await instance.addNode(node1);
          await instance.addNode(node2);
        });

        it('adds the first node', async () => {
          const actualNode1 = await instance.nodes(0);
          assert.equal(actualNode1, node1);
        });

        it('adds the second node', async () => {
          const actualNode2 = await instance.nodes(1);
          assert.equal(actualNode2, node2);
        });

        it('sets the count', async () => {
          const actualNodeCount = await instance.nodeCount();
          assert.equal(2, actualNodeCount);
        });
      });
      describe('called by someone else', async () => {
        it('throws an exception', async () => {
          try {
            await instance.addNode(node1, { from: accounts[1] });
          } catch (error) {
            return;
          }
          assert.fail('it should throw an exception');
        });
      });
    });
  });

  describe('deleteNode', async () => {
    describe('called by the owner', async () => {
      beforeEach(async () => {
        await instance.addNode(node1);
        await instance.addNode(node2);
        await instance.deleteNode(node1);
      });

      it('removes the first node', async () => {
        const actualNode2 = await instance.nodes(0);
        assert.equal(actualNode2, node2);
      });

      it('sets the count', async () => {
        const actualNodeCount = await instance.nodeCount();
        assert.equal(1, actualNodeCount);
      });
    });
    describe('called by someone else', async () => {
      it('throws an exception', async () => {
        try {
          await instance.deleteNode(node1, { from: accounts[1] });
        } catch (error) {
          return;
        }
        assert.fail('it should throw an exception');
      });
    });
  });

  describe('deleteAll', async () => {
    beforeEach(async () => {
      await instance.addNode(node1);
      await instance.addNode(node2);
      await instance.deleteAll();
    });

    it('empties the array', async () => {
      try {
        await instance.nodes(0);
      } catch (error) {
        return;
      }
      assert.fail('it should throw an exception');
    });

    it('sets the count', async () => {
      const actualNodeCount = await instance.nodeCount();
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
