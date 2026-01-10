const { ethers } = require("hardhat");
const { expect } = require("chai");

describe("Academic Certificate Verification Flow", function () {
  let didRegistry, certRegistry;
  let owner, issuer1, issuer2, randomUser;

  const CERT_HASH = ethers.keccak256(
    ethers.toUtf8Bytes("student-certificate-001")
  );

  beforeEach(async function () {
    [owner, issuer1, issuer2, randomUser] = await ethers.getSigners();

    // Deploy DIDRegistry
    const DIDRegistry = await ethers.getContractFactory("DIDRegistry");
    didRegistry = await DIDRegistry.deploy();
    await didRegistry.waitForDeployment();

    // Deploy CertificateRegistry (wired)
    const CertificateRegistry = await ethers.getContractFactory("CertificateRegistry");
    certRegistry = await CertificateRegistry.deploy(
      await didRegistry.getAddress()
    );
    await certRegistry.waitForDeployment();
  });

  /* ---------------- DID REGISTRY ---------------- */

  it("Owner can register issuer", async function () {
    await didRegistry.registerIssuer(issuer1.address, "did:univ:issuer1");
    expect(await didRegistry.isValidIssuer(issuer1.address)).to.equal(true);
  });

  it("Non-owner cannot register issuer", async function () {
    await expect(
      didRegistry.connect(issuer1).registerIssuer(
        issuer1.address,
        "did:fake"
      )
    ).to.be.revertedWith("Only owner");
  });

  /* ---------------- CERTIFICATE ISSUANCE ---------------- */

  it("Valid issuer can issue certificate", async function () {
    await didRegistry.registerIssuer(issuer1.address, "did:univ:issuer1");

    await certRegistry.connect(issuer1).issueCertificate(CERT_HASH);

    const cert = await certRegistry.getCertificate(CERT_HASH);
    expect(cert.issuer).to.equal(issuer1.address);
    expect(cert.revoked).to.equal(false);
  });

  it("Invalid issuer cannot issue certificate", async function () {
    await expect(
      certRegistry.connect(issuer1).issueCertificate(CERT_HASH)
    ).to.be.revertedWith("Unauthorized issuer");
  });

  it("Cannot issue duplicate certificate", async function () {
    await didRegistry.registerIssuer(issuer1.address, "did:univ:issuer1");

    await certRegistry.connect(issuer1).issueCertificate(CERT_HASH);

    await expect(
      certRegistry.connect(issuer1).issueCertificate(CERT_HASH)
    ).to.be.revertedWith("Certificate exists");
  });

  /* ---------------- CERTIFICATE REVOCATION ---------------- */

  it("Issuer can revoke their certificate", async function () {
    await didRegistry.registerIssuer(issuer1.address, "did:univ:issuer1");

    await certRegistry.connect(issuer1).issueCertificate(CERT_HASH);
    await certRegistry.connect(issuer1).revokeCertificate(CERT_HASH);

    const cert = await certRegistry.getCertificate(CERT_HASH);
    expect(cert.revoked).to.equal(true);
  });

  it("Other issuer cannot revoke certificate", async function () {
    await didRegistry.registerIssuer(issuer1.address, "did:univ:issuer1");
    await didRegistry.registerIssuer(issuer2.address, "did:univ:issuer2");

    await certRegistry.connect(issuer1).issueCertificate(CERT_HASH);

    await expect(
      certRegistry.connect(issuer2).revokeCertificate(CERT_HASH)
    ).to.be.revertedWith("Only issuer can revoke");
  });

  it("Random user cannot revoke certificate", async function () {
    await didRegistry.registerIssuer(issuer1.address, "did:univ:issuer1");

    await certRegistry.connect(issuer1).issueCertificate(CERT_HASH);

    await expect(
      certRegistry.connect(randomUser).revokeCertificate(CERT_HASH)
    ).to.be.revertedWith("Unauthorized issuer");
  });
});
