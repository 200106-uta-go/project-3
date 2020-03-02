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