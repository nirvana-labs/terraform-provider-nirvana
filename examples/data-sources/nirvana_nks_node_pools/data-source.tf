data "nirvana_nks_node_pools" "example_nks_node_pools" {
  cluster_id = "cluster_id"
  instance_type = "instance_type"
  name = "name"
  node_count_max = 0
  node_count_min = 0
  status = "ready"
  tags = ["string"]
}
