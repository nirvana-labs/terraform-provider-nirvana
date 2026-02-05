resource "nirvana_networking_vpc" "example_networking_vpc" {
  name = "my-vpc"
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  region = "us-wdc-1"
  subnet_name = "my-subnet"
  tags = ["production", "ethereum"]
}
