resource "porter_aks_cluster" "my_new_azure_cluster" {
  name         = "my-new-cluster-azure"
  region       = "eastus"
  version      = "v1.27"
  vpc_cidr     = "10.78.0.0/16"
  service_cidr = "172.17.0.0/16"
  node_groups = [
    {
      type          = "application"
      instance_type = "Standard_B2als_v2"
      min_nodes     = 1
      max_nodes     = 10
    },
    {
      type          = "gpu"
      instance_type = "Standard_NC4as_T4_v3"
      min_nodes     = 1
      max_nodes     = 5
    }
  ]
}
