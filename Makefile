setup:
	cp .env.dist .env && go mod vendor

vendor:
	go mod vendor

lint:
	golangci-lint run --allow-parallel-runners -c .golangci.yml

run:
	go run -mod=vendor cmd/main.go

debug:
	dlv --listen=:40000 --output /tmp/__debug_bin --headless --continue --accept-multiclient --build-flags="-v -x -gcflags '-N -l'" debug cmd/main.go