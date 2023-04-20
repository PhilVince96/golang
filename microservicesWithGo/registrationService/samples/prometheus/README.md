## Pull Docker container prom/prometheus
docker pull prom/prometheus

## Start a Docker Container
docker run -d --name prometheus-main -p 9090:9090 -v prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

## Stop the Docker container
docker stop prometheus-main
