In order to run the devnet you need to do the following:

1. Install kurtosis:

```bash
brew install kurtosis-tech/tap/kurtosis-cli
```

2. Have docker running

3. Run the devnet:

```bash
kurtosis run --enclave my-testnet github.com/ethpandaops/ethereum-package --args-file network_params.yaml
```

Note: The command itself might run relatively quickly however it will take around 45 minutes to fully sync the devnet.

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

4. Replace the `.env`

```bash
# ws from el-1-geth-lighthouse
export L1_WS=ws://127.0.0.1:32772

# http from cl-1-lighthouse-geth
export L1_BEACON_URL=http://127.0.0.1:32776
```

5. To deploy the contract use the rpc endpoint from the output of the devnet (in this case `http://127.0.0.1:32771`)

Currently im having issues getting 'prefunded' accounts to work with the devnet. So for now use a private key from metamask and fund it with some eth.


If there are any issues or you just want to stop the devnet please run:

```bash
kurtosis enclave rm -f my-testnet
```

I recommend doing this before you start the devnet again.

## Configuration

Please refer to the [README](https://github.com/ethpandaops/ethereum-package/blob/main/README.md#configuration) for more information.
