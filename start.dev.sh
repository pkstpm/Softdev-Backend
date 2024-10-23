#!/bin/sh

# Print message before starting Docker Compose
echo "Starting Docker Compose services..."

# Start Docker Compose services in the background
docker-compose -f docker-compose-dev.yml up --build -d

# Prune all unused Docker images
echo "Pruning all unused Docker images..."
docker image prune -a -f

# Print message after Docker Compose services are up
echo "Docker Compose services are up and running."
