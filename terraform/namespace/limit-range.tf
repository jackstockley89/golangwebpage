resource "kubernetes_limit_range_v1" "example" {
  count = var.resources.namespace_resources_enabled == true ? 1 : 0
  metadata {
    name = "${var.namespace}-limit-range"
  }
  spec {
    limit {
      default = {
        cpu    = "1000m"
        memory = "1000Mi"
      }
      default_request = {
        cpu    = "10m"
        memory = "100Mi"
      }
      type = "Container"
    }
  }
  depends_on = [kubernetes_namespace_v1.namespace]
}