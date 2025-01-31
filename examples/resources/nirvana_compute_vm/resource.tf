resource "nirvana_compute_vm" "example_compute_vm" {
  boot_volume = {
    size = 100
  }
  cpu = {
    cores = 2
  }
  name = "my-vm"
  os_image_name = "noble-2024-12-06"
  public_ip_enabled = true
  ram = {
    size = 2
  }
  region = "us-sea-1"
  ssh_key = {
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC1234567890"
  }
  data_volumes = [{
    size = 100
    type = "nvme"
  }]
  ports = ["22", "80", "443"]
  source_address = "0.0.0.0/0"
  subnet_id = "123e4567-e89b-12d3-a456-426614174000"
}
