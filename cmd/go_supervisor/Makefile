
GoSupervisor-build:
	@echo "Building GoSupervisor Docker Image"	
	docker build -t gosupervisor -f Dockerfile .

GoSupervisor-run:
	@echo "Running Single GoSupervisor Docker Container"
	docker run -p 8050:8050 -d gosupervisor

build:
	@echo "Building GoSupervisor Server"	
	go build -o bin/GoSupervisor main.go

fmt:
	@echo "go fmt GoSupervisor Server"	
	go fmt ./...

test:
	@echo "Testing GoSupervisor"	
	go test -v --cover ./...

bench: 
	@echo "Benchmarking GoSupervisor /parse"	
	ab -n 5000 -c 100 -T application/json -p benchmarks/bench.json -g benchmarks/out.data http://localhost:8080/crawler/parse
	gnuplot benchmarks/apache-benchmark.p