 .ONESHELL:

master:
	cd ./deployments/terraform
	terraform init
	terraform apply --auto-approve
	export masterip=$$(terraform output master_ip)
	ssh -i ./secrets/private.pem ubuntu@$$masterip

image:
	cd ./deployments/terraform/image
	terraform init
	terraform apply --auto-approve
	
destroy_master:
	cd ./deployments/terraform
	terraform destroy --auto-approve

destroy_image:
	cd ./deployments/terraform/image
	terraform destroy --auto-approve
