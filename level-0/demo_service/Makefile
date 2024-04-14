.PHONY: all build demo_service db nats publication clean

all: build db nats

build:
	go build -o bin/demo_service cmd/app/main.go 
	go build -o bin/publisher cmd/publication/publication-script.go 

demo_service:
	@./bin/demo_service

db:
	psql -U postgres -f deploy-pg.sql

nats:
	docker run -d --rm --name nats-streaming -p 4222:4222 057c745866d7 -cid demo-cluster

stop:
	docker stop nats-streaming

publication:
	@./bin/publisher

clean: stop
	rm -rf ./bin