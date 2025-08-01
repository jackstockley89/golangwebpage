resource "kubernetes_service_v1" "app_service" {
  metadata {
    name      = var.app_service.name
    namespace = var.namespace
    labels = {
      app = var.app_deployment.app
    }
  }

  spec {
    selector = {
      app = var.app_deployment.app
    }

    port {
      name        = var.app_service.protocol_name
      protocol    = var.app_service.protocol
      port        = var.app_service.port
      target_port = var.app_service.target_port
    }
  }
}