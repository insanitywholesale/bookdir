package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"google.golang.org/grpc"
	"net/http"
)

func RunGateway(grpcport string, restport string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBookDirHandlerFromEndpoint(ctx, mux, ":"+grpcport, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":"+restport, mux)
}
