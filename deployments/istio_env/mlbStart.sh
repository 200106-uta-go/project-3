kubectl apply -f https://raw.githubusercontent.com/google/metallb/v0.8.3/manifests/metallb.yaml
kubectl get nodes -o wide
echo put the range of ips into the config.yaml file
echo Then run istioStart.sh