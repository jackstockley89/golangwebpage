ci {
  include    = [ "kubectl-deploy/*/*promethus.yml", "kubectl-deploy/*/*promethus.yaml" ]
}

rule {
  
  match {
    kind = "alerting"
  }
}
