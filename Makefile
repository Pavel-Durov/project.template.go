SHELL := /bin/bash

DB_CONTAINER_NAME=my-postgres-container
DB_NAME=mydatabase
DB_USER=myuser
DB_PASSWORD=mypassword
DB_CONNECTION_STRING=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable

install:
	go mod download
	
start: db-init
	DB_CONNECTION_STRING=$(DB_CONNECTION_STRING) go run cmd/main.go

db-init:
	sqlc generate
	# Create schemas
	docker exec -i my-postgres-container psql -U $(DB_USER) -d $(DB_NAME) < ./app/services/db/sql/schema.sql

db-start:
	mkdir -p ${HOME}/postgres-data
	docker run --name $(DB_CONTAINER_NAME) -p 5432:5432 -e POSTGRES_DB=$(DB_NAME) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -v ${HOME}/postgres-data:/var/lib/postgresql/data postgres:latest

db-stop:
	docker stop $(DB_CONTAINER_NAME)

db-clean: stop-db
	docker rm $(DB_CONTAINER_NAME)

ssh-db:
	docker exec -it $(DB_CONTAINER_NAME) /bin/bash

buildmocks:
	mockery --all --dir "./app/"  --output ./test/mocks

build:
	go build -o tempalte cmd/main.go

.PHONY: test start
test:
	go test ./...
