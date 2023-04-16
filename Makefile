CURRENT_DIR="`pwd`"
LOCAL_BIN=$(CURRENT_DIR)/bin
PROTO_SRC=$(CURRENT_DIR)/api
PROTO_DST=$(CURRENT_DIR)/pkg/api
MIGRATIONS_FOLDER=$(CURRENT_DIR)/db/migrations

.PHONY: generate
generate: .generate-proto

.PHONY: .generate-proto
.generate-proto:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 && \
	buf generate --path $(PROTO_SRC)/*/*.proto

.PHONY: up
up:
	docker-compose up --build

.PHONY: down
down:
	docker-compose down

.PHONY: rm
rm:
	docker-compose down --volumes

.PHONY: fill-db
fill-db:
	curl -X POST "http://localhost:8080/add-users"

POSTGRES_CONNECTION=user=admin password=root dbname=social_network host=localhost port=5432 sslmode=disable

.PHONY: migration-up
migration-up:
	goose -dir $(MIGRATIONS_FOLDER) postgres "$(POSTGRES_CONNECTION)" up

.PHONY: migration-down
migration-down:
	goose -dir $(MIGRATIONS_FOLDER) postgres "$(POSTGRES_CONNECTION)" down



