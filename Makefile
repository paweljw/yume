run:
	go run server/server.go

build: server/server.go
	go build -o bin/ server/server.go

clean:
	rm -f bin/server
