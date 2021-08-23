run:
	go run server/server.go

build: server/server.go
	go build -o bin/ server/server.go

clean:
	rm -f bin/server

create:
	soda create -e development

drop:
	soda drop -e development

migrate:
	soda migrate -e development

rollback:
	soda migrate down -e development
