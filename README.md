# protobuf-gen-go

![Static Badge](https://img.shields.io/badge/Go-reference-00ADD8?style=flat-square&logo=go&logoColor=white)
![Static Badge](https://img.shields.io/badge/Go-golang-00ADD8?style=flat-square&logo=go&logoColor=white)
![Static Badge](https://img.shields.io/badge/Go%20version-%3E=1.16-00ADD8?style=flat-square&logo=go&logoColor=white)
![Static Badge](https://img.shields.io/badge/Protobuf-protocol--buffers-00979D?style=flat-square&logoColor=white)
![Static Badge](https://img.shields.io/badge/gRPC-%3E=1.2-00979D?style=flat-square&logo=grpc&logoColor=white)
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/fynntang/protobuf-gen-go?style=flat-square)

## Usage

### Install `protoc-gen-go-http`

```bash
go install github.com/fynntang/protobuf-gen-go/protoc-gen-go-http@latest
```

### protoc command

```bash
protoc --proto_path=./ \
  --go_out=paths=source_relative:./ \
  --go-http_out=paths=source_relative:./

```

### Install `protoc-go-inject-tags`

```bash
go install github.com/fynntang/protobuf-gen-go/protoc-go-inject-tags@latest

```

### command

```bash
protoc-go-inject-tags --tags=form,uri --folder=./

```

#### `protoc-go-inject-tags` parameter description

| Parameter | Description                                             | Default |
|-----------|---------------------------------------------------------|---------|
| --tags    | Used to specify the label that needs to be injected     |         |
| --folder  | Used to specify the directory where the file is located |         |
