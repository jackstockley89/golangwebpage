# cloud-platform-terraform-template

[![Releases](https://img.shields.io/github/v/release/ministryofjustice/cloud-platform-terraform-test-applications.svg)](https://github.com/ministryofjustice/cloud-platform-terraform-test-applictions/releases)

This Terraform module will create an SQS queue and a test application that sends messages to the queue.

## Usage

```hcl
module "namespace_create" {
  source = "github.com/ministryofjustice/cloud-platform-terraform-test-applications/namespace?ref=${var.namespace_module_release}"

  namespace_resources_enabled = "true"
  route53_enabled             = "true"
  rolebinding_enabled         = "true"

  namespace = var.namespace

  business_unit = var.business_unit
  slack_channel = var.slack_channel
  application   = var.application
  owner         = var.owner
  source_code   = var.source_code
  team_name     = var.team_name
}
```

See the [examples/](examples/) folder for more information.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.2.5 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 4.27.0 |
| kubernetes | ~> 2.18.0|
| github | ~> 5.17.0 |
| terraform | >= 1.2.5 |

## Modules

no modules.

## Resources

| Name |
|------|
| aws_sqs_queue_policy |
| kubernetes_deployment_v1 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| namespace | The namespace to deploy the test application to | `string` | module-test-namespace | no |
| business-unit | The business unit that owns the namespace | `string` | Platforms | no |
| team_name | The team name that owns the namespace | `string` | webops | no |
| environment | The environment name that the namespace is being deployed to | `string` | development | no |
| infrastructure-support | The team name that supports the infrastructure | `string` | platforms@digital.justice.gov.uk | no |
| is_production | Is this a production namespace? | `bool` | false | no |
| slack_channel | The slack channel to post alerts to | `string` | cloud-platform | no |
| application | The name of the application | `string` | module-test-application | no |
| owner | The owner of the application | `string` | Cloud Platform: platforms@digital.justice.gov.uk | no |
| source_code | The source code of the application | `string` | github.com/ministryofjustice/cloud-platform-terraform-test-application | no |
| route53_enabled | Enable Route53 DNS for the namespace | `bool` | `false` | no |
| namespace_enabled | Enable the namespace | `bool` | `false` | no |
| rolebinding_enabled | Enable the rolebinding | `bool` | `false` | no |


## Outputs

No outputs.
<!-- END_TF_DOCS -->

## Tags

Some of the inputs for this module are tags. All infrastructure resources must be tagged to meet the MOJ Technical Guidance on [Documenting owners of infrastructure](https://technical-guidance.service.justice.gov.uk/documentation/standards/documenting-infrastructure-owners.html).

You should use your namespace variables to populate these. See the [Usage](#usage) section for more information.

## Reading Material

<!-- Add links to useful documentation -->

- [Cloud Platform user guide](https://user-guide.cloud-platform.service.justice.gov.uk/#cloud-platform-user-guide)
