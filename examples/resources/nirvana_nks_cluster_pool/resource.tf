resource "nirvana_nks_cluster_pool" "example_nks_cluster_pool" {
  cluster_id = "cluster_id"
  name = "my-node-pool"
  node_config = {
    boot_volume = {
      size = 100
      type = "abs"
    }
    cpu_config = {
      vcpu = 4
    }
    memory_config = {
      size = 8
    }
  }
  node_count = 3
  tags = ["production", "ethereum"]
}
