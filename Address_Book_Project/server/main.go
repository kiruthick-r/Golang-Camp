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

type server struct {
	proto.UnimplementedServiceServer
}

func main() {
	service.InitAddress()
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
	proto.RegisterServiceServer(srv, &server{})

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) AddUser(ctx context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return service.Adduser(request), nil
}

func (s *server) GetUser(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return service.Getuser(request), nil
}

func (s *server) UserList(ctx context.Context, request *proto.Empty) (*proto.GetListResponse, error) {
	return service.Userlist(), nil
}

func (s *server) UpdateUser(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return service.Updateuser(request), nil
}

func (s *server) DeleteUser(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return service.Deleteuser(request), nil
}
