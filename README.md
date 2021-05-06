# bookdir
directory of books, used to illustrate useful gRPC-related development

## explanation
the api is written as a proto file and exposed over grpc and also exposed as a rest api through grpc-gateway

## way to quick test

### go test
run `go test -v` and see if it passes

### json
#### get
run `go run main.go` and do `curl http://localhost:8080/api/v1/books` to see it in action

#### post
run `go run main.go` and use `curl -H "Content-Type: application/json" -X POST --data-binary '@testdata/testbook.json' http://localhost:8080/api/v1/book` to add a book

## configuration
| Variable             | Description                    | Default Value        |
|----------------------|--------------------------------|----------------------|
| `BOOKDIR_GRPC_PORT`  | grpc port of the service       | 11000                |
| `BOOKDIR_REST_PORT`  | rest port of the service       | 8080                 |
| `PG_URL`             | postgres database URL          | empty                |
| `REDIS_URL`          | redis database URL             | empty                |

