run:
	go run ./cmd/main.go

build_tcp:
	mkdir -p ./bin/socket/tcp
	go build -o ./bin/socket/tcp ./demos/socket/tcp/client
	go build -o ./bin/socket/tcp ./demos/socket/tcp/server

run_tcp_c:
	./bin/socket/tcp/client

run_tcp_s:
	./bin/socket/tcp/server