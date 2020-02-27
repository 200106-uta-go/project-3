provider "aws" {
    // access_key and secret_key are personal files to access your own
    // aws account. They should not be pushed to github and are therefore
    // in the .gitignore file.
    // FOR THIS FILE TO WORK YOU MUST INCLUDE access_key, secret_key & Temp.pem IN THE TERRAFORM FOLDER
    access_key = file("./access_key")
    secret_key = file("./secret_key")
    region = "us-east-2"
}

resource "aws_instance" "worker" {
    // *Note* it is possible this image will not exist in the future in which case it will be best
    // to start from a base ubuntu image and install kubernetes and docker in order to reach the
    // equivelant starting image.
    ami           = "ami-0920a73d71dd0ab71"
    instance_type = "t2.small"
    
    count = "2"

    #Generate your own Key_Name from AWS and use that here, name it Temp.pem
    #DO NOT UPLOAD THESE FILES, make sure they are masked by the .gitignore
    key_name = "Temp"

    security_groups = [aws_security_group.SSH.name]

    connection {
        user = "ubuntu"
        type = "ssh"
        private_key = file("./Temp.pem")
        host =  self.public_ip
        timeout = "4m"
    }

    provisioner "file" {
        source = "./join.sh"
        destination = "/home/ubuntu/join.sh"
    }

    provisioner "remote-exec" {
        inline = [
            "chmod 777 join.sh",
            "sudo ./join.sh"
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
