resource "aws_db_instance" "cycling-db" {
  identifier                          = "cycling-blog-db"
  allocated_storage                   = 100
  auto_minor_version_upgrade          = true
  storage_type                        = "gp2"
  engine                              = "PostgresSQL"
  engine_version                      = "13.3-R1"
  instance_class                      = "db.t3.micro"
  name                                = "cycling-db"
  username                            = var.db_username
  password                            = var.db_password
  port                                = "5432"
  kms_key_id                          = ""
  storage_encrypted                   = true
  iam_database_authentication_enabled = false
  vpc_security_group_ids = [
    aws_security_group.cycling_db.id
  ]
  db_subnet_group_name    = "cycling-blog-db-sgn"
  option_group_name       = "cycling-blog-db-ogn"

  timeouts {
    create = "40m"
    delete = "40m"
    update = "80m"
  }

  tags = {
    Environment = "training"
    GithubRepo  = "golangwebpage"
  }
}

resource "aws_db_subnet_group" "cycling_db" {
  name = "cycling-db-subnet-group"
  subnet_ids = [
  
  ]
  tags = {
    Environment = "training"
    GithubRepo  = "golangwebpage"
  }
}

resource "aws_security_group" "cycling_db" {
  name        = "cycling-db-secrity-group"
  description = "Allow DB inbound traffic"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"
    cidr_blocks = [module.vpc.vpc_cidr_block]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Environment = "training"
    GithubRepo  = "golangwebpage"
  }
}

resource "aws_db_option_group" "cycling_db" {
  name                 = "cycling-postgres-db"
  engine_name          = "PostgresSQL"
  major_engine_version = "13.3-R1"

  option {
    option_name = "S3_INTEGRATION"
    port        = 0
    version     = "1.0"
  }
}
