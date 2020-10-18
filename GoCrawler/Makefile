
GoCrawler-build:
	@echo "Building GoCrawler Docker Image"	
	docker build -t gocrawler -f Dockerfile .

GoCrawler-run:
	@echo "Running Single GoCrawler Docker Container"
	docker run -p 8080:8080 -d gocrawler

build:
	@echo "Building GoCrawler Server"	
	go build -o bin/GoCrawler main.go

fmt:
	@echo "go fmt GoCrawler Server"	
	go fmt ./...

test:
	@echo "Testing GoCrawler"	
	@echo "Ensure RabbitMQ is up"	
	go test -v --cover ./...

bench: 
	@echo "Benchmarking GoCrawler /parse"	
	ab -n 5000 -c 100 -T application/json -p benchmarks/bench.json -g benchmarks/out.data http://localhost:8080/crawler/parse
	gnuplot benchmarks/apache-benchmark.p