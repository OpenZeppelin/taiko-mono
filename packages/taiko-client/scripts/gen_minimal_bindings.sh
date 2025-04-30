#!/bin/bash

# Generate go contract bindings for minimal-rollup interfaces.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

echo ""
echo "TAIKO_GETH_DIR: ${TAIKO_GETH_DIR}"
echo ""

cd ${TAIKO_GETH_DIR} &&
  make all &&
  cd -

cd ../protocol &&
  pnpm clean &&
  pnpm compile:minimal-rollup &&
  cd -

ABIGEN_BIN=$TAIKO_GETH_DIR/build/bin/abigen
PACKAGE_NAME=minimal

echo ""
echo "Start generating Go bindings for minimal-rollup interfaces..."
echo ""

# Create the minimal directory if it doesn't exist
mkdir -p $DIR/../bindings/${PACKAGE_NAME}

# Find all interface files in the minimal-rollup directory
INTERFACES=$(find ../protocol/contracts/minimal-rollup -name "I*.sol" -type f -exec basename {} .sol \;)

# Generate bindings for each interface
for interface in $INTERFACES; do
  echo "Generating bindings for $interface..."
  cat ../protocol/out/minimal-rollup/${interface}.sol/${interface}.json |
    jq .abi |
    ${ABIGEN_BIN} --abi - --type ${interface} --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_$(echo ${interface#I} | tr '[:upper:]' '[:lower:]').go
done

git -C ../../ log --format="%H" -n 1 >./bindings/${PACKAGE_NAME}/.githead

echo "🍻 Go contract bindings for minimal-rollup interfaces generated!" 