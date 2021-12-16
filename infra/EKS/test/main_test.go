package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformS3(t *testing.T) {
	//Test to see if S3 bucket is available
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./unit_test",
	})

	//defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	name := "jack-lnd-terraform-state"
	awsRegion := "eu-west-2"

	aws.AssertS3BucketExists(t, awsRegion, name)
}
