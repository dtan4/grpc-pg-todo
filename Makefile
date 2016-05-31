LDFLAGS="-w"

SERVICE_FILE := todo.proto

POSTGRES := postgres
POSTGRES_IMAGE := postgres:9.4
POSTGRES_PASSWORD := password

build:
	cd go/client && go build -ldflags=$(LDFLAGS) -o ../bin/client
	cd go/server && go build -ldflags=$(LDFLAGS) -o ../bin/server

deps:
	cd go && go get -u github.com/Masterminds/glide && glide install

generate-go:
	protoc --go_out=plugins=grpc:./go/proto $(SERVICE_FILE)

run-pg:
	@docker run \
	  -d \
	  -p 5432:5432 \
	  --name $(POSTGRES) \
	  $(POSTGRES_IMAGE)

stop-pg:
	@docker stop $(POSTGRES)
	@docker rm $(POSTGRES)

.PHONY: build deps generate-go run-pg stop-pg
