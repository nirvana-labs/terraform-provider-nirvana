resource "nirvana_volume" "example_volume" {
  size = 100
  vm_id = "vm_id"
  type = "nvme"
}
