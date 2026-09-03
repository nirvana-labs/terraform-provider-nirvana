data "nirvana_nks_clusters" "example_nks_clusters" {
  project_id = "project_id"
  autoscaling = true
  kubernetes_version = "kubernetes_version"
  name = "name"
  region = "region"
  status = "ready"
  tags = ["string"]
  vpc_id = "vpc_id"
}
