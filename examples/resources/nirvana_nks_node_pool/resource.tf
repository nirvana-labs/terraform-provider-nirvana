resource "nirvana_nks_node_pool" "example_nks_node_pool" {
  cluster_id = "cluster_id"
  name = "my-node-pool"
  node_config = {
    boot_volume = {
      size = 100
      type = "abs"
    }
    instance_type = "n1-standard-8"
    labels = ["env=prod", "team=platform"]
  }
  node_count = 3
  tags = ["production", "ethereum"]
}
