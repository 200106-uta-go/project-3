help:
    echo Makefile contains-
	echo	startLb - downloads metallb for istio
	echo 	startIstio - runs all commands to start a demo version of istio with bookinfo
	
all:
	kubectl delete gateway grafana-gateway -n istio-system
	kubectl delete virtualservice grafana-vts -n istio-system
	kubectl delete gateway jaeger-gateway -n istio-system
	kubectl delete virtualservice jaeger-vts -n istio-system
jDelete:
	kubectl delete gateway jaeger-gateway -n istio-system
	kubectl delete virtualservice jaeger-vts -n istio-system
gDelete:
	kubectl delete gateway grafana-gateway -n istio-system
	kubectl delete virtualservice grafana-vts -n istio-system
jApply:
	kubectl apply -f jaegerGate.yaml
gApply:
	kubectl apply -f grafGate.yaml
clean:
	killall istioctl
