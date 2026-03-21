resource "nirvana_nks_cluster" "example_nks_cluster" {
  name = "my-cluster"
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  region = "us-wdc-1"
  vpc_id = "123e4567-e89b-12d3-a456-426614174000"
  tags = ["production", "ethereum"]
}
