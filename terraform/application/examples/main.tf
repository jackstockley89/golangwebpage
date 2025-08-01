module "application_deploy" {
  source = "../"

  namespace = "cyclingblog"

  app_deployment = {
    app      = "blog"
    replicas = 2
    template = {
      containers = {
        image           = "jackstock8904/cycling-blog:latest"
        imagePullPolicy = "IfNotPresent"
        resources = {
          requests = {
            memory = "64Mi"
            cpu    = "250m"
          }
          limits = {
            memory = "128Mi"
            cpu    = "500m"
          }
        }
      }
    }
  }

  app_service = {
    name          = "blog-service"
    protocol      = "TCP"
    protocol_name = "http"
    port          = 8080
    target_port   = 8080
  }

  app_ingress = {
    name      = "blog-ingress"
    host      = "blog.example.com"
    paths     = [
      {
        path        = "/"
        path_type   = "Prefix"
        service_name = "blog_service"
        service_port = 8080
      }
    ]
    annotations = {
      "nginx.ingress.kubernetes.io/rewrite-target" = "/"
    }
  }
}




