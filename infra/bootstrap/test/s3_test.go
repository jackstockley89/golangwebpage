package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestBucket(t *testing.T) {
	var terraformOpts = &terraform.Options{
		TerraformDir: "./localstack",
	}
	defer terraform.Destroy(t, terraformOpts)
	terraform.InitAndApply(t, terraformOpts)
}
