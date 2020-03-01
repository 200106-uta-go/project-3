 .ONESHELL:

master:
	cd ./deployments/terraform
	terraform init
	terraform apply --auto-approve
	export masterip=$$(terraform output master_ip)
	ssh -i ./secrets/private.pem ubuntu@$$masterip

dev1:
	cd ./deployments/terraform/dev_env
	terraform init
	terraform apply --auto-approve
	export masterip=$$(terraform output -json master_ip | jq -j .[0])
	echo $$masterip
	ssh -i ./Temp.pem ubuntu@$$masterip

dev2:
	cd ./deployments/terraform/dev_env
	terraform init
	terraform apply --auto-approve
	export masterip=$$(terraform output -json master_ip | jq -j .[1])
	echo $$masterip
	ssh -i ./Temp.pem ubuntu@$$masterip

image:
	cd ./deployments/terraform/image
	terraform init
	terraform apply --auto-approve
	
destroy_dev:
	cd ./deployments/terraform/dev_env
	terraform destroy --auto-approve

