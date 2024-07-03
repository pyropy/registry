// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Basic is a contract which implements a basic filePath/CID store.
contract FileRegistry {
    // A mapping from filePath to CID value.
    mapping (string => string) public items;

    event ItemSet(string indexed filePath, string cid);

    function save(string memory filePath, string memory cid) external {
        items[filePath] = cid;
        emit ItemSet(filePath, cid);
    }

    function get(string memory filePath) external view returns (string memory) {
        return items[filePath];
    }
}
