require("dotenv").config({ path: "./backend-go/.env" });
const { ethers } = require("ethers");
const fs = require("fs");
const path = require("path");

async function main() {
  const RPC_URL = process.env.RPC_URL;
  const PRIVATE_KEY = process.env.PRIVATE_KEY;
  const DID_REGISTRY_ADDRESS = process.env.DID_REGISTRY_ADDRESS;

  if (!RPC_URL || !PRIVATE_KEY || !DID_REGISTRY_ADDRESS) {
    throw new Error("Missing env vars: RPC_URL / PRIVATE_KEY / DID_REGISTRY_ADDRESS");
  }

  const provider = new ethers.JsonRpcProvider(RPC_URL);
  const wallet = new ethers.Wallet(PRIVATE_KEY, provider);

  const issuer = wallet.address;
  const did = `did:ethr:${issuer}`; // ✅ DID string

  // Load ABI from artifact
  const abiPath = path.join(
    __dirname,
    "../smart-contracts/artifacts/contracts/DIDRegistry.sol/DIDRegistry.json"
  );
  const artifact = JSON.parse(fs.readFileSync(abiPath, "utf8"));

  const didRegistry = new ethers.Contract(
    DID_REGISTRY_ADDRESS,
    artifact.abi,
    wallet
  );

  console.log("DID Registry:", DID_REGISTRY_ADDRESS);
  console.log("Registering issuer:", issuer);
  console.log("DID:", did);

  // ✅ correct call: 2 params
  const tx = await didRegistry.registerIssuer(issuer, did);
  console.log("Tx sent:", tx.hash);

  await tx.wait();
  console.log("✅ Issuer registered successfully");

  // Optional: activate issuer if needed
  // const tx2 = await didRegistry.activateIssuer(issuer);
  // await tx2.wait();
  // console.log("✅ Issuer activated");
}

main().catch((err) => {
  console.error("❌ Error:", err);
  process.exit(1);
});
