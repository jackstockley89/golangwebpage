module "namespace_create" {
  source = "../"

  resources = {
    namespace_resources_enabled = true
    route53_enabled             = false
    role_binding_enabled        = false
  }

  namespace = "cyclingblog"

  namespace_annotations = {
    business_unit = "test"
    slack_channel = "n/a"
    application   = "cyclingblog"
    owner         = "jack_stockley"
    source_code   = "golangwebpage"
    team_name     = "n/a"
  }

  namespace_labels = {
    is_production          = "false"
    environment_name       = "test"
    pod_security           = "restricted"
    infrastructure_support = "n/a"
  }
}