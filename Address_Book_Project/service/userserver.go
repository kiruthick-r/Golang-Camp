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
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	if request.User.GetUsername() == "" {
		return nil, errors.New("username cannot be empty")
	} else if request.User.GetAddress() == "" {
		return nil, errors.New("address cannot be empty")
	} else if request.User.GetPhone() == "" {
		return nil, errors.New("phone cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	return addUser(request), nil
}

func (s *Server) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	if request.GetId() == "" {
		return nil, errors.New("id cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	res, err := getUser(request)
	if err != nil {
		return nil, errors.New("id not found")
	}
	return res, nil
}

func (s *Server) UserList(ctx context.Context, request *proto.UserListRequest) (*proto.UserListResponse, error) {
	return userList(), nil
}

func (s *Server) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	if request.User.GetUsername() == "" {
		return nil, errors.New("username cannot be empty")
	} else if request.User.GetAddress() == "" {
		return nil, errors.New("address cannot be empty")
	} else if request.User.GetPhone() == "" {
		return nil, errors.New("phone cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	res, err := updateUser(request)
	if err != nil {
		return nil, errors.New("id not found")
	}
	return res, nil
}

func (s *Server) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	if request.GetId() == "" {
		return nil, errors.New("id cannot be empty")
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	res, err := deleteUser(request)
	if err != nil {
		return nil, errors.New("id not found")
	}
	return res, nil
}
