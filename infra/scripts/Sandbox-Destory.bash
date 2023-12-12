#!/bin/bash

alias tf101='tfenv use 1.0.1; echo "Terraform Version: "; tfenv version-name'
alias laaops='export AWS_CONFIG_FILE=~/.aws/laa-ops/config; echo "AWS_CONFIG_FIL: "$AWS_CONFIG_FILE'

# These two commands will 
# Set the correct terraform version 
# Then set the correct local aws cli credential location to apply 
tf101
laaops

# Terraform Destory 

cd ../EKS
aws-vault exec laa-sandbox-lz -- terraform workspace select sandbox
aws-vault exec laa-sandbox-lz -- terraform init
aws-vault exec laa-sandbox-lz -- terraform destroy
