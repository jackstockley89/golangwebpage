resource "aws_route53_zone" "route53_zone" {
  count = var.resources.route53_enabled == true ? 1 : 0
  name  = "${var.namespace}.service.justice.gov.uk"

  tags = {
    team_name              = var.namespace_annotations.team_name
    business-unit          = var.namespace_annotations.business_unit
    application            = var.namespace_annotations.application
    is-production          = var.namespace_labels.is_production
    environment-name       = var.namespace_labels.environment_name
    owner                  = var.namespace_annotations.owner
    infrastructure-support = var.namespace_labels.infrastructure_support
    namespace              = var.namespace
  }
  depends_on = [kubernetes_namespace_v1.namespace]
}

resource "kubernetes_secret" "route53_zone_sec" {
  count = var.resources.route53_enabled == true ? 1 : 0
  metadata {
    name      = "${var.namespace}-route53-zone-output"
    namespace = var.namespace
  }

  data = {
    zone_id = aws_route53_zone.route53_zone[0].zone_id
  }

  depends_on = [aws_route53_zone.route53_zone]
}
