package service

import (
	"Address_Book_Project/proto"

	"gorm.io/gorm"
)

var DB *gorm.DB

func addUser(request *proto.AddUserRequest) *proto.AddUserResponse {
	DB.Create(&request.User)
	return &proto.AddUserResponse{User: request.User}
}

func getUser(request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var user *proto.User
	result := DB.First(&user, request.GetId())
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.GetUserResponse{User: user}, nil
}

func userList() *proto.UserListResponse {
	var list []*proto.User
	DB.Find(&list)
	return &proto.UserListResponse{User: list}
}

func updateUser(request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	var user *proto.User
	result := DB.First(&user, request.User.GetId())
	if result.Error != nil {
		return nil, result.Error
	}
	user.Address = request.User.GetAddress()
	user.Phone = request.User.GetPhone()
	DB.Save(&user)
	return &proto.UpdateUserResponse{User: request.User}, nil
}

func deleteUser(request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	var user *proto.User
	result := DB.First(&user, request.GetId())
	if result.Error != nil {
		return nil, result.Error
	}
	DB.Delete(&user)
	return &proto.DeleteUserResponse{User: user}, nil
}
