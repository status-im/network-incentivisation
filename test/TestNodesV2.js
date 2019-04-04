const Nodes = artifacts.require('NodesV2');

function ipToDecimal(dot) {
  const d = dot.split('.');
  return ((((((+d[0]) * 256) + (+d[1])) * 256) + (+d[2])) * 256) + (+d[3]);
}

/* function decimalToIp(num) {
  let d = num % 256;
  for (let i = 3; i > 0; i = i - 1) {
    num = Math.floor(num / 256);
    d = `${num % 256}.${d}`;
  }
  return d;
} */

const node1 = {
  pk: '0xaf80b90d25145da28c583359beb47b21796b2fe1a23c1511e443e7a64dfdb27d7434c380f0aa4c500e220aa1a9d068514b1ff4d5019e624e7ba1efe82b340a59',
  ip: ipToDecimal('47.91.156.93'),
  port: 443,
};

const node2 = {
  pk: '0xce7edc292d7b747fab2f23584bbafaffde5c8ff17cf689969614441e0527b90015ea9fee96aed6d9c0fc2fbe0bd1883dee223b3200246ff1e21976bdbc9a0fc8',
  ip: ipToDecimal('35.188.19.210'),
  port: 443,
};

/* const node3 = {
  pk: '0xebefab39b69bbbe64d8cd86be765b3be356d8c4b
  24660f65d493143a0c44f38c85a257300178f7845592a1b0332811542e9a58281c835babdd7535babb64efc1',
  ip: ipToDecimal('35.202.99.224'),
  port: 443
}; */

contract('Nodes', async (accounts) => {
  let instance;
  beforeEach(async () => {
    instance = await Nodes.new();
  });

  describe('registerNode', async () => {
    beforeEach(async () => {
      await instance.registerNode(node1.pk, node1.ip, node1.port);
      await instance.registerNode(node2.pk, node2.ip, node2.port, { from: accounts[1] });
    });

    it('adds the nodes', async () => {
      const actualNodeCount = await instance.pendingNodeCount();
      assert.equal(2, actualNodeCount);
    });

    describe('node is already in pending', async () => {
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
});
