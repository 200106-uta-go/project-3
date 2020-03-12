# modification of neo's makefile
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

jDelete: ## deletes all tracing gataeways
	kubectl delete gateway tracing-gateway -n istio-system
	kubectl delete virtualservice tracing-vts -n istio-system
	kubectl delete destinationrule tracing -n istio-system

istioStart: ## Start istio after adding ip range to config.yaml
	kubectl apply -f config.yaml
	sudo curl -L https://github.com/istio/istio/releases/download/1.4.5/istio-1.4.5-linux.tar.gz -o istio-1.4.5.tar.gz
	tar -xf istio-1.4.5.tar.gz
	sudo mv istio-1.4.5/bin/istioctl /bin/istioctl
	kubectl create ns monitoring
	istioctl manifest apply --set profile=demo
	kubectl label namespace default istio-injection=enabled
	kubectl label namespace monitoring istio-injection=enabled
	kubectl apply -f istio-1.4.5/samples/bookinfo/platform/kube/bookinfo.yaml
	kubectl apply -f istio-1.4.5/samples/bookinfo/networking/bookinfo-gateway.yaml
	kubectl apply -f istiometrics.yaml
	kubectl get gateway

mlbStart:## download and start mlb
	kubectl apply -f https://raw.githubusercontent.com/google/metallb/v0.8.3/manifests/metallb.yaml
	kubectl get nodes -o wide
	echo put the range of ips into the config.yaml file
	echo Then run istioStart.sh

gDelete: ## delete grafana gateways
	kubectl delete gateway grafana-gateway -n istio-system
	kubectl delete virtualservice grafana-vts -n istio-system
	kubectl delete destinationrule grafana -n istio-system

kDelete: ## delete kiali gateways
	kubectl delete gateway kiali-gateway -n istio-system
	kubectl delete virtualservice kiali-vts -n istio-system
	kubectl delete destinationrule kiali -n istio-system

pDelete: ## delete prometheus gateways
	kubectl delete gateway prometheus-gateway -n istio-system
	kubectl delete virtualservice prometheus-vts -n istio-system
	kubectl delete destinationrule prometheus -n istio-system

demoSetup: ## sets up environment for demo all the way through building binaries
	sudo snap install go --classic
	mkdir -p ~/go/src/github.com/200106-uta-go/project-3
	git init go/src/github.com/200106-uta-go/project-3
	cd go/src/github.com/200106-uta-go/project-3/ && git pull https://github.com/200106-uta-go/project-3
	sudo go get -d ./...
	sudo go build ./cmd/kreate
	sudo mv kreate /bin
	sudo chgrp roog /bin/kreate

delete: jDelete gDelete kDelete pDelete

jApply:
	kubectl apply -f jaegerGate.yaml

gApply:
	kubectl apply -f grafGate.yaml

kApply:
	kubectl apply -f kiaGate.yaml

pApply:
	kubectl apply -f promGate.yaml
	

apply: jApply gApply kApply pApply

clean:
	killall istioctl
