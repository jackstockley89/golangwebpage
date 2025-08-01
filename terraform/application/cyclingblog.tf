resource "kubernetes_deployment_v1" "cycleblog" {
  metadata {
    name      = var.app_deployment.app
    namespace = var.namespace
    labels = {
      app = var.app_deployment.app
    }
  }

  spec {
    replicas = var.app_deployment.replicas

    selector {
      match_labels = {
        app = var.app_deployment.app
      }
    }

    template {
      metadata {
        labels = {
          app = var.app_deployment.app
        }
      }

      spec {
        container {
          name              = var.app_deployment.app
          image             = var.app_deployment.template.containers.image
          image_pull_policy = var.app_deployment.template.containers.imagePullPolicy
          security_context {
            allow_privilege_escalation = false
            run_as_non_root = true
            run_as_user = 1000
            run_as_group = 1000
            seccomp_profile {
              type = "RuntimeDefault"
            }
            capabilities {
              drop = ["ALL"]
            }
          }
          resources {
            requests = {
              memory = var.app_deployment.template.containers.resources.requests.memory
              cpu    = var.app_deployment.template.containers.resources.requests.cpu
            }
            limits = {
              memory = var.app_deployment.template.containers.resources.limits.memory
              cpu    = var.app_deployment.template.containers.resources.limits.cpu
            }
          }
          port {
            container_port = var.app_service.port
          }
        }
      }
    }
  }
}