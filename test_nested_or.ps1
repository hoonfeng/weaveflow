$headers = @{
    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI1Mzg3MjcsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.8WRbk7oR8qYiQO5B2UA0pSnHkXxPfXLiJjdNP5v3dys"
}

try {
    # Test the nested or expression directly
    $response = Invoke-WebRequest -Uri "http://localhost:8081/api/admin/test_nested_or" -Headers $headers -Method GET
    Write-Host "Nested OR Test (no params): $($response.Content)"
    
    Write-Host "`n---`n"
    
    # Test with size parameter
    $response2 = Invoke-WebRequest -Uri "http://localhost:8081/api/admin/test_nested_or?size=20" -Headers $headers -Method GET
    Write-Host "Nested OR Test (size=20): $($response2.Content)"
    
    Write-Host "`n---`n"
    
    # Test with page_size parameter
    $response3 = Invoke-WebRequest -Uri "http://localhost:8081/api/admin/test_nested_or?page_size=30" -Headers $headers -Method GET
    Write-Host "Nested OR Test (page_size=30): $($response3.Content)"
    
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