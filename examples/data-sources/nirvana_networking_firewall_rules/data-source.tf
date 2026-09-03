data "nirvana_networking_firewall_rules" "example_networking_firewall_rules" {
  vpc_id = "vpc_id"
  name = "name"
  protocol = "tcp"
  status = "ready"
  tags = ["string"]
}
