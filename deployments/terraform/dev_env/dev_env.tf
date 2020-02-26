// creates two t2.medium ec2 instances each with kubernetes and docker installed

provider "aws" {
    // access_key and secret_key are personal files to access your own
    // aws account. They should not be pushed to github and are therefore
    // in the .gitignore file.
    // FOR THIS FILE TO WORK YOU MUST INCLUDE access_key, secret_key & Temp.pem IN THE DEV_ENV FOLDER
    access_key = file("./access_key")
    secret_key = file("./secret_key")
    region = "us-east-2"
}

resource "aws_instance" "cluster_A" {
    ami           = "ami-0a532537d305d1a34"
    instance_type = "t2.medium"

    #Generate your own Key_Name from AWS and use that here, name it Temp.pem
    #DO NOT UPLOAD THESE FILES, make sure they are masked by the .gitignore
    key_name = "Temp"

    security_groups = [aws_security_group.SSH.name]
}

resource "aws_instance" "cluster_B" {
    ami           = "ami-0a532537d305d1a34"
    instance_type = "t2.medium"

    #Generate your own Key_Name from AWS and use that here, name it Temp.pem
    #DO NOT UPLOAD THESE FILES, make sure they are masked by the .gitignore
    key_name = "Temp"

    security_groups = [aws_security_group.SSH.name]
}

resource "aws_security_group" "SSH" {
  description = "Allow SSH traffic"


  ingress {
    from_port   = 0 
    to_port     = 0
    protocol =   "-1"

    cidr_blocks =  ["0.0.0.0/0"]
  }

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}