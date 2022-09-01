ci {
  include    = [ "kubectl_deploy/*/*promethus.yml", "kubectl_deploy/*/*promethus.yaml" ]
}

rule {
  
  match {
    path = "kubectl_deploy/*/*promethus.yaml"
    kind = "alerting"
  }
}

rule {
  
  match {
    path = "kubectl_deploy/*/*promethus.yml"
    kind = "alerting"
  }
}
