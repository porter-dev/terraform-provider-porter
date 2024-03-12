resource "porter_aks_cluster" "my_new_azure_cluster" {
  name            = "my-new-cluster-azure"
  region          = "eastus"
  version         = "v1.27"
  is_soc2_enabled = false
  vpc_cidr        = "10.78.0.0/16"
  service_cidr    = "172.17.0.0/16"
  node_groups = [
    {
      type          = "application"
      instance_type = "Standard_B2als_v2"
      min_nodes     = 1
      max_nodes     = 10
    },
    {
      type          = "system"
      instance_type = "Standard_B2als_v2"
      min_nodes     = 1
      max_nodes     = 5
    },
    {
      type          = "custom"
      instance_type = "Standard_NC4as_T4_v3"
      min_nodes     = 1
      max_nodes     = 5
    },
    {
      type          = "monitoring"
      instance_type = "Standard_B2as_v2"
      min_nodes     = 1
      max_nodes     = 5
    }
  ]
}
