resource "porter_app" "my_app_on_porter" {
  name  = "my-application_name"
  image = "nginx:latest"
  services = [
    {
      name         = "my-nginx-pod"
      type         = "web"
      cpuCores     = 0.2
      ramMegabytes = 256
    },
    {
      name         = "restart-my-app"
      type         = "job"
      cpuCores     = 0.2
      ramMegabytes = 256
      command      = "echo 'Restarting my app'"
    },

  ]
}
