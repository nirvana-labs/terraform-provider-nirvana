resource "nirvana_vm" "example_vm" {
  cpu = {
    cores = 2
  }
  name = "my-vm"
  need_public_ip = true
  os_image_id = 1
  ports = ["22", "80", "443"]
  ram = {
    size = 2
    unit = "GB"
  }
  region = "amsterdam"
  source_address = "0.0.0.0/0"
  ssh_key = {
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC1234567890"
  }
  storage = [{
    size = 100
    type = "nvme"
    unit = "GB"
    disk_name = "disk_name"
  }]
  subnet_id = "123e4567-e89b-12d3-a456-426614174000"
}
