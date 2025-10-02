resource "nirvana_api_key" "example_api_key" {
  expires_at = "2025-12-31T23:59:59Z"
  name = "my-api-key"
  starts_at = "2025-01-01T00:00:00Z"
  tags = ["production", "ethereum"]
}
