resource "kubernetes_service" "app" {
  metadata {
    name      = var.app
    namespace = var.namespace
  }
  spec {
    selector = {
      app = kubernetes_deployment.app.metadata.0.labels.app
    }
    port {
      port        = 8080
      target_port = 8080
    }
    type = "LoadBalancer"
  }
  depends_on = [
    kubernetes_namespace.golang-app,
  ]
}
