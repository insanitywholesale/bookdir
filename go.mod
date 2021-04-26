module bookdir

go 1.16

replace gitlab.com/insanitywholesale/bookdir => ./

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/jackc/pgx/v4 v4.11.0
	github.com/lib/pq v1.10.1
	gitlab.com/insanitywholesale/bookdir v0.0.0-00010101000000-000000000000
	google.golang.org/genproto v0.0.0-20210423144448-3a41ef94ed2b
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
