kubectl apply -f config.yaml
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.4.5
export PATH=$PWD/bin:$PATH
istioctl manifest apply --set profile=demo
