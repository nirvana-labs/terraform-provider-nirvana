resource "nirvana_storage_bucket" "example_storage_bucket" {
  name = "my-bucket"
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  region = "us-sva-2"
  tags = ["production", "ethereum"]
}
