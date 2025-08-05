## Overview
- This repository exists to modify the Taiko client so it works with our contracts.
- I am finding the existing PRs difficult to follow. I intend to use this branch to make very incremental changes (probably copied from the other branches) so I can track what's going on.
- I will update this description as a I go, so any commit on this branch will explain my current understanding and the current state of the codebase.

## Background

### PR 1
- [Reference](https://github.com/OpenZeppelin/taiko-mono/pull/1) 
- As I understand it, this adds the minimal rollup contracts to the repository and modifies the taiko client to interact with them.
- It also includes a `run-me.md` file to explain how to set up a local environment.

### PR 3
- [Reference](https://github.com/OpenZeppelin/taiko-mono/pull/1)
- This uses the latest version of the minimal rollup contracts (with the publication feed as part of the inbox).
- It also starts the process of modifying the client to reorganise publications into blobs using the new format, and interpreting those blobs on the driver side.
- It also uses the new anchor contract on the L2, and updates the genesis state accordingly.
- It also renames the `run-me.md` file to `runme.md` and uses `kurtosis` for the L1 devnet instead of `anvil`.

## Changes on this branch

### Linux dev container
- I was having issues with `kurtosis` on my mac so I decided to try this in a linux dev container.
- Seppi's devcontainer (that he provided us in his workshop) was not suitable because it has a non-root user that cannot install new utilities.
- I decided to follow [this explanation](https://code.visualstudio.com/docs/devcontainers/create-dev-container) to create a fresh one, and to install all the relevant dependencies in there.
- Hopefully this also implies that other people can recreate my setup using the same Dockerfile.


### L1 Development Environment
- I copied (from PR 3)
   - the _L1 Node_ section of the `runme.md` file
   - the `network_params.yaml` file
- I placed them in the main directory (not _taiko-client_) because they're not specifically related to the client.
- I put the kurtosis install instructions in the `Dockerfile`. 
   - `kurtosis version` returns `CLI Version:   1.10.2`
- The `scripts/run-kurtosis.sh` script can be used to create a kurtosis enclave.
- I ran `cast wallet new` to produce a deployer account with a known private key, and prefunded this account in `network_params.yaml`

#### Local output
After running the script I get the error

```
Adding service with name 'blockscout-postgres' and image 'library/postgres:alpine'
There was an error executing Starlark code 
An error occurred executing instruction (number 44) at github.com/kurtosis-tech/postgres-package/main.star[115:40]:
  add_service(name="blockscout-postgres", config=ServiceConfig(image="library/postgres:alpine", ports={"postgresql": PortSpec(number=5432, application_protocol="postgresql")}, files={}, cmd=["-c", "max_connections=1000"], env_vars={"POSTGRES_DB": "blockscout", "POSTGRES_PASSWORD": "MyPassword1!", "POSTGRES_USER": "postgres"}, max_cpu=1000, min_cpu=10, max_memory=1024, min_memory=32, node_selectors={}))
  Caused by: Unexpected error occurred starting service 'blockscout-postgres'
  Caused by: An error occurred starting the user service container for user service with UUID '69c33437a7d44f1788fab39a826396e1'
  Caused by: Could not start Docker container from image 'library/postgres:alpine'.
  Caused by: Could not start Docker container with ID 'e5ae609cef20671cab7f65c46f4462ab8cfef23462603590297b7dfe5f65c555'; logs are below:
  --------------------- CONTAINER LOGS -----------------------
  
  ------------------- END CONTAINER LOGS --------------------
  Caused by: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: unable to apply cgroup configuration: cannot enter cgroupv2 "/sys/fs/cgroup/docker" with domain controllers -- it is in threaded mode: unknown

Error encountered running Starlark code.
```

and then the output

```
========================================== User Services ==========================================
UUID           Name                                             Ports                                         Status
26c40ed8583a   cl-1-lighthouse-geth                             http: 4000/tcp -> http://127.0.0.1:32779      RUNNING
                                                                metrics: 5054/tcp -> http://127.0.0.1:32780   
                                                                quic-discovery: 9001/udp -> 127.0.0.1:32771   
                                                                tcp-discovery: 9000/tcp -> 127.0.0.1:32781    
                                                                udp-discovery: 9000/udp -> 127.0.0.1:32770    
ddadbc095284   cl-2-lighthouse-geth                             http: 4000/tcp -> http://127.0.0.1:32782      RUNNING
                                                                metrics: 5054/tcp -> http://127.0.0.1:32783   
                                                                quic-discovery: 9001/udp -> 127.0.0.1:32773   
                                                                tcp-discovery: 9000/tcp -> 127.0.0.1:32784    
                                                                udp-discovery: 9000/udp -> 127.0.0.1:32772    
ae05ac94f991   el-1-geth-lighthouse                             engine-rpc: 8551/tcp -> 127.0.0.1:32771       RUNNING
                                                                metrics: 9001/tcp -> http://127.0.0.1:32772   
                                                                rpc: 8545/tcp -> 127.0.0.1:32769              
                                                                tcp-discovery: 30303/tcp -> 127.0.0.1:32773   
                                                                udp-discovery: 30303/udp -> 127.0.0.1:32768   
                                                                ws: 8546/tcp -> 127.0.0.1:32770               
62ca42aba9bf   el-2-geth-lighthouse                             engine-rpc: 8551/tcp -> 127.0.0.1:32776       RUNNING
                                                                metrics: 9001/tcp -> http://127.0.0.1:32777   
                                                                rpc: 8545/tcp -> 127.0.0.1:32774              
                                                                tcp-discovery: 30303/tcp -> 127.0.0.1:32778   
                                                                udp-discovery: 30303/udp -> 127.0.0.1:32769   
                                                                ws: 8546/tcp -> 127.0.0.1:32775               
2743a532c715   validator-key-generation-cl-validator-keystore   <none>                                        RUNNING
3a6c03ce33ea   vc-1-geth-lighthouse                             metrics: 8080/tcp -> http://127.0.0.1:32785   RUNNING
96dcdfff0a9a   vc-2-geth-lighthouse                             metrics: 8080/tcp -> http://127.0.0.1:32786   RUNNING
```

I will ignore the blockscout error for now.