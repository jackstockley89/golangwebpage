resource "aws_db_instance" "cycling-db" {
  identifier                          = "cycling-blog-db"
  allocated_storage                   = 100
  auto_minor_version_upgrade          = true
  storage_type                        = "gp2"
  engine                              = "postgres"
  engine_version                      = "13.4"
  instance_class                      = "db.t3.micro"
  db_name                                = "cycling_db"
  username                            = var.db_username
  password                            = var.db_password
  skip_final_snapshot                 = true
  port                                = "5432"
  storage_encrypted                   = true
  iam_database_authentication_enabled = false
  vpc_security_group_ids = [
    aws_security_group.cycling_db.id
  ]
  db_subnet_group_name = aws_db_subnet_group.cycling_db.id
  option_group_name    = aws_db_option_group.cycling_db.id

  timeouts {
    create = "40m"
    delete = "40m"
    update = "80m"
  }

  tags = {
    Name        = var.app-name
    Environment = "training"
    GithubRepo  = "golangwebpage"
  }
}

resource "aws_db_subnet_group" "cycling_db" {
  name       = "cycling-db-subnet-group"
  subnet_ids = module.vpc.private_subnets

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
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [module.vpc.vpc_cidr_block]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = var.app-name
    Environment = "training"
    GithubRepo  = "golangwebpage"
  }
}

resource "aws_db_option_group" "cycling_db" {
  name                 = "cycling-db-option-group"
  engine_name          = "postgres"
  major_engine_version = "13"

  tags = {
    Name        = var.app-name
    Environment = "training"
    GithubRepo  = "golangwebpage"
  }
}
