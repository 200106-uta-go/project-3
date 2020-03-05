# Istio "Demo" dev environment
You may need to chmod the commands. Terraform isn't running the command at the monent
Three commands can be used to start the environment

./mlbStart.sh
This will start the load balancer and return a list of the nodes. A range containing these needs to be put
into the config.yaml. Run the next command after that is completed.

./istioStart.sh
This will start the istio sdn, running the demo profile which will take awhile. 
You can go to the ip of the master node :32077/productpage to test that bookinfo is running

# Istio's Demo profile *with bookinfo* set up environment step by step:

1. Download release:
The next steps occur within the master node of a Kubernetes Cluster: 
Download the lastest edition of istio:
`curl -L https://istio.io/downloadIstio | sh -`
Move to the istio package directory, at the time of this project creation it would be named istio-1.4.6:
`cd istio-1.4.6`
Add istioctl to .profile within the home directory, or export this directory to the envirnoment path:
`$ export PATH=$PWD/bin:$PATH`
is 
2. Install Istio:
This guide will use the demo profile to have a sample set of packages to be ran on istio service mesh:
Note, the demo profile that comes with istio, which includes jaeger prometheus, grafana, and kaili.
`istioctl manifest apply --set profile=demo`
Command to enable sidecar injection for new nodes on the *default* namespace:
`kubectl label namespace default istio-injection=enabled`
3. Deploy Bookinfo Application:
Change directory to where the yaml file for bookinfo configuration is located at, istio directory (istio-1.4.6):
`cd ~/istio-1.4.6`
Apply the configuration of bookinfo,and launch the bookinfo
`kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml`
4. Deploy gateway and destination rules:
`kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml`
`kubectl apply -f samples/bookinfo/networking/destination-rule-all.yaml`
5. Obtain the external ip from the istio-ingressgateway
Note this will list all services within istio, and the user must record the value of the IP.
`kubectl get services -n istio-system`
Specfic to *KOPS*, we can run this command to export the external *HOSTNAME*.
`export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')`
At this point, the productpage endpoint would be up and you could check this by visting:
http://$INGRESS_HOST/productpage

## Expose these telemetric applications to external usage:
`cd github.com/200106-uta-go/project-3/deployments/terraform/dev_env/istio_env`
`kubectl apply -f istiometrics.yaml`

## Testing enpoints:
- Kiali
`$INGRESS_HOST:15029`

- Prometheus
`$INGRESS_HOST:15030`

- Grafana
`$INGRESS_HOST:15031`

- Jaeger
`$INGRESS_HOST:15032`

## Optional Kiali set up:
1. Create username and passphrase:
Single users need to have env variables setup with *$KIALI_USERNAME* and *$KIALI_PASSPHRASE* by: 
KIALI_USERNAME=$(read -p 'Kiali Username: ' uval && echo -n $uval | base64)
KIALI_PASSPHRASE=$(read -sp 'Kiali Passphrase: ' pval && echo -n $pval | base64)

2. Add these lines to istiometrics.yaml
```sh
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: kiali
  namespace: $NAMESPACE
  labels:
    app: kiali
type: Opaque
data:
  username: $KIALI_USERNAME
  passphrase: $KIALI_PASSPHRASE
EOF
```

