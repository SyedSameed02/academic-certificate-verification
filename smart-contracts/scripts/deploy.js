const hre = require("hardhat");
const fs = require("fs");

async function main() {
  console.log("Deploying contracts...");

  // 1. Deploy DIDRegistry
  const DIDRegistry = await hre.ethers.getContractFactory("DIDRegistry");
  const didRegistry = await DIDRegistry.deploy();
  await didRegistry.waitForDeployment();

  const didRegistryAddress = await didRegistry.getAddress();
  console.log("DIDRegistry deployed to:", didRegistryAddress);

  // 2. Deploy CertificateRegistry (pass DIDRegistry address)
  const CertificateRegistry = await hre.ethers.getContractFactory("CertificateRegistry");
  const certificateRegistry = await CertificateRegistry.deploy(didRegistryAddress);
  await certificateRegistry.waitForDeployment();

  const certificateRegistryAddress = await certificateRegistry.getAddress();
  console.log("CertificateRegistry deployed to:", certificateRegistryAddress);

  // 3. Save addresses to config file
  const config = {
    didRegistry: didRegistryAddress,
    certificateRegistry: certificateRegistryAddress,
  };

  fs.mkdirSync("../academic-certificate-verification-backend/config", { recursive: true });
  fs.writeFileSync(
    "../academic-certificate-verification-backend/config/contracts.json",
    JSON.stringify(config, null, 2)
  );

  console.log("Contract addresses saved to contracts.json");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
