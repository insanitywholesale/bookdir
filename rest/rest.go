package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"google.golang.org/grpc"
	"net/http"
)

//TODO: implement
func RunGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	/*
		conn, err := grpc.DialContext(
			context.Background(),
			addr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			return fmt.Errorf("failed to dial server: %w", err)
		}
	*/

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBookDirHandlerFromEndpoint(ctx, mux, ":11000", opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux)
}
