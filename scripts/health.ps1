try {
  $resp = Invoke-WebRequest http://localhost:8080/v1/healthz -UseBasicParsing
  $resp.Content
} catch {
  Write-Host "health check failed"
  exit 1
}