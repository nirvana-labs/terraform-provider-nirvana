data "nirvana_networking_connect_connections" "example_networking_connect_connections" {
  project_id = "project_id"
  bandwidth_mbps = 50
  name = "name"
  networking_connect_connection_provider = "provider"
  provider_region = "provider_region"
  region = "region"
  status = "pending"
  tags = ["string"]
}
