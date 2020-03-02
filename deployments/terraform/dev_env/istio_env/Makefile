help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

startMLB: ## Starts the Metal Loadbalancer and then prints the Kube nodes to get the ip range
	kubectl apply -f https://raw.githubusercontent.com/google/metallb/v0.8.3/manifests/metallb.yaml
	kubectl get nodes -o wide
	echo put the range of ips into the config.yaml file
	echo Then run istioStart.sh

startIstio: ## Starts the istio-system with demo profile and opens all gateways
	kubectl apply -f config.yaml
	curl -L https://istio.io/downloadIstio | sh -
	sudo mv ./istio-1.4.5/bin/istioctl /usr/bin/istioctl
	istioctl manifest apply --set profile=demo
	kubectl label namespace default istio-injection=enabled
	kubectl apply -f ./istio-1.4.5/samples/bookinfo/platform/kube/bookinfo.yaml
	kubectl apply -f ./istio-1.4.5/samples/bookinfo/networking/bookinfo-gateway.yaml
	kubectl apply -f istiometrics.yaml
	kubectl get gateway

allDelete: ## Deletes all gateways and virtualservices
	kubectl delete gateway grafana-gateway -n istio-system
	kubectl delete virtualservice grafana-vts -n istio-system
	kubectl delete gateway jaeger-gateway -n istio-system
	kubectl delete virtualservice jaeger-vts -n istio-system

jDelete: ## Deletes the jeager gateway and virtualservice
	kubectl delete gateway jaeger-gateway -n istio-system
	kubectl delete virtualservice jaeger-vts -n istio-system

gDelete: ## Deletes the grafana gateway and virtualservice
	kubectl delete gateway grafana-gateway -n istio-system
	kubectl delete virtualservice grafana-vts -n istio-system

jApply: ## Starts the jaeger gateway & virtualservice
	kubectl apply -f jaegerGate.yaml

gApply: ## Starts the grafana gateway & virtualservice
	kubectl apply -f grafGate.yaml

clean: ## Clears the istio system
	killall istioctl
