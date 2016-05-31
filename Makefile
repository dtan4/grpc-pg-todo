GO_DIR := go

PROTO_DIR := proto
SERVICE_FILE := todo.proto

POSTGRES := postgres
POSTGRES_IMAGE := postgres:9.4
POSTGRES_PASSWORD := password

generate-go:
	protoc --go_out=plugins=grpc:./$(GO_DIR)/$(PROTO_DIR) $(SERVICE_FILE)

run-pg:
	@docker run \
	  -d \
	  -p 5432:5432 \
	  --name $(POSTGRES) \
	  $(POSTGRES_IMAGE)

stop-pg:
	@docker stop $(POSTGRES)
	@docker rm $(POSTGRES)

.PHONY: generate-go run-pg stop-pg
