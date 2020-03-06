 .ONESHELL:

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

master: ## Not currently functional
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
	
destroy_dev: ## Tear down whole dev env
	cd ./deployments/terraform/dev_env
	export masterip=$$(terraform output -json master_ip | jq -j .[0])
	export masterip2=$$(terraform output -json master_ip | jq -j .[1])
	ssh -i ./Temp.pem ubuntu@$$masterip 'sudo terraform destroy --auto-approve' & disown
	ssh -i ./Temp.pem ubuntu@$$masterip2 'sudo terraform destroy --auto-approve'
	cd ./deployments/terraform/dev_env
	terraform destroy --auto-approve
