1. Run a local L1 execution node using anvil
```sh
anvil --fork-url https://eth-holesky.g.alchemy.com/v2/sNeHddUXY-9axjz5eRRDg0uox6Q7LijO
```
By default this spins up a node at `127.0.0.1:8545` and the default chain ID is `31337`

2. Deploy the L1 contracts
```sh
TODO: Write a script that deploys the `PublicationFeed` at least

3. Run a local L2 execution client(`taiko-geth`). Inside the `taiko-geth` project
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

This will spin up a node at `127.0.0.1:8551` both for RPC and ws connections.