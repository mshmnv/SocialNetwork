CURRENT_DIR="`pwd`"
LOCAL_BIN=$(CURRENT_DIR)/bin
PROTO_SRC=$(CURRENT_DIR)/api
PROTO_DST=$(CURRENT_DIR)/pkg/api
MIGRATIONS_FOLDER=$(CURRENT_DIR)/db/migrations
MIGRATIONS_SHARDED_FOLDER=$(CURRENT_DIR)/db/migrations-sharded

.PHONY: generate
generate: .generate-proto

.PHONY: .generate-proto
.generate-proto:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 && \
	buf generate --path $(PROTO_SRC)/user/*.proto
	buf generate --path $(PROTO_SRC)/friend/*.proto
	buf generate --path $(PROTO_SRC)/post/*.proto
	buf generate --path $(PROTO_SRC)/dialog/*.proto

.PHONY: up
up:
	#docker-compose up -d --build
	sudo docker-compose up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: restart
restart:
	docker compose restart

.PHONY: rm
rm:
	docker-compose down --volumes

.PHONY: fill-db
fill-db:
	curl -X POST "http://localhost:8080/add-users"

POSTGRES_CONNECTION=user=admin password=root dbname=social_network host=localhost port=5432 sslmode=disable
POSTGRES_SHARD_1_CONNECTION=user=admin password=root dbname=social_network host=localhost port=5433 sslmode=disable
POSTGRES_SHARD_2_CONNECTION=user=admin password=root dbname=social_network host=localhost port=5434 sslmode=disable

.PHONY: migration-up
migration-up:
	goose -dir $(MIGRATIONS_FOLDER) postgres "$(POSTGRES_CONNECTION)" up
	goose -dir $(MIGRATIONS_SHARDED_FOLDER) postgres "$(POSTGRES_SHARD_1_CONNECTION)" up
	goose -dir $(MIGRATIONS_SHARDED_FOLDER) postgres "$(POSTGRES_SHARD_2_CONNECTION)" up

.PHONY: migration-down
migration-down:
	goose -dir $(MIGRATIONS_FOLDER) postgres "$(POSTGRES_CONNECTION)" down
	goose -dir $(MIGRATIONS_SHARDED_FOLDER) postgres "$(POSTGRES_SHARD_1_CONNECTION)" up
	goose -dir $(MIGRATIONS_SHARDED_FOLDER) postgres "$(POSTGRES_SHARD_2_CONNECTION)" up

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yaml ./...


ws-request:
	curl --output - \
         --include \
         --no-buffer \
         --header "Connection: Upgrade" \
		 --header "Upgrade: websocket" \
		 --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
		 --header "Sec-WebSocket-Version: 13" \
		 localhost:8080/post/feed/posted \
		 --cookie "session-token=909e8d05-d652-4297-855a-5ca1229a7603"
