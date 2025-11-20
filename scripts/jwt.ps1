$ErrorActionPreference = "Stop"
param(
  [string]$Secret = "devsecret",
  [string[]]$Roles = @("admin"),
  [int]$ExpiresMinutes = 30
)

function To-Base64Url($bytes){
  $b64 = [System.Convert]::ToBase64String($bytes)
  $b64 = $b64.TrimEnd('=') -replace '\+','-' -replace '/','_'
  return $b64
}

function Get-Bytes($text){ [System.Text.Encoding]::UTF8.GetBytes($text) }

$header = @{ alg = "HS256"; typ = "JWT" }
$exp = [int]([DateTimeOffset]::UtcNow.AddMinutes($ExpiresMinutes).ToUnixTimeSeconds())
$payload = @{ roles = $Roles; exp = $exp }

$headerJson = ($header | ConvertTo-Json -Compress)
$payloadJson = ($payload | ConvertTo-Json -Compress)

$headerB64 = To-Base64Url (Get-Bytes $headerJson)
$payloadB64 = To-Base64Url (Get-Bytes $payloadJson)
$unsigned = "$headerB64.$payloadB64"

$hmac = New-Object System.Security.Cryptography.HMACSHA256 ((Get-Bytes $Secret))
$sig = $hmac.ComputeHash((Get-Bytes $unsigned))
$sigB64 = To-Base64Url $sig

$token = "$unsigned.$sigB64"
Write-Output $token
