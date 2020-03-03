##setup istio
sudo curl -L https://istio.io/downloadIstio | sh -
cd istio-1.4.5
export PATH=$PWD/bin:$PATH
sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz
tar xf helm.tar.gz
cd linux-amd64/
sudo cp helm /bin/helm
sudo cp tiller /bin/tiller
cd ..
sudo kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
sudo helm init --service-account tiller
echo "waiting for tiller pod to be ready ..."
sudo kubectl -n kube-system wait --for=condition=Ready pod -l name=tiller --timeout=300s
sudo helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
echo "waiting for istio-system jobs to complete (may take about a min)"
kubectl -n istio-system wait --for=condition=complete job --all
sudo helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml