$headers = @{
    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI1Mzg3MjcsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.8WRbk7oR8qYiQO5B2UA0pSnHkXxPfXLiJjdNP5v3dys"
}

try {
    $response = Invoke-WebRequest -Uri "http://localhost:8081/api/admin/test_template?page=1&page_size=10" -Headers $headers -Method GET
    Write-Host "Status Code: $($response.StatusCode)"
    Write-Host "Response Body: $($response.Content)"
} catch {
    Write-Host "Error: $($_.Exception.Message)"
    if ($_.Exception.Response) {
        Write-Host "Status Code: $($_.Exception.Response.StatusCode)"
        $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
        $reader.BaseStream.Position = 0
        $reader.DiscardBufferedData()
        $responseBody = $reader.ReadToEnd()
        Write-Host "Error Response: $responseBody"
    }
}