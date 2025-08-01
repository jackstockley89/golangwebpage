variable "aws_region" {
  description = "The AWS region to deploy the resources"
  type        = string
  default     = "us-west-2"
}

variable "kubeconfig_path" {
  description = "The path to the kubeconfig file"
  type        = string
  default     = "~/.kube/kind/config"
}

variable "s3" {
  description = "The S3 bucket to store the Terraform state"
  type = object({
    bucket   = string
    key      = string
    owner_id = string
  })
  default = {
    bucket = "value-of-bucket"
    key    = "terraform/state"
  }
}

variable "namespace" {
  description = "value of namespace"
  type        = string
  default     = "module-test-namespace"
}

variable "app_deployment" {
  description = "The deployment configuration for the application"
  type = object({
    app      = string
    replicas = number
    template = object({
      containers = object({
        image = string
        imagePullPolicy = string
        resources = object({
          requests = object({
            memory = string
            cpu    = string
          })
          limits = object({
            memory = string
            cpu    = string
          })
        })
      })
    })
  })
}

variable "app_service" {
  description = "The service configuration for the application"
  type = object({
    name          = string
    protocol_name = string
    protocol      = string
    port          = number
    target_port   = number
  })
}

variable "app_ingress" {
  description = "The ingress configuration for the application"
  type = object({
    name       = string
    host       = string
    paths      = list(object({
      path        = string
      path_type   = string
      service_name = string
      service_port = number
    }))
    annotations = map(string)
  })
}