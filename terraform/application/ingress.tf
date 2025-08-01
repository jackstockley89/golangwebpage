resource "kubernetes_ingress_v1" "app_ingress" {
  metadata {
    name      = var.app_ingress.name
    namespace = var.namespace
    labels = {
      app = var.app_deployment.app
    }
    annotations = var.app_ingress.annotations
  }

  spec {
    rule {
      host = var.app_ingress.host
      http {
        dynamic "path" {
          for_each = var.app_ingress.paths
          content {
            path     = path.value.path
            path_type = path.value.path_type

            backend {
              service {
                name = path.value.service_name
                port {
                  number = path.value.service_port
                }
              }
            }
          }
        }
      }
    }
  }
}