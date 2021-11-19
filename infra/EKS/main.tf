terraform {
  required_version = ">= 1.0.1"
  backend "s3" {
    bucket = "jack-lnd-terraform-state"
    region = "eu-west-2"
    key    = "terraform.tfstate"
  }
}

data "aws_kms_alias" "ebs" {
  name = "alias/aws/ebs"
}