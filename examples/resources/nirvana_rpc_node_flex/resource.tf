resource "nirvana_rpc_node_flex" "example_rpc_node_flex" {
  blockchain = "ethereum"
  name = "my-ethereum-node"
  network = "mainnet"
  tags = ["production", "ethereum"]
}
