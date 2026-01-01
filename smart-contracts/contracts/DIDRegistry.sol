// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract DIDRegistry {

    address public owner;

    struct Issuer {
        string did;
        bool isRegistered;
    }

    mapping(address => Issuer) public issuers;

    event IssuerRegistered(address issuer, string did);
    event IssuerRemoved(address issuer);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner allowed");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function registerIssuer(address _issuer, string memory _did) public onlyOwner {
        issuers[_issuer] = Issuer(_did, true);
        emit IssuerRegistered(_issuer, _did);
    }

    function removeIssuer(address _issuer) public onlyOwner {
        issuers[_issuer].isRegistered = false;
        emit IssuerRemoved(_issuer);
    }

    function isValidIssuer(address _issuer) public view returns (bool) {
        return issuers[_issuer].isRegistered;
    }
}
