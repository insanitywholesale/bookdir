.PHONY: getprotodeps protos protosplus

getprotodeps:
	which protoc
	export GO111MODULE=on
	go get -u -v google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

protos:
	protoc -I ./proto/ \
	--go_out=. --go_opt=module=gitlab.com/insanitywholesale/bookdir --go-grpc_out=. --go-grpc_opt=module=gitlab.com/insanitywholesale/bookdir proto/v1/*.proto

protosplus:
	protoc -I ./proto/ -I third_party/googleapis -I third_party/grpc-gateway \
    --go_out=. --go_opt=module=gitlab.com/insanitywholesale/bookdir \
    --go-grpc_out=. --go-grpc_opt=module=gitlab.com/insanitywholesale/bookdir \
    --openapiv2_out=./openapiv2 --openapiv2_opt logtostderr=true \
    proto/v1/*.proto
