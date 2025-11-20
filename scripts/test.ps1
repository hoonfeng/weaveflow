$ErrorActionPreference = "Stop"
$base = "http://localhost:8080"

function Get-Json($url) {
  $resp = Invoke-WebRequest -Uri $url -UseBasicParsing
  return $resp.Content
}

function Post-Json($url, $json) {
  $resp = Invoke-WebRequest -Uri $url -Method POST -ContentType 'application/json' -Body $json -UseBasicParsing
  return $resp.Content
}

function Post-Multipart($url, $filePath) {
  $file = Get-Item $filePath
  $form = @{ files = $file }
  $resp = Invoke-WebRequest -Uri $url -Method Post -Form $form -UseBasicParsing
  return $resp.Content
}

"[GET] /docs"
Get-Json "$base/docs" | Write-Output

"[GET] /docs/openapi.json"
Get-Json "$base/docs/openapi.json" | Write-Output

"[GET] /debug/plugins"
Get-Json "$base/debug/plugins" | Write-Output

"[POST] /api/text/reverse"
Post-Json "$base/api/text/reverse" '{"text":"Hello World"}' | Write-Output

"[POST] /api/files/upload"
$tmp = [System.IO.Path]::GetTempFileName()
[System.IO.File]::WriteAllText($tmp, "hello")
Post-Multipart "$base/api/files/upload" $tmp | Write-Output
Remove-Item $tmp -Force

"[POST] /api/vector/demo"
Post-Json "$base/api/vector/demo" '{"id":"u1","vec":[0.1,0.2,0.3]}' | Write-Output
