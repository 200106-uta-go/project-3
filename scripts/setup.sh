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
echo "Please wait this will take about a min"
sleep 10
sudo helm init --service-account tiller
sleep 5
helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
sleep 5 
kubectl -n istio-system wait --for=condition=complete job --all
sleep 5
helm install install/kubernetes/helm/istio --name istio --namespace istio-system \
    --values install/kubernetes/helm/istio/values-istio-demo.yaml
