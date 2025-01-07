resource "nirvana_firewall_rule" "example_firewall_rule" {
  vpc_id = "vpc_id"
  destination = {
    address = "0.0.0.0/0"
    ports = ["22", "80", "443"]
  }
  name = "my-firewall-rule"
  protocol = "tcp"
  source = {
    address = "0.0.0.0/0"
    ports = ["22", "80", "443"]
  }
}
