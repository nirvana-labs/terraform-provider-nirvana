resource "nirvana_networking_connect_connection" "example_networking_connect_connection" {
  bandwidth_mbps = 50
  cidrs = ["10.0.0.0/16"]
  name = "my-connect-connection"
  provider_cidrs = ["172.16.0.0/16"]
  region = "us-wdc-1"
  aws = {
    account_id = "523816707215"
    region = "us-east-1"
  }
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  tags = ["production", "ethereum"]
}
