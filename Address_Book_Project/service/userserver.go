package service

import (
	"Address_Book_Project/proto"
	"context"
)

type Server struct {
	proto.UnimplementedServiceServer
}

func (s *Server) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return Adduser(request), nil
}

func (s *Server) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return Getuser(request), nil
}

func (s *Server) UserList(ctx context.Context, request *proto.UserListRequest) (*proto.UserListResponse, error) {
	return Userlist(), nil
}

func (s *Server) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return Updateuser(request), nil
}

func (s *Server) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return Deleteuser(request), nil
}
