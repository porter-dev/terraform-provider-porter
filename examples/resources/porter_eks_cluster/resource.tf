resource "porter_eks_cluster" "my_new_eks_cluster" {
  name            = "my-new-cluster-aws"
  region          = "us-west-2"
  version         = "v1.27"
  is_soc2_enabled = false
  vpc_cidr        = "10.78.0.0/16"
  service_cidr    = "172.17.0.0/16"
  node_groups = [
    {
      type          = "application"
      instance_type = "t3.medium"
      min_nodes     = 1
      max_nodes     = 10
    },
    {
      type          = "system"
      instance_type = "t3.medium"
      min_nodes     = 1
      max_nodes     = 5
    },
    {
      type          = "custom"
      instance_type = "g4dn.2xlarge"
      min_nodes     = 1
      max_nodes     = 5
    },
    {
      type          = "monitoring"
      instance_type = "t3.large"
      min_nodes     = 1
      max_nodes     = 5
    }
  ]
}
