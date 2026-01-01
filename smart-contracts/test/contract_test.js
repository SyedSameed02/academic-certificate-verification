const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Academic Certificate Verification Smart Contracts", function () {

  let DIDRegistry;
  let didRegistry;
  let CertificateRegistry;
  let certificateRegistry;

  let owner;
  let issuer;
  let attacker;

  let certificateHash;

  beforeEach(async function () {
    [owner, issuer, attacker] = await ethers.getSigners();

    // Deploy DIDRegistry
    DIDRegistry = await ethers.getContractFactory("DIDRegistry");
    didRegistry = await DIDRegistry.connect(owner).deploy();
    await didRegistry.waitForDeployment();

    // Deploy CertificateRegistry with DIDRegistry address
    CertificateRegistry = await ethers.getContractFactory("CertificateRegistry");
    certificateRegistry = await CertificateRegistry
      .connect(owner)
      .deploy(await didRegistry.getAddress());
    await certificateRegistry.waitForDeployment();

    // Sample certificate hash
    certificateHash = ethers.keccak256(
      ethers.toUtf8Bytes("SAMEED-22691A05J6-BTECH-CSE")
    );
  });

  // ---------------- DID REGISTRY TESTS ----------------

  it("Owner should register an issuer DID", async function () {
    await didRegistry
      .connect(owner)
      .registerIssuer(issuer.address, "did:university:mits");

    expect(await didRegistry.isValidIssuer(issuer.address)).to.equal(true);
  });

  it("Non-owner should NOT register issuer", async function () {
    await expect(
      didRegistry
        .connect(attacker)
        .registerIssuer(attacker.address, "did:fake:123")
    ).to.be.revertedWith("Only owner allowed");
  });

  // ---------------- CERTIFICATE ISSUANCE TESTS ----------------

  it("Registered issuer should issue certificate", async function () {
    await didRegistry
      .connect(owner)
      .registerIssuer(issuer.address, "did:university:mits");

    await certificateRegistry
      .connect(issuer)
      .issueCertificate(certificateHash);

    const isValid = await certificateRegistry.verifyCertificate(certificateHash);
    expect(isValid).to.equal(true);
  });

  it("Unregistered issuer should NOT issue certificate", async function () {
    await expect(
      certificateRegistry
        .connect(attacker)
        .issueCertificate(certificateHash)
    ).to.be.revertedWith("Issuer not authorized");
  });

  // ---------------- CERTIFICATE REVOCATION TESTS ----------------

  it("Issuer should revoke issued certificate", async function () {
    await didRegistry
      .connect(owner)
      .registerIssuer(issuer.address, "did:university:mits");

    await certificateRegistry
      .connect(issuer)
      .issueCertificate(certificateHash);

    await certificateRegistry
      .connect(issuer)
      .revokeCertificate(certificateHash);

    const isValid = await certificateRegistry.verifyCertificate(certificateHash);
    expect(isValid).to.equal(false);
  });

  it("Revoking an already revoked certificate should fail", async function () {
    await didRegistry
      .connect(owner)
      .registerIssuer(issuer.address, "did:university:mits");

    await certificateRegistry
      .connect(issuer)
      .issueCertificate(certificateHash);

    await certificateRegistry
      .connect(issuer)
      .revokeCertificate(certificateHash);

    await expect(
      certificateRegistry
        .connect(issuer)
        .revokeCertificate(certificateHash)
    ).to.be.revertedWith("Already revoked");
  });

  // ---------------- SECURITY TESTS ----------------

  it("Verifier should only read certificate validity", async function () {
    await didRegistry
      .connect(owner)
      .registerIssuer(issuer.address, "did:university:mits");

    await certificateRegistry
      .connect(issuer)
      .issueCertificate(certificateHash);

    const result = await certificateRegistry
      .connect(attacker)
      .verifyCertificate(certificateHash);

    expect(result).to.equal(true);
  });
});
