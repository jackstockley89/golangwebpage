ci {
  include    = [ "namespaces/live.cloud-platform.service.justice.gov.uk/*/*promethus.yml", "namespaces/live.cloud-platform.service.justice.gov.uk/*/*promethus.yaml" ]
}

rule {
  match {
    path = "kubectl-deploy/*/.*"
    kind = "alerting"
    annotation = "(.*)" {
      value = "(.*)"
    }
    expr = "(.*)" {
      value = "(.*)"
    }
  }
  ignore {
    command = "watch"
  }
}
