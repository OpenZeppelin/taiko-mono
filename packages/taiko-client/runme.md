# Local Environment Setup

These are the instructions to run a full L1 & L2 environment locally. There are likely easier ways to do this, but this is what worked for me :)

## L1 Node

1. Run a local L1 execution node using anvil. We fork Holesky since that is where the Taiko L1 contracts were deployed, and parts of the codebase still need them.

   ```sh
   anvil --fork-url https://eth-holesky.g.alchemy.com/v2/<yourKey>
   ```

   By default this spins up a node at `127.0.0.1:8545` and the default chain ID is `31337`

1. Deploy the minimal-rollup L1 contracts.
   TODO: Write a script that deploys the `PublicationFeed` at least. For now, we just deploy manually using `forge`.

   From the `minimal-rollup` repository:

   ```sh
   forge create src/protocol/PublicationFeed.sol:PublicationFeed --private-key YOUR_PRIVATE_KEY --rpc-url http://127.0.0.1:8545
   ```

## L2 Execution Client

1.  Run a local L2 execution client(`taiko-geth`).
    First make sure you generate a JWT secret and save it as `jwt.txt` as explained [here](https://docs.taiko.xyz/guides/node-operators/run-a-node-for-taiko-alethia/).

        Inside the `taiko-geth` project:

        ```sh
        make all
        ```

        ```sh
        ./build/bin/geth \
            --dev \
            --taiko \
            --authrpc.addr "0.0.0.0" \
            --authrpc.port 28551 \
            --authrpc.vhosts "*" \
            --http \
            --http.api admin,debug,eth,net,web3,txpool,miner,taiko \
            --http.addr "0.0.0.0" \
            --http.port 28545 \
            --ws \
            --ws.api admin,debug,eth,net,web3,txpool,miner,taiko \
            --ws.addr "0.0.0.0" \
            --ws.port 28546 \
            --ws.origins "*" \
            --authrpc.addr=0.0.0.0 \
            --authrpc.vhosts="*" \
            --authrpc.jwtsecret ./jwt.txt
            --syncmode full
        ```

        This will spin up a node for both RPC and ws connections.

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
   cast send <PUBLICATION_FEED_ADDRESS> "publish(bytes[])" "[0x1234, 0x5678]" --private-key <YOUR_PRIVATE_KEY> --rpc-url http://127.0.0.1:8545
   ```
