resource "porter_app" "my_app_on_porter" {
  name  = "my-application_name"
  image = "nginx:latest"
  services = [
    {
      name          = "my-nginx-pod"
      type          = "web"
      cpu_cores     = 0.2
      ram_megabytes = 256
    },
    {
      name          = "restart-my-app"
      type          = "job"
      cpu_cores     = 0.2
      ram_megabytes = 256
      command       = "echo 'Restarting my app'"
    },

  ]
}
