resource "nirvana_compute_vm" "example_compute_vm" {
  boot_volume = {
    size = 100
  }
  cpu_config = {
    vcpu = 2
  }
  memory_config = {
    size = 2
  }
  name = "my-vm"
  os_image_name = "ubuntu-noble-2025-04-03"
  public_ip_enabled = true
  region = "us-wdc-1"
  ssh_key = {
    public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2"
  }
  subnet_id = "123e4567-e89b-12d3-a456-426614174000"
}
