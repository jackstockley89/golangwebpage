resource "kubernetes_namespace" "golang-app" {
  metadata {
    annotations = {
      name = var.namespace
    }

    labels = {
      mylabel = var.namespace
    }

    name = var.namespace
  }
}