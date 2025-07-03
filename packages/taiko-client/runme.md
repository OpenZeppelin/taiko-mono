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


2. Deploy the minimal-rollup L1 contracts.


   From the `minimal-rollup` repository checkout the `client-work` branch and run the following command:

   ```sh
   forge script script/DeployTaikoInbox.s.sol:DeployTaikoInbox --rpc-url <RPC_URL> --private-key bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31 --broadcast
   ```

   Note: Safe to share the PK because its one of the default kurtosis ones

## L2 Execution Client

1.  Run a local L2 execution client(`taiko-geth`).
    First make sure you generate a JWT secret and save it as `jwt.txt` as explained [here](https://docs.taiko.xyz/guides/node-operators/run-a-node-for-taiko-alethia/).
Inside the `taiko-geth` project:

```sh
make all
```

 ```sh
        ./build/bin/geth \
        --taiko \
        --networkid 167001 \
        --gcmode archive \
        --datadir ./data/taiko-geth \
        --metrics \
        --metrics.expensive \
        --metrics.addr "0.0.0.0" \
        --bootnodes enode://7a8955b27eda2ddf361b59983fce9c558b18ad60d996ac106629f7f913247ef13bc842c7cf6ec6f87096a3ea8048b04873c40d3d873c0276d38e222bddd72e88@43.153.44.186:30303,enode://704a50da7e727aa10c45714beb44ece04ca1280ad63bb46bb238a01bf55c19c9702b469fb12c63824fa90f5051f7091b1c5069df1ec9a0ba1e943978c09d270f@49.51.202.127:30303,enode://f52e4e212a15cc4f68df27282e616d51d7823596c83c8c8e3b3416d7ab531cefc7b8a493d01964e1918315e6b0c7a4806634aeabb9013642a9159a53f4ebc094@43.153.16.47:30303,enode://57f4b29cd8b59dc8db74be51eedc6425df2a6265fad680c843be113232bbe632933541678783c2a5759d65eac2e2241c45a34e1c36254bccfe7f72e52707e561@104.197.107.1:30303,enode://87a68eef46cc1fe862becef1185ac969dfbcc050d9304f6be21599bfdcb45a0eb9235d3742776bc4528ac3ab631eba6816e9b47f6ee7a78cc5fcaeb10cd32574@35.232.246.122:30303 \
        --authrpc.addr "0.0.0.0" \
        --authrpc.port 28551 \
        --authrpc.vhosts "*" \
        --authrpc.jwtsecret ./jwt.txt \
        --http \
        --http.api admin,debug,eth,net,web3,txpool,miner,taiko \
        --http.addr "0.0.0.0" \
        --http.port 28545 \
        --http.vhosts "*" \
        --ws \
        --ws.api admin,debug,eth,net,web3,txpool,miner,taiko \
        --ws.addr "0.0.0.0" \
        --ws.port 28546 \
        --ws.origins "*" \
        --gpo.defaultprice "10000000" \
        --port 30304 \
        --syncmode full \
        --maxpeers 0
        ```


This will spin up a node for both RPC and ws connections. These ports should be the same as on the one present in the `.env.example` file.

## L2 consensus client

Now it is time to spin up `taiko-client`. We'll do this for `driver` mode, which means the client will be able to listen to events from the L1 and send it to the execution client for processing L2 blocks.

1. First, build the project

   ```sh
   make build
   ```

1. Now fill all the environment variables in an `.env` file and source it. You can find an example in `.env.example`

   ```sh
   source .env
   ```

1. Copy the `jwt.txt` file that you used for `taiko-geth` to the `taiko-client` project.

1. Now you can start the client

   ```sh
   ./bin/taiko-client driver --jwtSecret jwt.txt
   ```

   To test it out send a `publish` transaction to the L1 node. This emits a `Published` event, that if everything has been configured correctly should be picked up by our client and processed.

   ```sh
   cast send  0x703848F4c85f18e3acd8196c8eC91eb0b7Bd0797 "publish(uint256,uint64)" 1 100 --private-key  bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31  --rpc-url <RPC_URL>
   ```
