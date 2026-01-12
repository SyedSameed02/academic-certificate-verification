

## Prerequisites
- Node.js + npm installed
- Go installed
- Docker Desktop running (only needed if compiling Circom with Docker)
- Project structure:
  - `smart-contracts/` contains `hardhat.config.js`
  - `backend-go/` contains `main.go` and `test-dev.ps1`
  - `scripts/registerIssuer.js` 

### 1) Start local Hardhat blockchain
📍 Run inside **smart-contracts/** (terminal 1):

```powershell
npx hardhat node
```

Keep this terminal running.

---

### 2) Deploy contracts
Open **new terminal** inside **smart-contracts/** (terminal 2):

```powershell
npx hardhat run scripts/deploy.js --network localhost
```

This deploys `DIDRegistry` + `CertificateRegistry` and prints their addresses.

---

### 3) Register & activate issuer
Run from **project root** (terminal 3):

```powershell
node scripts/registerIssuer.js
```

This registers the backend signer as an authorized issuer in `DIDRegistry`.

---

### 4) Start Go backend server
Run inside **backend-go/** (terminal 4):

```powershell
go run ./main.go
```

Backend runs at:
- `http://localhost:8080`

---

### 5) Run end-to-end dev test
Open **new terminal** inside **backend-go/** (terminal 5):

```powershell
./test-dev.ps1
```

✅ Expected output:
- Issue → prints certificate hash
- Verify PASS → `{ "verified": true }`
- Revoke → `{ "status": "revoked" }`
- Verify again FAIL → `{ "verified": false, "reason": "certificate_revoked" }`
