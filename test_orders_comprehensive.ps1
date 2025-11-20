$headers = @{
    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI1Mzg3MjcsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.8WRbk7oR8qYiQO5B2UA0pSnHkXxPfXLiJjdNP5v3dys"
}

# Test different page sizes and pages
try {
    Write-Host "Testing page=2, page_size=5:"
    $response = Invoke-WebRequest -Uri "http://localhost:8081/api/admin/orders?page=2&page_size=5" -Headers $headers -Method GET
    Write-Host "Status Code: $($response.StatusCode)"
    Write-Host "Response Body: $($response.Content)"
    
    Write-Host "`n---`n"
    
    Write-Host "Testing default parameters (no page/page_size):"
    $response2 = Invoke-WebRequest -Uri "http://localhost:8081/api/admin/orders" -Headers $headers -Method GET
    Write-Host "Status Code: $($response2.StatusCode)"
    Write-Host "Response Body: $($response2.Content)"
    
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