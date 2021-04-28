module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket = "jack-lnd-terraform-state"
  acl    = "private"

  versioning = {
    enabled = true
  }

}