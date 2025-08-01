terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.27.0"
    }
    github = {
      source  = "integrations/github"
      version = "~> 5.17.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.18.0"
    }
  }
  required_version = ">= 1.2.5"

  backend "s3" {
    bucket         = var.s3.bucket
    key            = var.s3.key
    region         = var.aws_region
  }
}