package main

import (
	"Address_Book_Project/db"
	"Address_Book_Project/proto"
	"Address_Book_Project/service"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

func main() {
	service.DB = db.Init()
	go func() {
		mux := runtime.NewServeMux()
		proto.RegisterServiceHandlerServer(context.Background(), mux, &service.Server{})
		log.Fatalln(http.ListenAndServe(":8081", mux))
	}()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterServiceServer(srv, &service.Server{})

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
