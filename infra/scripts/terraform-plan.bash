#!/bin/bash

set -o pipefail
set -e

# This script runs terraform plan with input set to false and no color outputs, suitable for running as part of a CI/CD pipeline.
# You need to pass through a Terraform directory as an argument, e.g.
# sh terraform-plan.bash terraform/environments

# This script pipes the output of terraform plan to ./infra/scripts/redact-output.bash to redact sensitive things, such as AWS keys if they
# are exposed via terraform plan.

# Make redact-output.bash executable
chmod +x ./infra/scripts/redact-output.bash

if [ -z "$1" ]; then
  echo "Unsure where to run terraform, exiting"
  exit 1
fi

if [ ! -z "$2" ]; then
  options="$2"
  terraform -chdir="$1" plan -input=false -no-color $options | ./infra/scripts/redact-output.bash
else
  terraform -chdir="$1" plan -input=false -no-color | ./infra/scripts/redact-output.bash
fi
