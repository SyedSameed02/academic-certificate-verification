const hre = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
  console.log("🚀 Deploying smart contracts...");

  // 1️⃣ Deploy DIDRegistry (no constructor args)
  const DIDRegistry = await hre.ethers.getContractFactory("DIDRegistry");
  const didRegistry = await DIDRegistry.deploy();
  await didRegistry.waitForDeployment();

  const didRegistryAddress = await didRegistry.getAddress();
  console.log("✅ DIDRegistry deployed to:", didRegistryAddress);

  // 2️⃣ Deploy CertificateRegistry (PASS DIDRegistry ADDRESS)
  const CertificateRegistry = await hre.ethers.getContractFactory(
    "CertificateRegistry"
  );

  const certificateRegistry = await CertificateRegistry.deploy(
    didRegistryAddress // 👈 REQUIRED CONSTRUCTOR ARG
  );

  await certificateRegistry.waitForDeployment();
  const certificateRegistryAddress =
    await certificateRegistry.getAddress();

  console.log(
    "✅ CertificateRegistry deployed to:",
    certificateRegistryAddress
  );

  // 3️⃣ Write config for Go backend
  const configDir = path.join(
    __dirname,
    "../../backend-go/config"
  );

  if (!fs.existsSync(configDir)) {
    fs.mkdirSync(configDir, { recursive: true });
  }

  const contractsConfig = {
    didRegistry: didRegistryAddress,
    certificateRegistry: certificateRegistryAddress,
  };

  fs.writeFileSync(
    path.join(configDir, "contracts.json"),
    JSON.stringify(contractsConfig, null, 2)
  );

  console.log("📄 contracts.json written for Go backend");
}

main()
  .then(() => {
    console.log("🎉 Deployment complete");
    process.exit(0);
  })
  .catch((error) => {
    console.error("❌ Deployment failed");
    console.error(error);
    process.exit(1);
  });
