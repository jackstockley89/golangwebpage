#!/bin/bash

alias tf15='tfenv use 0.15.0; echo "Terraform Version: "; tfenv version-name'
alias laaops='export AWS_CONFIG_FILE=~/.aws/laa-ops/config; echo "AWS_CONFIG_FIL: "$AWS_CONFIG_FILE'

# These two commands will 
# Set the correct terraform version 
# Then set the correct aws cli credential location to apply 
tf15
laaops

# Terraform Apply 
# This will select the correct workspace
# Run a init
# Then apply the Infra with the auto approve 
cd EKS
aws-vault exec laa-sandbox-lz -- terraform workspace select sandbox
aws-vault exec laa-sandbox-lz -- terraform init
aws-vault exec laa-sandbox-lz -- terraform apply -auto-approve # Auto approve plan on apply 

# Kubernetes Deploy
# Once the Terraform has completed creating the Infra 
# Next will need to apply the deployment script on the EKS cluster
## Need away of connecting to the cluster after the infra has been created to run script 
cd ../../deploy
./create-golang-cluster.bash