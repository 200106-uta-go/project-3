kubectl label namespace default istio-injection=enabled
cd istio-1.4.5
kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
kubectl exec -it $(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}') -c ratings -- curl productpage:9080/productpage | grep -o "<title>.*</title>"
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
kubectl get gateway