terraform {
  required_version = ">= 0.14"
  backend "s3" {
    bucket = "jack-lnd-terraform-state"
    region = "eu-west-2"
    key    = "terraform.tfstate"
  }
}