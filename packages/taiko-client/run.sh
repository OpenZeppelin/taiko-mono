#!/bin/bash

set -e

OUTPUT=$(kurtosis run --enclave my-testnet github.com/ethpandaops/ethereum-package)

L1_WS_PORT=$(echo "$OUTPUT" | grep -A 10 "el-1-geth-lighthouse" | grep "ws:" | grep -oE "127\.0\.0\.1:[0-9]+" | head -n 1)
L1_BEACON_HTTP_PORT=$(echo "$OUTPUT" | grep -A 10 "cl-1-lighthouse-geth" | grep "http:" | grep -oE "127\.0\.0\.1:[0-9]+" | head -n 1)

if [[ -z "$L1_WS_PORT" || -z "$L1_BEACON_HTTP_PORT" ]]; then
  echo "❌ Could not extract ports."
  exit 1
fi

ENV_FILE=".env"

# Append new values
{
  echo "export L1_WS=ws://$L1_WS_PORT"
  echo "export L1_BEACON_URL=http://$L1_BEACON_HTTP_PORT"
} >> "$ENV_FILE"

echo "✅ Updated $ENV_FILE:"
grep -E 'L1_WS=|L1_BEACON_URL=' "$ENV_FILE"
