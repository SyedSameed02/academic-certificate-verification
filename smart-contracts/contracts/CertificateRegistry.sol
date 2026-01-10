// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IDIDRegistry {
    function isValidIssuer(address issuer) external view returns (bool);
}

contract CertificateRegistry {

    struct Certificate {
        address issuer;
        uint256 issuedAt;
        bool revoked;
    }

    mapping(bytes32 => Certificate) private certificates;

    IDIDRegistry public didRegistry;

    event CertificateIssued(bytes32 indexed certHash, address indexed issuer);
    event CertificateRevoked(bytes32 indexed certHash, address indexed issuer);

    constructor(address _didRegistry) {
        require(_didRegistry != address(0), "Invalid DIDRegistry address");
        didRegistry = IDIDRegistry(_didRegistry);
    }

    function issueCertificate(bytes32 certHash) external {
        require(didRegistry.isValidIssuer(msg.sender), "Unauthorized issuer");
        require(certHash != bytes32(0), "Invalid hash");
        require(certificates[certHash].issuer == address(0), "Certificate exists");

        certificates[certHash] = Certificate({
            issuer: msg.sender,
            issuedAt: block.timestamp,
            revoked: false
        });

        emit CertificateIssued(certHash, msg.sender);
    }

    function revokeCertificate(bytes32 certHash) external {
        require(didRegistry.isValidIssuer(msg.sender), "Unauthorized issuer");

        Certificate storage cert = certificates[certHash];
        require(cert.issuer != address(0), "Certificate not found");
        require(cert.issuer == msg.sender, "Only issuer can revoke");
        require(!cert.revoked, "Already revoked");

        cert.revoked = true;

        emit CertificateRevoked(certHash, msg.sender);
    }

    function getCertificate(bytes32 certHash)
        external
        view
        returns (address issuer, bool revoked, uint256 issuedAt)
    {
        Certificate storage cert = certificates[certHash];
        require(cert.issuer != address(0), "Certificate not found");

        return (cert.issuer, cert.revoked, cert.issuedAt);
    }

    function exists(bytes32 certHash) external view returns (bool) {
        return certificates[certHash].issuer != address(0);
    }
}
