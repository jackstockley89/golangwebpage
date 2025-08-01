variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "eu-west-2"
}

variable "kubeconfig_path" {
  description = "Path to kubeconfig file"
  type        = string
  default     = "~/.kube/kind/config"
}

variable "namespace" {
  description = "value of namespace"
  type        = string
  default     = "module-test-namespace"
}

variable "namespace_annotations" {
  description = "Labels to apply to the namespace"
  type = object({
    business_unit = string
    slack_channel = string
    application   = string
    owner         = string
    source_code   = string
    team_name     = string
  })
}

variable "namespace_labels" {
  description = "Labels to apply to the namespace"
  type = object({
    is_production          = string
    environment_name       = string
    pod_security           = string
    infrastructure_support = string
  })
}


variable "resources" {
  description = "enable choosen resources"
  type = object({
    route53_enabled             = bool
    namespace_resources_enabled = bool
    role_binding_enabled        = bool
  })
  default = {
    route53_enabled             = false
    namespace_resources_enabled = true
    role_binding_enabled        = false
  }
}
