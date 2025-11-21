resource "nirvana_rpc_node_flex" "example_rpc_node_flex" {
  blockchain = "ethereum"
  name = "my-ethereum-node"
  network = "mainnet"
  project_id = "123e4567-e89b-12d3-a456-426614174000"
  tags = ["production", "ethereum"]
}
