// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract DIDRegistry {

    address public owner;

    struct Issuer {
        string did;        // DID string (for audit / off-chain resolution)
        bool active;
    }

    mapping(address => Issuer) private issuers;

    event IssuerRegistered(address indexed issuer, string did);
    event IssuerDeactivated(address indexed issuer);
    event IssuerActivated(address indexed issuer);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    // Register a new issuer
    function registerIssuer(address issuer, string calldata did)
        external
        onlyOwner
    {
        require(issuer != address(0), "Invalid address");
        require(!issuers[issuer].active, "Issuer already active");

        issuers[issuer] = Issuer({
            did: did,
            active: true
        });

        emit IssuerRegistered(issuer, did);
    }

    // Deactivate an issuer (revokes future issue/revoke rights)
    function deactivateIssuer(address issuer)
        external
        onlyOwner
    {
        require(issuers[issuer].active, "Issuer not active");

        issuers[issuer].active = false;
        emit IssuerDeactivated(issuer);
    }

    // Reactivate an issuer
    function activateIssuer(address issuer)
        external
        onlyOwner
    {
        require(!issuers[issuer].active, "Issuer already active");

        issuers[issuer].active = true;
        emit IssuerActivated(issuer);
    }

    // --- Read functions ---

    function isValidIssuer(address issuer)
        external
        view
        returns (bool)
    {
        return issuers[issuer].active;
    }

    function getIssuerDID(address issuer)
        external
        view
        returns (string memory)
    {
        require(issuers[issuer].active, "Issuer not active");
        return issuers[issuer].did;
    }
}
