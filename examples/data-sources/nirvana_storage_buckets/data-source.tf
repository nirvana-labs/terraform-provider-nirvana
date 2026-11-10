data "nirvana_storage_buckets" "example_storage_buckets" {
  project_id = "project_id"
  name = "name"
  region = "region"
  status = "ready"
  tags = ["string"]
}
