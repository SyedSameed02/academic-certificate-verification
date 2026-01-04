
## About

This folder contains Ethereum smart contracts for academic certificate verification.
The contracts ensure secure, tamper-proof issuance and verification of certificates
using blockchain technology. They are deployed using Hardhat.

## Smart Contracts

1. DIDRegistry.sol  
2. CertificateRegistry.sol

### DIDRegistry.sol
- Registers authorized institutions
- Prevents unauthorized certificate issuance
- Used to verify issuer identity

### CertificateRegistry.sol
- Stores hashes of issued certificates
- Allows only registered institutions to issue certificates
- Enables public certificate verification

CertificateRegistry depends on DIDRegistry to verify whether an institution
is authorized before issuing a certificate.

## Deployment

1. Install dependencies
   npm install

2. Start local blockchain
   npx hardhat node

3. Deploy contracts
   npx hardhat run scripts/deploy.js --network localhost

The deployment script outputs the deployed contract addresses.
These addresses are used by the backend application to interact
with the smart contracts.

## Flow
Institution → DIDRegistry → CertificateRegistry → Verifier

