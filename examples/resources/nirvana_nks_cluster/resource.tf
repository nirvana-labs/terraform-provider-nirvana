resource "nirvana_nks_cluster" "example_nks_cluster" {
  autoscaling = true
  kubernetes_version = "v1.34.4"
  name = "my-cluster"
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  region = "us-sva-2"
  vpc_id = "123e4567-e89b-12d3-a456-426614174000"
  tags = ["production", "ethereum"]
}
