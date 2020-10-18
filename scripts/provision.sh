#!/bin/bash

# Initialize workers list
# A list of all the ports
# of the gocrawler workers
declare -a workers=()
declare -a worker_id_list=()

# Name of the GoCrawler cluster network
# connecting the workers
network="gocrawler_network"

# Number of workers to be provisioned
# Default 3 workers are provisioned
workers_count=$1

# Error check number of workers
# If no workers count is given default to 3
if [[ $workers_count -eq "" ]]; then
    workers_count=3
fi

echo "Number of workers: $workers_count"

# Exit when there are less than 1 worker
if [[ $workers_count -lt 1 ]]; then
    echo "Number of workers cannot be less than 1"
    exit 255
fi

# Exit when there are more than 10 workers
if [[ $workers_count -ge 10 ]]; then
    echo "Number of workers cannot be more than 10"
    exit 255
fi

# Check if port is available and then
# append to workers starting from 8000
available_port=8000
provisioned_ports_count=0

# Remove previous stale docker infra components
echo "Cleaning previous stale components"
docker ps -a | awk '$2 ~ /postgres/ {print $1}' | xargs -I {} docker rm -f {}
docker ps -a | awk '$2 ~ /rabbitmq/ {print $1}' | xargs -I {} docker rm -f {}
docker ps -a | awk '$2 ~ /gocrawler/ {print $1}' | xargs -I {} docker rm -f {}
docker ps -a | awk '$2 ~ /gosupervisor/ {print $1}' | xargs -I {} docker rm -f {}
docker network rm "$network"

# Build and deploy GoCrawler infra components

echo "Building GoCrawler Cluster Network"
docker network create "$network"

echo "Provisioning GoCrawler RabbitMQ Docker Component"
docker run -d --hostname rabbitmq --net $network -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:3-management

# Wait for RabbitMQ to come up
sleep 30

echo "Provisioning GoCrawler Postgres Docker Component"
docker build -t go-crawler-postgres -f Postgres/Dockerfile .

if [[ $? -ne 0 ]]; then
    echo "Unable To Build GoCrawler Postgres Docker Image"
    exit 255
fi

docker run -d -p 5432:5432 --net $network -e POSTGRES_PASSWORD=postgres --name postgres go-crawler-postgres

sleep 30

echo "Reserving ports for workers"

# Iterate through all ports and reserve free ports for GoCrawler workers
for port in {8000..9000}; do
    if [[ provisioned_ports_count -eq workers_count ]]; then
        break
    fi

    netstat -an | grep $port
    if [[ $? -ne 0 ]]; then
        workers+=($port)
        ((provisioned_ports_count++))
    fi
done

if [[ provisioned_ports_count -ne workers_count ]]; then
    echo "Unable to reserve ports for workers"
    exit 255
fi

echo "Reserved ports:" ${workers[*]}
comma_separated_workers=$(
    IFS=,
    echo "${workers[*]}"
)

# Docker create workers from worker list
# and pass PORT = workers[[i]]
echo "Provisioning GoCrawler Docker Cluster"

echo "Building GoCrawler Docker Image"
cd GoCrawler
docker build -t gocrawler -f Dockerfile .

if [[ $? -ne 0 ]]; then
    echo "Unable To Build GoCrawler Docker Image"
    exit 255
fi

for ((id = 0; id < $workers_count; ++id)); do
    worker_id_list+=(worker-$id)
done

comma_separated_worker_id_list=$(
    IFS=,
    echo "${worker_id_list[*]}"
)

for worker_index in "${!workers[@]}"; do
    docker run -p "${workers[$worker_index]}":8080 --net $network --name="gocrawler-$worker_index" -d gocrawler
done
cd -

echo "Building GoCrawler GoSupervisor Docker Image"
cd GoSupervisor
docker build -t gosupervisor -f Dockerfile .

if [[ $? -ne 0 ]]; then
    echo "Unable To Build GoCrawler GoSupervisor Docker Image"
    exit 255
fi

docker run -p 8050:8050 --net $network -d gosupervisor
cd -

# Docker list workers on success
echo "GoCrawler Cluster Nodes"
docker ps | grep 'gosupervisor'
docker ps | grep 'gocrawler'
docker ps | grep 'postgres'
docker ps | grep 'rabbitmq'
docker network ls | grep "$network"
