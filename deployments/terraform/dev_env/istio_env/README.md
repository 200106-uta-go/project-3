# Istio "Demo" dev environment
You may need to chmod the commands. Terraform isn't running the command at the monent
Two commands can be used to start the environment

make startMLB
This will start the load balancer and return a list of the nodes. A range containing these needs to be put
into the config.yaml. Run the next command after that is completed.

make startIstio
This will start the istio sdn, running the demo profile which will take awhile. 
You can go to the ip of the master node :32077/productpage to test that bookinfo is running

# Istio "Demo" Under KOPS Deployment:
