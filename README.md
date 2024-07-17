```shell
protoc --go_out=. --go-grpc_out=. proto/category.proto;
```

```shell
go install github.com/ktr0731/evans@latest
```

```shell
evans -r repl;
package pb
service CategoryService
call CategoryService
```

```shell
docker run --rm -v "$(pwd):/mount:ro" \
    ghcr.io/ktr0731/evans:latest \
      --path ./proto/files \
      --proto file-name.proto \
      --host localhost \
      --port 50051 \
      repl
```