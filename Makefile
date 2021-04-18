.PHONY: getprotodeps makeprotos

getprotodeps:
	which protoc
	export GO111MODULE=on
	go get -u -v google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc

makeprotos:
	export GO111MODULE=on
	protoc -I ./proto/ --go_out=. --go_opt=module=gitlab.com/insanitywholesale/bookdir --go-grpc_out=. --go-grpc_opt=module=gitlab.com/insanitywholesale/bookdir proto/bookdir/v1/*.proto
