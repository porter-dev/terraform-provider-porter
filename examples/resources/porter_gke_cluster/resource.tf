resource "porter_gke_cluster" "my_new_gke_cluster" {
  name            = "my-new-cluster-gcp"
  region          = "us-east4"
  version         = "v1.27"
  is_soc2_enabled = false
  vpc_cidr        = "10.78.0.0/16"
  service_cidr    = "172.17.0.0/16"
  node_groups = [
    {
      type          = "application"
      instance_type = "e2-standard-4"
      min_nodes     = 1
      max_nodes     = 10
    },
    {
      type          = "system"
      instance_type = "e2-standard-4"
      min_nodes     = 1
      max_nodes     = 5
    },
    {
      type          = "custom"
      instance_type = "n1-standard-32"
      min_nodes     = 1
      max_nodes     = 5
    },
    {
      type          = "monitoring"
      instance_type = "e2-standard-8"
      min_nodes     = 1
      max_nodes     = 5
    }
  ]
}
