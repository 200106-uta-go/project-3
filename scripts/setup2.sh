helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
kubectl -n istio-system wait --for=condition=complete job --all
helm install install/kubernetes/helm/istio --name istio --namespace istio-system \
    --values install/kubernetes/helm/istio/values-istio-demo.yaml
