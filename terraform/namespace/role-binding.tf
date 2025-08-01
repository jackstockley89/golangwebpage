resource "kubernetes_role_binding_v1" "role-binding" {
  count = var.resources.role_binding_enabled ? 1 : 0
  metadata {
    name      = "${var.namespace}-admin"
    namespace = var.namespace
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "admin"
  }
  subject {
    kind      = "Group"
    name      = "github:webops"
    api_group = "rbac.authorization.k8s.io"
  }
  depends_on = [kubernetes_namespace_v1.namespace]
}