resource "nirvana_compute_volume" "example_compute_volume" {
  name = "my-data-volume"
  size = 100
  vm_id = "vm_id"
  tags = ["production", "ethereum"]
  type = "nvme"
}
