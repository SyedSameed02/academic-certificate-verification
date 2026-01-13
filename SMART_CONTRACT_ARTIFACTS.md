# SMART_CONTRACT_ARTIFACTS.md — Generate & Use Hardhat Contract Artifacts

This guide explains how to generate smart contract artifacts using Hardhat:
- Compile Solidity contracts
- Start local Hardhat node
- Deploy contracts
- Export deployed addresses / ABIs
- Register and activate issuer
- Verify deployment outputs

> Run commands inside the **smart-contracts/** folder unless specified.

---

## 0) Go to smart-contracts folder

```powershell
cd C:\Users\shifa\OneDrive\Desktop\sameed\academic-certificate-verification\smart-contracts
```

---

## 1) Install dependencies (first time only)

```powershell
npm install
```

---

## 2) Compile contracts (generate artifacts)

```powershell
npx hardhat compile
```

Artifacts generated under:
- `smart-contracts/artifacts/`
- `smart-contracts/cache/`

> `artifacts/` includes the ABI + bytecode used by deployment scripts and external apps.

---

## 3) Start local blockchain (Hardhat node)

Open terminal 1:

```powershell
npx hardhat node
```

Keep it running. It starts at:
- RPC: `http://127.0.0.1:8545`
- Chain ID: `31337`

---

## 4) Deploy contracts on localhost

Open terminal 2 (same folder):

```powershell
npx hardhat run scripts/deploy.js --network localhost
```

Deployment prints the contract addresses for:
- DIDRegistry
- CertificateRegistry

Copy these addresses into backend `.env`:

Example:

```env
DID_REGISTRY_ADDRESS=0x...
CERTIFICATE_REGISTRY_ADDRESS=0x...
```

---

## 5) Export ABI (optional)

Hardhat artifacts already contain ABI, but to extract clean ABI JSON you can do:

### Option A: Use the generated artifact
ABI is inside:

```
smart-contracts/artifacts/contracts/<ContractFile.sol>/<ContractName>.json
```

Example:
- `artifacts/contracts/DIDRegistry.sol/DIDRegistry.json`
- `artifacts/contracts/CertificateRegistry.sol/CertificateRegistry.json`

### Option B: Copy ABI to backend folder (recommended)
Copy ABI JSON files into:

```
backend-go/config/abi/
```

Then backend can load ABIs consistently.

---

## 6) Update backend contract binding config (important)

If your backend uses `contracts.json`, update it after deployment with addresses.

Example file format:

```json
{
  "didRegistry": "0x...",
  "certificateRegistry": "0x..."
}
```

---

## 7) Register + activate issuer

Your smart contract enforces:
- issuer must be registered in DIDRegistry
- issuer must be active before issuing certificates

Run issuer registration script (from project root):

```powershell
cd ..
node scripts/registerIssuer.js
```

This calls:
- `registerIssuer(address, did)` OR equivalent
- `activateIssuer(address)`

---

## 8) Quick verification (contract exists on chain)

Run (optional) to check contract code is deployed:

```powershell
$addr="0xPASTE_CERTIFICATE_REGISTRY_ADDRESS"
curl -X POST http://127.0.0.1:8545 `
  -H "Content-Type: application/json" `
  -d "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getCode\",\"params\":[`"$addr`",`"latest`"],\"id\":1}"
```

Expected:
- `"result"` should NOT be `"0x"`

---

## 9) Clean build artifacts (optional)

To clean Hardhat build cache:

```powershell
Remove-Item -Recurse -Force artifacts, cache -ErrorAction SilentlyContinue
```

---

## Notes
- If you restart `npx hardhat node`, the chain resets and you MUST redeploy.
- Artifacts (`artifacts/`, `cache/`) should be in `.gitignore`.
