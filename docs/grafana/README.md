# Grafana & Prometheus Configuration Explanation

# USER STORY:
- [x] Discover configuration yaml file for Prometheus.
- [ ] Discover configuration ini for Grafana, or its yaml equiliant. 
- [ ] Generate data for Prometheus, to perform a showcase.
- [x] Understand how to use Grafana.
- [ ]Create presentation for demo/training.
- [ ] Discover what variables are needed to stay constant and what variables are customizable in each configuration.

---
## Notes/Commands: 

### Docker image pulls for prometheus and grafana:
`docker pull grafana/grafana`

`docker pull prom/prometheus`

### Docker create for a user defined bridge network to communicate between containers and volume to persist data for Prometheus, in this project we will not be currently using the volume: 
`//docker volume create myVolume`

`docker network create myNetwork`

### Docker run commands for each program, note that Grafana runs on port 3000, and Prometheus on port 9090:
`docker run --rm -p 9090:9090 --name=myProm --network=myNetwork -v myVolume:/etc/prometheus -d prom/prometheus`

`docker run --rm -p 3000:3000 --name=myGrafana --network=myTest -d  grafana/grafana`




