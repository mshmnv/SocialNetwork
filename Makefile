CURRENT_DIR="`pwd`"
LOCAL_BIN=$(CURRENT_DIR)/bin
PROTO_SRC=$(CURRENT_DIR)/api
PROTO_DST=$(CURRENT_DIR)/pkg/api

.PHONY: generate
generate: .generate-proto

.PHONY: .generate-proto
.generate-proto:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 && \
	buf generate --path $(PROTO_SRC)/*/*.proto

up:
	docker-compose up --build

down:
	docker-compose down

rm:
	docker-compose down --volumes
