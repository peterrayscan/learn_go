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

build_udp:
	mkdir -p ./bin/socket/udp
	go build -o ./bin/socket/udp ./demos/socket/udp/client
	go build -o ./bin/socket/udp ./demos/socket/udp/server

run_udp_c:
	./bin/socket/udp/client

run_udp_s:
	./bin/socket/udp/server