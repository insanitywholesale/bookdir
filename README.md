# bookdir
directory of books, used to illustrate useful gRPC-related development

## quick test

### json
#### get
run `go run main.go` and do `curl http://localhost:8080/api/v1/books` to see it in action

#### post
run `go run main.go` and use `curl -H "Content-Type: application/json" -X POST --data-binary '@testdata/testbook.json' http://localhost:8080/api/v1/book` to add a book
