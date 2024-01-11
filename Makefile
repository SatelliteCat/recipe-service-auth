include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) dbname=$(POSTGRES_NAME) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.17.0

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/note_v1/note.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

docker-rebuild:
	cd docker && docker compose --env-file ../.env up -d --build

#copy-to-server:
#	scp service_linux root@45.12.231.86:

#docker-build-and-push:
#	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/olezhek28/test-server:v0.0.1 .
#	docker login -u token -p CRgAAAAA6ELvdoGnA7EL4qSbkICKeDkVldDC2OOU cr.selcloud.ru/olezhek28
#	docker push cr.selcloud.ru/olezhek28/test-server:v0.0.1
