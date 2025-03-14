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
  os_image_name = "noble-2024-12-06"
  public_ip_enabled = true
  region = "us-sea-1"
  ssh_key = {
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDJiJabIUkXw7VrQG+yBohvhEsyoKEYvejZc4RFzV5maybqQei1punVsoe4r6gJttMM1Gr3cNr3OfepikCQAhAchw5ww94ZWqDsDYIqMrlDFbqhGTXDNzFAjeVIKptCOlz9k+7aM69YtLXJ6gFUCq1fbK9PjY+AK28UpMfKYUcyHQ== noname"
  }
  subnet_id = "123e4567-e89b-12d3-a456-426614174000"
  data_volumes = [{
    name = "my-data-volume"
    size = 100
  }]
}
