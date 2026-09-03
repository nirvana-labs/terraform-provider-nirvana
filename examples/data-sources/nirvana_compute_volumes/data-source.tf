data "nirvana_compute_volumes" "example_compute_volumes" {
  project_id = "project_id"
  attached = true
  kind = "boot"
  name = "name"
  region = "region"
  status = "ready"
  tags = ["string"]
  type = "nvme"
  vm_id = "vm_id"
}
