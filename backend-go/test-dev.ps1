Write-Host "=== DEV END-TO-END TEST START ===" -ForegroundColor Cyan

# -------------------------
# 1. ISSUE CERTIFICATE
# -------------------------
Write-Host "`n[1] Issuing certificate..." -ForegroundColor Yellow

$issueBody = @{
    degreeHash = "123456789"
    cgpa = 85
    issuerDidHash = "987654321"
    issuerSignatureHash = "111222333"
} | ConvertTo-Json

$issueResponse = Invoke-RestMethod `
    -Uri "http://localhost:8080/api/v1/issue" `
    -Method POST `
    -ContentType "application/json" `
    -Body $issueBody

$certHash = $issueResponse.certificateHash
Write-Host "Issued certificate hash: $certHash" -ForegroundColor Green


# -------------------------
# 2. VERIFY CERTIFICATE
# -------------------------
Write-Host "`n[2] Verifying certificate (should PASS)..." -ForegroundColor Yellow

# ⚠️ Update this Poseidon hash to match your ZKP public[0]
# This must be the SAME decimal value used during proof generation
$poseidonHashDecimal = "1458734928347928374928374"

$verifyBody = @{
    proof  = (Get-Content "./zkp/proof.json" | ConvertFrom-Json)
    public = @($poseidonHashDecimal, "80")
} | ConvertTo-Json -Depth 10

$verifyResponse1 = Invoke-RestMethod `
    -Uri "http://localhost:8080/api/v1/verify" `
    -Method POST `
    -ContentType "application/json" `
    -Body $verifyBody

Write-Host "Verify response:" -ForegroundColor Green
$verifyResponse1 | ConvertTo-Json


# -------------------------
# 3. REVOKE CERTIFICATE
# -------------------------
Write-Host "`n[3] Revoking certificate..." -ForegroundColor Yellow

$revokeBody = @{
    certificateHash = $certHash
} | ConvertTo-Json

$revokeResponse = Invoke-RestMethod `
    -Uri "http://localhost:8080/api/v1/revoke" `
    -Method POST `
    -ContentType "application/json" `
    -Body $revokeBody

Write-Host "Revoke response:" -ForegroundColor Green
$revokeResponse | ConvertTo-Json


# -------------------------
# 4. VERIFY AGAIN (FAIL EXPECTED)
# -------------------------
Write-Host "`n[4] Verifying certificate again (should FAIL)..." -ForegroundColor Yellow

$verifyResponse2 = Invoke-RestMethod `
    -Uri "http://localhost:8080/api/v1/verify" `
    -Method POST `
    -ContentType "application/json" `
    -Body $verifyBody

Write-Host "Verify response after revoke:" -ForegroundColor Green
$verifyResponse2 | ConvertTo-Json


Write-Host "`n=== DEV END-TO-END TEST COMPLETE ===" -ForegroundColor Cyan
