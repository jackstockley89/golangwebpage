resource "kubernetes_resource_quota_v1" "resource-quota" {
  count = var.resources.namespace_resources_enabled == true ? 1 : 0
  metadata {
    name      = "${var.namespace}-resource-quota"
    namespace = var.namespace
  }
  spec {
    hard = {
      pods = 50
    }
  }
  depends_on = [kubernetes_namespace_v1.namespace]
}