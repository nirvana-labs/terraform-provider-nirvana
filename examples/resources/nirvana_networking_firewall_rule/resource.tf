resource "nirvana_networking_firewall_rule" "example_networking_firewall_rule" {
  vpc_id = "vpc_id"
  destination_address = "10.0.0.0/25"
  destination_ports = ["22", "80", "443"]
  name = "my-firewall-rule"
  protocol = "tcp"
  source_address = "0.0.0.0/0"
  tags = ["production", "api", "access"]
}
