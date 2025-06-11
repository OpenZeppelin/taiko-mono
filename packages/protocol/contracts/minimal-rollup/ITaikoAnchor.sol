// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface TaikoAnchor {
    struct BlockHeader {
        bytes32 parentHash;
        bytes32 omnersHash;
        address coinbase;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        bytes logsBloom;
        uint256 difficulty;
        uint256 number;
        uint64 gasLimit;
        uint64 gasUsed;
        uint64 timestamp;
        bytes extraData;
        bytes32 mixedHash;
        uint64 nonce;
        bytes32 baseFeePerGas;
        bytes32 withdrawalsRoot;
        uint64 blobGasUsed;
        uint64 excessBlobGas;
        bytes32 parentBeaconBlockRoot;
        bytes32 requestsHash;
    }

    event Anchor(
        uint256 publicationId, uint256 anchorBlockId, bytes32 anchorBlockHash, bytes32 parentGasUsed
    );

    function anchor(
        uint256 _publicationId,
        uint256 _anchorBlockId,
        bytes32 _anchorBlockHash,
        BlockHeader calldata _anchorBlockHeader,
        bytes32 _parentGasUsed
    )
        external;

    function l1BlockHashes(uint256 blockId) external view returns (bytes32 blockHash);
}
