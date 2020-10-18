
provision:
	@echo "Provisioning GoCrawler Cluster"	
	bash scripts/provision.sh 3

info:
	@echo "GoCrawler Cluster Nodes"
	docker ps | grep 'gosupervisor' || true
	docker ps | grep 'gocrawler' || true
	docker ps | grep 'postgres' || true
	docker ps | grep 'rabbitmq' || true
	docker network ls | grep "gocrawler_network" || true

clean:
	echo "Cleaning GoCrawler Cluster"
	docker ps -a | awk '$$2 ~ /postgres/ {print $$1}' | xargs -I {} docker rm -f {} || true
	docker ps -a | awk '$$2 ~ /rabbitmq/ {print $$1}' | xargs -I {} docker rm -f {} || true
	docker ps -a | awk '$$2 ~ /gocrawler/ {print $$1}' | xargs -I {} docker rm -f {} || true
	docker ps -a | awk '$$2 ~ /gosupervisor/ {print $$1}' | xargs -I {} docker rm -f {} || true
	docker network rm "gocrawler_network" || true

start-postgres:
	@echo "Starting Postgres Database"	
	docker stop postgres || true && docker rm postgres || true
	docker build -t go-crawler-postgres -f Postgres/Dockerfile .
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name postgres go-crawler-postgres

stop-postgres:
	@echo "Stopping Postgres Database"	
	docker stop postgres || true && docker rm postgres || true

start-rabbitmq:
	@echo "Starting RabbitMQ Database"	
	docker stop rabbitmq || true && docker rm rabbitmq || true
	docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:3-management

stop-rabbitmq:
	@echo "Stopping RabbitMQ Database"	
	docker stop rabbitmq || true && docker rm rabbitmq || true
