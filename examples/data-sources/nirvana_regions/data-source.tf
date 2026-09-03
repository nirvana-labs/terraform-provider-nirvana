data "nirvana_regions" "example_regions" {
  availability = "live"
  compute_vms = true
  networking_connect = true
  networking_vpcs = true
  nks_autoscaling = true
  nks_clusters = true
  storage_abs = true
  storage_local_nvme = true
}
