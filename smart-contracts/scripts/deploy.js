const hre = require("hardhat");

async function main() {
  console.log("Starting deployment...");

  // Get deployer account
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deployer address:", deployer.address);
  console.log(
    "Deployer balance:",
    (await hre.ethers.provider.getBalance(deployer.address)).toString()
  );

  // ---------------- Deploy DIDRegistry ----------------
  const DIDRegistry = await hre.ethers.getContractFactory("DIDRegistry");
  const didRegistry = await DIDRegistry.deploy();
  await didRegistry.waitForDeployment();

  const didRegistryAddress = await didRegistry.getAddress();
  console.log("DIDRegistry deployed at:", didRegistryAddress);

  // ---------------- Deploy CertificateRegistry ----------------
  const CertificateRegistry = await hre.ethers.getContractFactory(
    "CertificateRegistry"
  );
  const certificateRegistry = await CertificateRegistry.deploy(
    didRegistryAddress
  );
  await certificateRegistry.waitForDeployment();

  const certificateRegistryAddress =
    await certificateRegistry.getAddress();
  console.log(
    "CertificateRegistry deployed at:",
    certificateRegistryAddress
  );

  console.log("🎉 Deployment completed successfully!");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Deployment failed:", error);
    process.exit(1);
  });
