// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./DIDRegistry.sol";

contract CertificateRegistry {

    DIDRegistry public didRegistry;

    struct Certificate {
        address issuer;
        bool isValid;
        bool exists;
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

    // Issue a new certificate
    function issueCertificate(bytes32 _certHash) public onlyValidIssuer {
        require(!certificates[_certHash].exists, "Certificate already exists");

        certificates[_certHash] = Certificate({
            issuer: msg.sender,
            isValid: true,
            exists: true
        });

        emit CertificateIssued(_certHash, msg.sender);
    }

    // Revoke an existing certificate
    function revokeCertificate(bytes32 _certHash) public onlyValidIssuer {
        require(certificates[_certHash].exists, "Certificate does not exist");
        require(certificates[_certHash].isValid, "Certificate already revoked");

        certificates[_certHash].isValid = false;

        emit CertificateRevoked(_certHash);
    }

    // Verify certificate details
    function verifyCertificate(bytes32 _certHash)
        public
        view
        returns (
            bool exists,
            address issuer,
            bool isValid
        )
    {
        Certificate memory cert = certificates[_certHash];
        return (cert.exists, cert.issuer, cert.isValid);
    }
}
