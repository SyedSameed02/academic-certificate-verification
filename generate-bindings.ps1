$ErrorActionPreference = "Stop"

$ARTIFACTS = "smart-contracts/artifacts/contracts"
$OUT_DIR = "backend-go/contracts"
$TMP_DIR = "tmp-abi"

# Create temp and output dirs
New-Item -ItemType Directory -Force -Path $OUT_DIR | Out-Null
New-Item -ItemType Directory -Force -Path $TMP_DIR | Out-Null

Write-Host "Extracting ABI files..."

# Extract ABI using node (no jq needed)
node -e "const fs=require('fs'); fs.writeFileSync('$TMP_DIR/DIDRegistry.abi.json', JSON.stringify(require('./$ARTIFACTS/DIDRegistry.sol/DIDRegistry.json').abi))"
node -e "const fs=require('fs'); fs.writeFileSync('$TMP_DIR/CertificateRegistry.abi.json', JSON.stringify(require('./$ARTIFACTS/CertificateRegistry.sol/CertificateRegistry.json').abi))"

Write-Host "Generating Go bindings..."

# DIDRegistry
abigen `
  --abi "$TMP_DIR/DIDRegistry.abi.json" `
  --pkg contracts `
  --type DIDRegistry `
  --out "$OUT_DIR/did_registry.go"

# CertificateRegistry
abigen `
  --abi "$TMP_DIR/CertificateRegistry.abi.json" `
  --pkg contracts `
  --type CertificateRegistry `
  --out "$OUT_DIR/certificate_registry.go"

Write-Host "Go bindings generated successfully in backend-go/contracts"
