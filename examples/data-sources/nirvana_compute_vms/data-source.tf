data "nirvana_compute_vms" "example_compute_vms" {
  project_id = "project_id"
  name = "name"
  public_ip_enabled = true
  region = "region"
  status = "pending"
  subnet_id = "subnet_id"
  tags = ["string"]
  vpc_id = "vpc_id"
}
