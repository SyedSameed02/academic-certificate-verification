#Requires -Version 5.1
<#
bootstrap.ps1 — One-command dev bootstrap for the project.

What it does:
1) Starts Hardhat node (new PowerShell window)
2) Deploys contracts to localhost
3) Updates backend-go/.env with deployed addresses
4) Runs issuer registration script
5) Starts Go backend server (new PowerShell window)
6) Runs backend-go/test-dev.ps1 end-to-end tests

Run from project root:
  Set-ExecutionPolicy -Scope Process Bypass
  .\scripts\bootstrap.ps1
#>

$ErrorActionPreference = "Stop"

function Write-Step($msg) {
  Write-Host "`n==> $msg" -ForegroundColor Cyan
}

function Fail($msg) {
  Write-Host "`n[ERROR] $msg" -ForegroundColor Red
  exit 1
}

# -------------------------
# Resolve paths
# -------------------------
$root = (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
$smartContractsDir = Join-Path $root "smart-contracts"
$backendDir = Join-Path $root "backend-go"
$envPath = Join-Path $backendDir ".env"
$registerScript = Join-Path $root "scripts\registerIssuer.js"
$deployScript = Join-Path $smartContractsDir "scripts\deploy.js"
$testScript = Join-Path $backendDir "test-dev.ps1"

if (!(Test-Path $smartContractsDir)) { Fail "smart-contracts folder not found: $smartContractsDir" }
if (!(Test-Path $backendDir)) { Fail "backend-go folder not found: $backendDir" }
if (!(Test-Path $deployScript)) { Fail "deploy.js not found: $deployScript" }
if (!(Test-Path $registerScript)) { Fail "registerIssuer.js not found: $registerScript" }
if (!(Test-Path $testScript)) { Fail "test-dev.ps1 not found: $testScript" }

Write-Host "Project root: $root" -ForegroundColor DarkGray

# -------------------------
# 1) Start Hardhat node in new window
# -------------------------
Write-Step "Starting Hardhat node (new window)..."
Start-Process powershell -ArgumentList @(
  "-NoExit",
  "-Command",
  "cd `"$smartContractsDir`"; npx hardhat node"
) | Out-Null

Write-Host "Waiting 4 seconds for Hardhat node..." -ForegroundColor DarkGray
Start-Sleep -Seconds 4

# -------------------------
# 2) Deploy contracts
# -------------------------
Write-Step "Deploying contracts (localhost)..."
Push-Location $smartContractsDir

# Ensure deps installed
if (!(Test-Path (Join-Path $smartContractsDir "node_modules"))) {
  Write-Host "node_modules not found in smart-contracts. Installing..." -ForegroundColor Yellow
  npm install
}

$deployOutput = & npx hardhat run scripts/deploy.js --network localhost 2>&1 | Out-String
Write-Host $deployOutput

Pop-Location

# Extract addresses from deploy output (supports common patterns)
# We look for 0x...40hex
$addresses = [regex]::Matches($deployOutput, "0x[a-fA-F0-9]{40}") | ForEach-Object { $_.Value }
$addresses = $addresses | Select-Object -Unique

if ($addresses.Count -lt 2) {
  Fail "Could not extract both contract addresses from deploy output. Make sure deploy.js prints addresses."
}

# Heuristic: first is DID, second is CERT (matches common deploy script ordering)
$didAddr  = $addresses[0]
$certAddr = $addresses[1]

Write-Host "Detected DID_REGISTRY_ADDRESS: $didAddr" -ForegroundColor Green
Write-Host "Detected CERTIFICATE_REGISTRY_ADDRESS: $certAddr" -ForegroundColor Green

# -------------------------
# 3) Update backend-go/.env addresses
# -------------------------
Write-Step "Updating backend-go/.env with contract addresses..."

if (!(Test-Path $envPath)) {
  Fail ".env not found at: $envPath (create it first)"
}

$envText = Get-Content $envPath -Raw

# Replace or append DID_REGISTRY_ADDRESS
if ($envText -match "(?m)^DID_REGISTRY_ADDRESS=") {
  $envText = [regex]::Replace($envText, "(?m)^DID_REGISTRY_ADDRESS=.*$", "DID_REGISTRY_ADDRESS=$didAddr")
} else {
  $envText += "`nDID_REGISTRY_ADDRESS=$didAddr"
}

# Replace or append CERTIFICATE_REGISTRY_ADDRESS
if ($envText -match "(?m)^CERTIFICATE_REGISTRY_ADDRESS=") {
  $envText = [regex]::Replace($envText, "(?m)^CERTIFICATE_REGISTRY_ADDRESS=.*$", "CERTIFICATE_REGISTRY_ADDRESS=$certAddr")
} else {
  $envText += "`nCERTIFICATE_REGISTRY_ADDRESS=$certAddr"
}

Set-Content -Path $envPath -Value $envText -Encoding UTF8
Write-Host ".env updated." -ForegroundColor Green

# -------------------------
# 4) Register + activate issuer
# -------------------------
Write-Step "Registering & activating issuer..."
Push-Location $root

# Ensure deps installed at root if script uses ethers/dotenv from root node_modules
if (!(Test-Path (Join-Path $root "node_modules"))) {
  Write-Host "node_modules not found at root. Installing dependencies for issuer script..." -ForegroundColor Yellow
  npm install
}

$regOutput = & node $registerScript 2>&1 | Out-String
Write-Host $regOutput

Pop-Location

# -------------------------
# 5) Start Go backend (new window)
# -------------------------
Write-Step "Starting Go backend server (new window)..."
Start-Process powershell -ArgumentList @(
  "-NoExit",
  "-Command",
  "cd `"$backendDir`"; go run ./main.go"
) | Out-Null

Write-Host "Waiting 3 seconds for Go server..." -ForegroundColor DarkGray
Start-Sleep -Seconds 3

# -------------------------
# 6) Run end-to-end dev test
# -------------------------
Write-Step "Running end-to-end test (test-dev.ps1)..."
Push-Location $backendDir
& powershell -ExecutionPolicy Bypass -File $testScript
Pop-Location

Write-Host "`n✅ Bootstrap complete." -ForegroundColor Green
