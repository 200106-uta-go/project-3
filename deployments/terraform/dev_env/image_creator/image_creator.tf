// Use this file to create an image that can be used in dev_env.tf to create two basic t2.mediums with
// the desired starting state for development
// You need an access_key and secret_key for your own aws account in the dev_env folder

provider "aws" {
  #Two localfiles names as such. Each contains what they say, given to you from AWS.
  #DO NOT UPLOAD THESE FILES, make sure they are masked by the .gitignore
  access_key = file("../access_key")
  secret_key = file("../secret_key")
  region     = "us-east-2"
}

resource "aws_ami_from_instance" "dev_env_image" {
  name               = "dev_env_image"
  source_instance_id = "i-0f4c016c2e0ec06db" # Need to change this to the id of the instance created by base_for_image.tf
}
