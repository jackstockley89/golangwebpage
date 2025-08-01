resource "kubernetes_network_policy_v1" "network-policy-default" {
  count = var.resources.namespace_resources_enabled == "true" ? 1 : 0
  metadata {
    name      = "${var.namespace}-network-policy"
    namespace = var.namespace
  }

  spec {
    pod_selector {}
    policy_types = ["Ingress"]
    ingress {
      from {
        pod_selector {}
      }
    }
  }
  depends_on = [kubernetes_namespace_v1.namespace]
}

resource "kubernetes_network_policy_v1" "network-policy-allow-ingress-controllers" {
  metadata {
    name      = "allow-ingress-ingress-controllers"
    namespace = var.namespace
  }

  spec {
    pod_selector {}
    policy_types = ["Ingress"]
    ingress {
      from {
        pod_selector {
          match_labels = {
            conponent = "ingress-controllers"
          }
        }
      }
    }
  }
  depends_on = [kubernetes_namespace_v1.namespace]
}