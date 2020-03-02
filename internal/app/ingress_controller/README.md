# Ingress Router

## How It Works

The Ingress Router is composed of four elements:

1. ### Ingress Manifest

    The Ingress Manifest (or sometimes called ingress) is simplly a _.yml_ (or _.json_) file that conforms to the standard rules of the Kubernetes __Kind: ingress__. It describes rules wihin it for routing control, these rules are well documented elsewhere ([Official Documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/)). Once a user creates/configures their own personal _.yml_ (or _.json_) file with the routing rules they want, they can create the ingress in the Kubernetes API server by either running the command 

        kubectl create -f path/to/ingress.yml

    ___or___ by manually curling to the Kubernetes cluster and making a POST requrest to the API server ([how to access API server](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/)) using a command similar to

        POST /apis/extensions/v1beta1/namespaces/{namespace}/ingresses

    ___Note:___ Creating an ingress by using the POST method directly above does not provide it routing rules. It simply initilizes it into the API as a referenceable object. More POST requests will need to be made. Unless needed, most users should use kubectl as its entire job is to take care of communicating with the API server. 

    Once created the ingress manifest does nothing. It has no logic, instead it may be more helpful to think of the ingress object as more of a table: as it  is just a static _.json_ file in the API server to be read from.

2. ### Ingress Config Map

    The Ingress Config Map is a critical component of the Ingress Router. It defines a file as a mountable volume for our Ingress Controller which will run inside a Docker container. This volume contains the TLS certificate config file inside it. This is used with the command 

        kubectl proxy
    
    to set up a side car proxy in our pod that can communicate with the overlay network and thus the API server. It gives admin access to the entire pod and thus all containers within the pod so they can edit and issue any command to the Kubernetes cluser. 
    
    ***DO NOT MOUNT THIS VOLUME ON ANY CONTAINER/POD THAT CAN EXECUTE CODE REMOTELY***

3. ### Ingress Controller  

    The Ingress Controller is composed of two parts. The two parts are two seperate programs that work togeather by sharing data to create the Ingress Controller functionality.

    1. #### Ingress Scanner

        The first component program is the Ingress Scanner. This program makes constant queries to the Kubernetes cluster it resides within using the TLS certificate config given by the config map.

        These queries are for three seperate kinds of data which will be used in routing:

        * Portals
        


            If any custom resource of kind portal exists this command will return the data stored within it. It uses this portal as additional routing control. This is the object that provides inter-cluster networking.

        * Ingress Manifest

                GET /apis/extensions/v1beta1/ingresses

            The Ingress Manifest returns the routing rules configured by the user. These are used to tell Proxy how to discriminate incomming requests and where to route them.

        * Services

                GET /api/v1/services

            The services are logged for their cluster IP address. This is an overlay IP address that can be dialed to contact the deployments connected to them. Any service referenced by the Ingress Manifest has its cluster IP address looked up at this stage. 

    2. #### Ingress Proxy

        The Ingress Proxy is where all the actual routing for the Ingress Controller happens. Here incomming requests are parsed for their protocol and path. If a match is found, it is routed forwards into the cluster according to the routing rules. If no match is found the request is tried against all portals in the cluster one at a time. If no portals exist in the cluster, the request is returned as a 400 error.

    

Other elements can be used to extend the functionality of this ingress controller, such as the custom resource _Portals_. These excess elements are ___Not Required___ for use of this ingress controller, only the four elements stated above are ___Required___ to be used.

## How To Configure

Below is a sample _.yml_ file that can be used to luanch both the Ingress Manifest (kind: ingress) and the Ingress Controller (kind: pod). 

    apiVersion: extensions/v1beta1
    kind: Ingress
    metadata:
        name: ingressname
    spec:
        tls:
        - hosts:
            - "*.our-server.net"
            secretName: our-server-net-tls
        rules:
        - host: www.our-server.com
            http:
            paths: 
            - path: /
                backend:
                serviceName: my-index-app-service
                servicePort: 8080
        - host: www.our-server.com
            http:
            paths: 
            - path: /videos
                backend:
                serviceName: my-video-app-service
                servicePort: 8080
    ---
    apiVersion: v1
    kind: Pod
    metadata:
        name: ingresscontroller
    labels:
        app: ingress
    spec:
        containers:
        - name: ingresscontroller
            image: jtheiss19/ingress-controller
            ports:
            - containerPort: 4000
            volumeMounts:
            - name: config-volume
                mountPath: /root/.kube/
        volumes:
        - name: config-volume
            configMap:
                name: ingress

If you look at a section of the rules within the main _.yml_ file we can show how routing rules can be configured. 

1. Host 

    This simply refers to the incomming address dialed by the user, obsuring the implyied port in this case. Any paths specified later are used against the this host discriminator. 
    
    Example: even though www.Our-Server.com/ and www.Our-Server.net/ both have the same path, they both can be routed to different service backends because of their differing host names.

2. Path

    Path describes the path, ignoring keys and fragments. It is used as the specific, unique ending to a url. 

    Example: the URL https://www.our-server.net/videos/cats?c=mykitty#00h02m30s has the path /videos/cats

3. Service Name

    Service Name describes the backend service in the Kubernetes cluster for this routing protocol. It has to be known before declaring this _.yml_ file. If you add a new service to the Kubernetes cluster and want to add it to the routing protocol you have to patch the _.yml_ file with the new service in mind.
    

    
<pre>
    - host: www.Our-Server.com
        http:
        paths: 
        - path: /
            backend:
            serviceName: MyIndexAppService
            servicePort: 8080
</pre>

Using the above example any requests to the URL http://www.Our-Server.com/?key=5#00h02m30s would redirect to our backend service MyIndexAppService whose IP address will be automatically pulled from the API server. 
