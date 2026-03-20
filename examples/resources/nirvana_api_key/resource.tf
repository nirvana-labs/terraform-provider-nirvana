resource "nirvana_api_key" "example_api_key" {
  expires_at = "2025-12-31T23:59:59Z"
  name = "My API Key"
  permissions = [{
    permission = "edit"
    resource_type = "vm"
  }]
  project_ids = ["123e4567-e89b-12d3-a456-426614174000", "123e4567-e89b-12d3-a456-426614174001"]
  source_ip_rule = {
    allowed = ["192.168.1.0/24", "10.0.0.0/8"]
    blocked = ["192.168.1.100/32"]
  }
  starts_at = "2025-01-01T00:00:00Z"
  tags = ["production", "ethereum"]
}
