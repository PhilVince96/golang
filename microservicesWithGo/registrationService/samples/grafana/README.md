## Pull Docker container grafana
docker pull grafana/grafana

## Start a Docker Container
docker run -d --name grafana -p 3000:3000 grafana/grafana

## Stop the Docker container
docker stop grafana
