DB_URL=postgresql://root:secret@localhost:5432/waas_dev?sslmode=disable

network:
	docker network create waas-service

postgres:
	docker run --name postgres --network waas-service -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root waas_dev

dropdb:
	docker exec -it postgres dropdb waas_dev

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server proto
.PHONY: all build deploy clean build-linux

include .env

PATH_CURRENT := $(shell pwd)
PATH_BUILT := $(PATH_CURRENT)/build/server
GIT_COMMIT_LOG := $(shell git log --oneline -1 HEAD)

all: build-linux deploy clean

build-linux:
	echo "current commit: ${GIT_COMMIT_LOG}"
	env GOOS=linux GOARCH=amd64 go build -v -o ./build/server -ldflags "-X 'main.GitCommitLog=${GIT_COMMIT_LOG}'"

deploy: build-linux
	gcloud run deploy --source . --region asia-southeast1 --project ${GCP_PROJECT}; \
	echo "Done deploy."

clean:
	rm -fr "${PATH_BUILT}"; \
	echo "Clean built."

build:
	go build -v -o ./build/server-local -ldflags "-X 'main.GitCommitLog=${GIT_COMMIT_LOG}'"

dev-build: build
	./build/server-local
