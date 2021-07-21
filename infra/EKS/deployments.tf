resource "kubernetes_deployment" "app" {
  metadata {
    name      = var.app
    namespace = var.namespace
    labels = {
      app = var.app
    }
  }
  spec {
    replicas = 2

    selector {
      match_labels = {
        app = var.app
      }
    }
    template {
      metadata {
        labels = {
          app = var.app
        }
      }
      spec {
        container {
          image = var.docker-image
          name  = var.app
          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "64Mi"
            }
          }
          port {
            name           = "port-8080"
            container_port = 8080
            protocol       = "TCP"
          }
          env {
            name  = "PGPASSFILE"
            value = "/home/appuser/.pgpass"
          }
        }
      }
    }
  }
  depends_on = [
    kubernetes_namespace.golang-app,
  ]
}