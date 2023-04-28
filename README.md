# protobuf-gen-go

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
