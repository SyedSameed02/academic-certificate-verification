Write-Host "=== DEV END-TO-END TEST START ===" -ForegroundColor Cyan

# -------------------------
# 1. ISSUE CERTIFICATE
# -------------------------
Write-Host "`n[1] Issuing certificate..." -ForegroundColor Yellow

# IssueHandler expects decimal strings for hashes.
$issueBody = @{
    degreeHash = "123456789"
    cgpa = 85
    issuerDidHash = "987654321"
    issuerSignatureHash = "111222333"
} | ConvertTo-Json

try {
  $issueResponse = Invoke-RestMethod `
      -Uri "http://localhost:8080/api/v1/issue" `
      -Method POST `
      -ContentType "application/json" `
      -Body $issueBody
} catch {
  Write-Host "Issue API failed:" -ForegroundColor Red
  if ($_.Exception.Response -ne $null) {
    Write-Host "Status:" $_.Exception.Response.StatusCode.value__ -ForegroundColor Red
    $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
    $respBody = $reader.ReadToEnd()
    Write-Host "Response body:" -ForegroundColor Yellow
    Write-Host $respBody
  } else {
    Write-Host $_.Exception.Message -ForegroundColor Red
  }
  exit 1
}

$certHash = $issueResponse.certificateHash
Write-Host "Issued certificate hash: $certHash" -ForegroundColor Green


# -------------------------
# 2. VERIFY CERTIFICATE
# -------------------------
Write-Host "`n[2] Verifying certificate (should PASS)..." -ForegroundColor Yellow

# IMPORTANT:
# PowerShell's ConvertTo-Json can reformat large integers and break JSON decoding in Go.
# So we embed proof.json as RAW text (exact snarkjs output), and only stringify public signals.
$proofText = Get-Content "./zkp/proof.json" -Raw

$publicRaw = Get-Content "./zkp/public.json" | ConvertFrom-Json
$publicStr = @($publicRaw | ForEach-Object { "$_" })  # force []string

# Build verify JSON manually (proof inserted as raw JSON object)
$verifyBody = @"
{
  "proof": $proofText,
  "public": $(ConvertTo-Json $publicStr -Compress)
}
"@

try {
  $verifyResponse1 = Invoke-RestMethod `
      -Uri "http://localhost:8080/api/v1/verify" `
      -Method POST `
      -ContentType "application/json" `
      -Body $verifyBody
} catch {
  Write-Host "Verify API failed:" -ForegroundColor Red
  if ($_.Exception.Response -ne $null) {
    Write-Host "Status:" $_.Exception.Response.StatusCode.value__ -ForegroundColor Red
    $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
    $respBody = $reader.ReadToEnd()
    Write-Host "Response body:" -ForegroundColor Yellow
    Write-Host $respBody
  } else {
    Write-Host $_.Exception.Message -ForegroundColor Red
  }
  exit 1
}

Write-Host "Verify response:" -ForegroundColor Green
$verifyResponse1 | ConvertTo-Json


# -------------------------
# 3. REVOKE CERTIFICATE
# -------------------------
Write-Host "`n[3] Revoking certificate..." -ForegroundColor Yellow

$revokeBody = @{ certificateHash = $certHash } | ConvertTo-Json

try {
  $revokeResponse = Invoke-RestMethod `
      -Uri "http://localhost:8080/api/v1/revoke" `
      -Method POST `
      -ContentType "application/json" `
      -Body $revokeBody
} catch {
  Write-Host "Revoke API failed:" -ForegroundColor Red
  if ($_.Exception.Response -ne $null) {
    Write-Host "Status:" $_.Exception.Response.StatusCode.value__ -ForegroundColor Red
    $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
    $respBody = $reader.ReadToEnd()
    Write-Host "Response body:" -ForegroundColor Yellow
    Write-Host $respBody
  } else {
    Write-Host $_.Exception.Message -ForegroundColor Red
  }
  exit 1
}

Write-Host "Revoke response:" -ForegroundColor Green
$revokeResponse | ConvertTo-Json


# -------------------------
# 4. VERIFY AGAIN (FAIL EXPECTED)
# -------------------------
Write-Host "`n[4] Verifying certificate again (should FAIL)..." -ForegroundColor Yellow

try {
  $verifyResponse2 = Invoke-RestMethod `
      -Uri "http://localhost:8080/api/v1/verify" `
      -Method POST `
      -ContentType "application/json" `
      -Body $verifyBody
} catch {
  Write-Host "Verify API (after revoke) failed:" -ForegroundColor Red
  if ($_.Exception.Response -ne $null) {
    Write-Host "Status:" $_.Exception.Response.StatusCode.value__ -ForegroundColor Red
    $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
    $respBody = $reader.ReadToEnd()
    Write-Host "Response body:" -ForegroundColor Yellow
    Write-Host $respBody
  } else {
    Write-Host $_.Exception.Message -ForegroundColor Red
  }
  exit 1
}

Write-Host "Verify response after revoke:" -ForegroundColor Green
$verifyResponse2 | ConvertTo-Json

Write-Host "`n=== DEV END-TO-END TEST COMPLETE ===" -ForegroundColor Cyan
