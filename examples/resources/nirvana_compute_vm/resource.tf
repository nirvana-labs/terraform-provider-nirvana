resource "nirvana_compute_vm" "example_compute_vm" {
  boot_volume = {
    size = 100
  }
  cpu = {
    cores = 2
  }
  name = "my-vm"
  need_public_ip = true
  os_image_name = "noble-2024-12-06"
  ports = ["22", "80", "443"]
  ram = {
    size = 2
  }
  region = "amsterdam"
  source_address = "0.0.0.0/0"
  ssh_key = {
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC1234567890"
  }
  data_volumes = [{
    size = 100
    type = "nvme"
  }]
  subnet_id = "123e4567-e89b-12d3-a456-426614174000"
}
