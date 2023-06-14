GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CONF_PROTO_FILES=$(shell find configs -name *.proto)
PROTO_PATH=./api
PROTO_FILES=$(shell find ${PROTO_PATH} -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/fynntang/protobuf-gen-go/protoc-gen-go-http@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/fynntang/protobuf-gen-go/protoc-go-inject-tags@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest


.PHONY: proto
# generate internal proto
proto:
	protoc --proto_path=${PROTO_PATH} \
	       --proto_path=./pkg/third_party/protos \
 	       --go_out=paths=source_relative:${PROTO_PATH} \
 	       --go-crud_out=paths=source_relative:${PROTO_PATH} \
	       $(PROTO_FILES)


.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...