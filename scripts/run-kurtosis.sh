sudo service docker start

kurtosis clean -a
kurtosis run --enclave my-testnet github.com/ethpandaops/ethereum-package --args-file network_params.yaml