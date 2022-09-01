ci {
  include    = [ "kubectl_deploy/monitoring/*promethus.yml", "kubectl_deploy/monitoring/*promethus.yaml" ]
}

rule {
  
  match {
    path = "kubectl_deploy/monitoring/*promethus.yaml"
    kind = "alerting"
  }

  match {
    path = "kubectl_deploy/monitoring/*promethus.yml"
    kind = "alerting"
  }
}
