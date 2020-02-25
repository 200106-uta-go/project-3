# Project-03

## Revature Hybrid Ingress Controller
* Create a CLI tool that functions similar to `helm init`, which creates an empty Helm chart scaffold set up with the following dependencies fully configured
    * Istio
    * Jaeger
    * Grafana

* Implement a custom Kubernetes resource `Cluster`, which represents necessary details of some Revature Kubernetes Cluster, _Cluster B_

* Create a Custom Ingress Controller that can be deployed in some Kubernetes Cluster _Cluster A_ such that
    * If a request is made to _Cluster A_ and fails for any reason, the request will be retried against the same `Service` in _Cluster B_
    * If the retried request is made to _Cluster B_, then returning a failed response is acceptable
