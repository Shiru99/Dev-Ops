provider "aws" {
  region                      = "us-east-1"
  access_key                  = "EXAMPLEKEY"
  secret_key                  = "UtnFwJalrXDENG/bPxRfEMI/K7MiCYEXAMPLEKEY"
  skip_credentials_validation = true
  skip_requesting_account_id  = true
}

resource "aws_vpc" "dev-vpc" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Project_Name = "hello-world"
    Author       = "John Doe"
  }
}
 
resource "aws_subnet" "dev-vpc-subnet" {
  vpc_id     = aws_vpc.dev-vpc.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Project_Name = "hello-world"
    Author       = "John Doe"
  }
}


# data "aws_vpc" "existing_vpc" {
#   # id = "vpc-1234567890"
#   tags = {
#     Project_Name = "hello-world"
#     Author       = "John Doe"
#   }
# }

# resource "aws_subnet" "existing_vpc-subnet" {
#   vpc_id     = data.aws_vpc.existing_vpc.id
#   cidr_block = "10.0.2.0/24"

#   tags = {
#     Project_Name = "hello-world"
#     Author       = "John Doe"
#   }
# }

output "dev-vpc-id" {
  value = aws_vpc.dev-vpc.id
}

output "dev-vpc-subnet-id" {
  value = aws_subnet.dev-vpc-subnet.id
}