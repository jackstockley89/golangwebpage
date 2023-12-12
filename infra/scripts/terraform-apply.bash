#!/bin/bash

set -o pipefail
set -e

# This script runs terraform apply with input set to false and no color outputs, suitable for running as part of a CI/CD pipeline.
# You need to pass through a Terraform directory as an argument, e.g.
# sh terraform-apply.bash terraform/environments

# This script pipes the output of terraform apply to ./infra/scripts/redact-output.bash to redact sensitive things, such as AWS keys if they
# are exposed via terraform apply.

# Make redact-output.bash executable
chmod +x ./infra/scripts/redact-output.bash

if [ -z "$1" ]; then
  echo "Unsure where to run terraform, exiting"
  exit 1
fi

if [ ! -z "$2" ]; then
  options="$2"
  terraform -chdir="$1" apply -input=false -no-color -auto-approve $options | ./infra/scripts/redact-output.bash
else
  terraform -chdir="$1" apply -input=false -no-color -auto-approve | ./infra/scripts/redact-output.bash
fi
