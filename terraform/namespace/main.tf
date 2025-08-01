provider "aws" {
  region = var.aws_region
}

provider "kubernetes" {
  config_path = var.kubeconfig_path
}

# resource "aws_s3_bucket" "terraform_state" {
#   bucket = var.s3.bucket
# }

# resource "aws_s3_bucket_acl" "terraform_state_acl" {
#   bucket = aws_s3_bucket.terraform_state.id
#   acl    = "private"

#   lifecycle {
#     prevent_destroy = true
#   }
# }

# resource "aws_s3_bucket_versioning" "terraform_state_versioning" {
#   bucket = aws_s3_bucket.terraform_state.id

#   versioning_configuration {
#     status = "Enabled"
#   }
# }