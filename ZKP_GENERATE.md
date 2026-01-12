# ZKP_GENERATE.md — Generate Circom + Groth16 Proof Artifacts (Windows / PowerShell)

This document explains how to regenerate all ZKP artifacts for the project:
- circuit compilation (`.r1cs`, `.wasm`, `.sym`)
- Groth16 setup (`.ptau`, `.zkey`)
- witness generation (`.wtns`)
- proof generation (`proof.json`, `public.json`)
- verification key export (`verification_key.json`)

> ✅ Run commands from **backend-go/** directory.

---

## 0) Go to backend directory

```powershell
cd C:\Users\shifa\OneDrive\Desktop\sameed\academic-certificate-verification\backend-go
```

---

## 1) Ensure circuit has public outputs

Edit:

```
zkp/circuits/certificate.circom
```

Make sure `main` exports public signals:

```circom
component main {public [onChainHash, MIN_CGPA]} = CertificateProof();
```

---

## 2) Clean old artifacts (recommended)

```powershell
Remove-Item -Recurse -Force zkp\build -ErrorAction SilentlyContinue
Remove-Item -Recurse -Force zkp\keys  -ErrorAction SilentlyContinue
Remove-Item -Force zkp\proof.json, zkp\public.json, zkp\verification_key.json -ErrorAction SilentlyContinue

mkdir zkp\build
mkdir zkp\keys
```

---

## 3) Compile circuit (Circom → build artifacts)

> Uses Docker image (recommended).

```powershell
docker run --rm `
  -v "${PWD}:/work" `
  -w /work `
  saleel/circom:2.1.6 `
  circom zkp/circuits/certificate.circom `
    -l node_modules `
    --r1cs --wasm --sym -o zkp/build
```

Output:
- `zkp/build/certificate.r1cs`
- `zkp/build/certificate_js/certificate.wasm`
- `zkp/build/certificate.sym`

---

## 4) Powers of Tau (Phase 1 setup)

> This can be reused across circuits, but safe to regenerate for demo.

```powershell
snarkjs powersoftau new bn128 12 zkp/keys/pot12_0000.ptau -v
snarkjs powersoftau contribute zkp/keys/pot12_0000.ptau zkp/keys/pot12_0001.ptau --name="First contribution" -v
snarkjs powersoftau prepare phase2 zkp/keys/pot12_0001.ptau zkp/keys/pot12_final.ptau -v
```

---

## 5) Groth16 setup (Phase 2) → generate `.zkey`

```powershell
snarkjs groth16 setup zkp/build/certificate.r1cs zkp/keys/pot12_final.ptau zkp/keys/certificate_0000.zkey
snarkjs zkey contribute zkp/keys/certificate_0000.zkey zkp/keys/certificate_final.zkey --name="Key contributor" -v
```

Output:
- `zkp/keys/certificate_final.zkey` ✅

---

## 6) Witness generation

### Fix CommonJS vs ESM issue (important)
Your project has `"type": "module"` in `package.json`, but Circom-generated witness code is CommonJS.

✅ Create a local `package.json` inside:

```
zkp/build/certificate_js/package.json
```

Content:

```json
{
  "type": "commonjs"
}
```

Then generate witness:

```powershell
node zkp/build/certificate_js/generate_witness.js `
  zkp/build/certificate_js/certificate.wasm `
  zkp/inputs/input.json `
  zkp/build/witness.wtns
```

Output:
- `zkp/build/witness.wtns` ✅

---

## 7) Generate proof + public signals

```powershell
snarkjs groth16 prove `
  zkp/keys/certificate_final.zkey `
  zkp/build/witness.wtns `
  zkp/proof.json `
  zkp/public.json
```

Outputs:
- `zkp/proof.json`
- `zkp/public.json` (should contain 2 values)

---

## 8) Export verification key

```powershell
snarkjs zkey export verificationkey `
  zkp/keys/certificate_final.zkey `
  zkp/verification_key.json
```

---

## 9) Confirm public signals count

```powershell
Get-Content zkp/verification_key.json | Select-String "nPublic"
Get-Content zkp/public.json
```

Expected:
- `nPublic: 2`
- `public.json` has 2 entries: `[onChainHash, MIN_CGPA]`

---

## 10) Verify proof manually (optional)

```powershell
snarkjs groth16 verify `
  zkp/verification_key.json `
  zkp/public.json `
  zkp/proof.json
```

Expected output:
- `OK` ✅
