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

- So far, I have just updated this markdown file.
