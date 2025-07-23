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
