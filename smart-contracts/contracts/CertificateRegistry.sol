// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./DIDRegistry.sol";

contract CertificateRegistry {

    DIDRegistry public didRegistry;

    struct Certificate {
        bytes32 certHash;
        address issuer;
        bool isValid;
    }

    mapping(bytes32 => Certificate) public certificates;

    event CertificateIssued(bytes32 certHash, address issuer);
    event CertificateRevoked(bytes32 certHash);

    constructor(address _didRegistryAddress) {
        didRegistry = DIDRegistry(_didRegistryAddress);
    }

    modifier onlyValidIssuer() {
        require(
            didRegistry.isValidIssuer(msg.sender),
            "Issuer not authorized"
        );
        _;
    }

    function issueCertificate(bytes32 _certHash) public onlyValidIssuer {
        certificates[_certHash] = Certificate(
            _certHash,
            msg.sender,
            true
        );

        emit CertificateIssued(_certHash, msg.sender);
    }

    function revokeCertificate(bytes32 _certHash) public onlyValidIssuer {
        require(certificates[_certHash].isValid, "Already revoked");
        certificates[_certHash].isValid = false;
        emit CertificateRevoked(_certHash);
    }

    function verifyCertificate(bytes32 _certHash) public view returns (bool) {
        return certificates[_certHash].isValid;
    }
}
