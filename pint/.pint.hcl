ci {
  include    = [ "kubectl-deploy/*/*promethus.yml", "kubectl-deploy/*/*promethus.yaml" ]
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
