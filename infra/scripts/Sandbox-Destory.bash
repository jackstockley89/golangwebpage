#!/bin/bash

alias tf15='tfenv use 0.15.0; echo "Terraform Version: "; tfenv version-name'
alias laaops='export AWS_CONFIG_FILE=~/.aws/laa-ops/config; echo "AWS_CONFIG_FIL: "$AWS_CONFIG_FILE'

tf15
laaops

# Terraform Destory 

cd EKS
aws-vault exec laa-sandbox-lz -- terraform select sandbox
aws-vault exec laa-sandbox-lz -- terraform init
aws-vault exec laa-sandbox-lz -- terraform destory
