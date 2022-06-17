package main

import (
	"Address_Book_Project/proto"
	"Address_Book_Project/service"
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:4040", "gRPC server endpoint")
)

func main() {
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		proto.RegisterServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	}()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterServiceServer(srv, &service.Server{})

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
