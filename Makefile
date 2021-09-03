server: build migrate
	bin/yume server

migrate: build
	bin/yume migrate

build: exe/yume.go exe/server.go
	go build -o bin/ exe/yume.go exe/server.go

clean:
	rm -f bin/yume
