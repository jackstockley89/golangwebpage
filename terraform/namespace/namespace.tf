resource "kubernetes_namespace_v1" "namespace" {
  metadata {

    annotations = {
      "cloud-platform.justice.gov.uk/business-unit" = var.namespace_annotations.business_unit
      "cloud-platform.justice.gov.uk/slack-channel" = var.namespace_annotations.slack_channel
      "cloud-platform.justice.gov.uk/application"   = var.namespace_annotations.application
      "cloud-platform.justice.gov.uk/owner"         = var.namespace_annotations.owner
      "cloud-platform.justice.gov.uk/source-code"   = var.namespace_annotations.source_code
      "cloud-platform.justice.gov.uk/team-name"     = var.namespace_annotations.team_name
    }

    labels = {
      "cloud-platform.justice.gov.uk/is-production"    = var.namespace_labels.is_production
      "cloud-platform.justice.gov.uk/environment-name" = var.namespace_labels.environment_name
      "pod-security.kubernetes.io/enforce"             = var.namespace_labels.pod_security
    }

    name = var.namespace
  }
}

