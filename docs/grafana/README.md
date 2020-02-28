# Grafana & Prometheus Configuration Explanation:

# USER STORY:
- [x] Discover configuration .yml file for Prometheus.
- [ ] Discover configuration ini for Grafana, or its .yml equiliant. 
- [x] Generate data for Prometheus, to perform a showcase.
- [x] Understand how to use Grafana.
- [ ] Create presentation for demo/training.
- [ ] Discover what variables are needed to stay constant and what variables are customizable in each configuration.

---
## Notes/Commands: 

1. Docker image pulls for prometheus and grafana:
`docker pull grafana/grafana`

`docker pull prom/prometheus`

2. Docker create for a user defined bridge network to communicate between containers and volume to persist data for Prometheus, in this project we will not be currently using the volume: 
`//docker volume create myVolume`

`docker network create myNetwork`

3. Docker run commands for each program, note that Grafana runs on port 3000, and Prometheus on port 9090:
`docker run --rm -p 9090:9090 --name=myProm --network=myNetwork -v myVolume:/etc/prometheus -d prom/prometheus`

`docker run --rm -p 3000:3000 --name=myGrafana --network=myTest -d  grafana/grafana`

---
## Configuration Notes for Prometheus:
- Once Prometheus is running, we can look at Prometheus' current configurations under the config endpoint: 
    *http://localhost:9090/config*
- A starting configuartion yml file can be found from Prometheus' Github: /docs > getting started >  
prometheus.yml is based off their starting configuration, and will be used for this demo. *Note*, the path to the .yml(Path/To/.Yml_File) will be used in the next command.
- Final step is to edit the prometheus.yml and set it up as a flag in the docker command below: 
`docker run --rm -v /home/${USER}/{Path/To/.Yml_File}/prometheus.yml:/etc/prometheus/prometheus.yml -d -p 9090:9090 --name=myProm prom/prometheus` 







