# modification of neo's makefile
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	
jDelete:
	kubectl delete gateway tracing-gateway -n istio-system
	kubectl delete virtualservice tracing-vts -n istio-system
	kubectl delete destinationrule tracing -n istio-system
gDelete:
	kubectl delete gateway grafana-gateway -n istio-system
	kubectl delete virtualservice grafana-vts -n istio-system
	kubectl delete destinationrule grafana -n istio-system
kDelete:
	kubectl delete gateway kiali-gateway -n istio-system
	kubectl delete virtualservice kiali-vts -n istio-system
	kubectl delete destinationrule kiali -n istio-system
pDelete:
	kubectl delete gateway prometheus-gateway -n istio-system
	kubectl delete virtualservice prometheus-vts -n istio-system
	kubectl delete destinationrule prometheus -n istio-system

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
