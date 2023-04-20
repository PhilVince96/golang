## Pull Docker container nats
docker pull nats

## Start a Docker Container
docker run -d --name nats -p 4222:4222 -p 8222:8222 nats

## Stop the Docker container
docker stop nats-main
