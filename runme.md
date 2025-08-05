# Local Environment Setup

These are the instructions to run a full L1 & L2 environment locally. There are likely easier ways to do this, but this is what worked for me :)

## L1 Node

1. Use kurtosis to run a local L1 devnet. Please refer to the [kurtosis documentation](https://github.com/ethpandaops/ethereum-package) for further details. However, in short this spins up a devnet with 2 execution client (EL) and 2 consensus nodes (CL) as well as some nice QoL features such as a block and beacon chain explorer.

   ```sh
   kurtosis run --enclave my-testnet github.com/ethpandaops/ethereum-package --args-file network_params.yaml
   ```

You will get an output similar to this:

```bash
UUID           Name                   Ports                                         Status
681cd3d3a7a0   cl-1-lighthouse-geth   http: 4000/tcp -> http://127.0.0.1:32776      RUNNING
                                      metrics: 5054/tcp -> http://127.0.0.1:32777
                                      quic-discovery: 9001/udp -> 127.0.0.1:32770
                                      tcp-discovery: 9000/tcp -> 127.0.0.1:32778
                                      udp-discovery: 9000/udp -> 127.0.0.1:32769
d61be60e1795   el-1-geth-lighthouse   engine-rpc: 8551/tcp -> 127.0.0.1:32773       RUNNING
                                      metrics: 9001/tcp -> http://127.0.0.1:32774
                                      rpc: 8545/tcp -> 127.0.0.1:32771
                                      tcp-discovery: 30303/tcp -> 127.0.0.1:32775
                                      udp-discovery: 30303/udp -> 127.0.0.1:32768
                                      ws: 8546/tcp -> 127.0.0.1:32772

```

Replace the `.env`

```bash
# ws from el-1-geth-lighthouse
export L1_WS=ws://127.0.0.1:32772

# http from cl-1-lighthouse-geth
export L1_BEACON_URL=http://127.0.0.1:32776
```
To interact with the EL you can use the rpc endpoint (i.e when deploying a contracts), in this cause it will be `http://127.0.0.1:32771`