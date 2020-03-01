// creates one t2.medium ec2 instances with kubernetes and docker installed
output "master_ip" {
  value = aws_instance.cluster.*.public_ip
  description = "The Private IP address of the master"
}



provider "aws" {
    // access_key and secret_key are personal files to access your own
    // aws account. They should not be pushed to github and are therefore
    // in the .gitignore file.
    // FOR THIS FILE TO WORK YOU MUST INCLUDE access_key, secret_key & Temp.pem IN THE DEV_ENV FOLDER
    access_key = file("./access_key")
    secret_key = file("./secret_key")
    region = "us-east-2"
}

resource "aws_instance" "cluster" {
    ami           = "ami-0f35b7d4467861303"
    instance_type = "t2.medium"

    #Generate your own Key_Name from AWS and use that here, name it Temp.pem
    #DO NOT UPLOAD THESE FILES, make sure they are masked by the .gitignore
    key_name = "Temp"

    count = "2"

    security_groups = [aws_security_group.SSH.name]

    connection {
        user = "ubuntu"
        type = "ssh"
        private_key = file("./Temp.pem")
        host =  self.public_ip
        timeout = "15m"
    }

    provisioner "file" {
      source = "./access_key"
      destination = "/home/ubuntu/access_key"
    }

    provisioner "file" {
      source = "./secret_key"
      destination = "/home/ubuntu/secret_key"
    }

    provisioner "file" {
      source = "./Temp.pem"
      destination = "/home/ubuntu/Temp.pem"
    }

    provisioner "remote-exec"{
        inline = [
          "chmod 777 setup_dev_env.sh",
          "sudo ./setup_dev_env.sh"
        ]
    }
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