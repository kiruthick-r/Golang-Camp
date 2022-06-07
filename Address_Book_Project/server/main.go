package main

import (
	"Address_Book_Project/proto"
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var address []*proto.Address

type server struct {
	proto.UnimplementedServiceServer
}

func main() {
	initAddress()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func initAddress() {

	address1 := &proto.Address{Username: "kiruthick", Address: "india", Phone: "223"}
	address2 := &proto.Address{Username: "kiruthick1", Address: "usa", Phone: "456"}
	address3 := &proto.Address{Username: "kiruthick2", Address: "canada", Phone: "890"}
	address = append(address, address1)
	address = append(address, address2)
	address = append(address, address3)
}

func (s *server) Add(ctx context.Context, request *proto.Address) (*proto.Address, error) {
	a, b, c := request.GetUsername(), request.GetAddress(), request.GetPhone()
	addressServer := &proto.Address{Username: a, Address: b, Phone: c}
	address = append(address, addressServer)
	return request, nil
}

func (s *server) GetAllData(request *proto.Empty, stream proto.Service_GetAllDataServer) error {

	for _, address := range address {
		if err := stream.Send(address); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetByUsername(ctx context.Context, request *proto.GetUser) (*proto.Address, error) {

	res := &proto.Address{}

	for _, address := range address {
		if address.Username == request.GetUsername() {
			res = address
			break
		}
	}

	return res, nil
}

func (s *server) GetByAddress(ctx context.Context, request *proto.GetAddress) (*proto.Address, error) {

	res := &proto.Address{}

	for _, address := range address {
		if address.Address == request.GetAddress() {
			res = address
			break
		}
	}

	return res, nil
}

func (s *server) GetByPhone(ctx context.Context, request *proto.GetPhone) (*proto.Address, error) {

	res := &proto.Address{}

	for _, address := range address {
		if address.Phone == request.GetPhone() {
			res = address
			break
		}
	}

	return res, nil
}

func (s *server) Update(ctx context.Context, request *proto.Address) (*proto.Address, error) {
	a, b, c := request.GetUsername(), request.GetAddress(), request.GetPhone()
	for _, address := range address {
		if address.Username == a {
			address.Address = b
			address.Phone = c
			break
		}
	}
	return request, nil
}

func (s *server) Delete(ctx context.Context, request *proto.GetUser) (*proto.Address, error) {

	address2 := &proto.Address{}
	for i, address1 := range address {
		if address1.Username == request.GetUsername() {
			address = append(address[:i], address[i+1:]...)
			address2 = address1
			break
		}
	}
	return address2, nil
}
