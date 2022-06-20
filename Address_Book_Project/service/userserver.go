package service

import (
	"Address_Book_Project/proto"
	"context"
	"errors"
)

type Server struct {
	proto.UnimplementedServiceServer
}

func (s *Server) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	if request.User.GetUsername() == "" || request.User.Address == "" || request.User.Phone == "" {
		return nil, errors.New("values cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return addUser(request), nil
}

func (s *Server) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	if request.GetID() == "" {
		return nil, errors.New("username cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return getUser(request), nil
}

func (s *Server) UserList(ctx context.Context, request *proto.UserListRequest) (*proto.UserListResponse, error) {
	return userList(), nil
}

func (s *Server) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	if request.User.GetUsername() == "" || request.User.Address == "" || request.User.Phone == "" {
		return nil, errors.New("values cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return updateUser(request), nil
}

func (s *Server) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	if request.GetID() == "" {
		return nil, errors.New("username cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return deleteUser(request), nil
}
