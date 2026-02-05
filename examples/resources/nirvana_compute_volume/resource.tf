resource "nirvana_compute_volume" "example_compute_volume" {
  name = "my-data-volume"
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  region = "us-wdc-1"
  size = 100
  type = "nvme"
  tags = ["production", "ethereum"]
  vm_id = "123e4567-e89b-12d3-a456-426614174000"
}
