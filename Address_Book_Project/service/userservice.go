package service

import (
	"Address_Book_Project/db"
	"Address_Book_Project/proto"
)

func addUser(request *proto.AddUserRequest) *proto.AddUserResponse {
	DB := db.Init()
	DB.Create(&request.User)
	res := &proto.AddUserResponse{User: request.User}
	return res
}

func getUser(request *proto.GetUserRequest) *proto.GetUserResponse {
	var user *proto.User
	DB := db.Init()
	DB.Find(&user, request.GetID())
	res := &proto.GetUserResponse{User: user}
	return res
}

func userList() *proto.UserListResponse {
	var list []*proto.User
	DB := db.Init()
	DB.Find(&list)
	res := &proto.UserListResponse{User: list}
	return res
}

func updateUser(request *proto.UpdateUserRequest) *proto.UpdateUserResponse {
	var user *proto.User
	DB := db.Init()
	DB.First(&user, request.User.GetID())
	user.Address = request.User.GetAddress()
	user.Phone = request.User.GetPhone()
	DB.Save(&user)
	res := &proto.UpdateUserResponse{User: request.User}
	return res
}

func deleteUser(request *proto.DeleteUserRequest) *proto.DeleteUserResponse {
	var user *proto.User
	DB := db.Init()
	DB.First(&user, request.GetID())
	DB.Delete(&user)
	res := &proto.DeleteUserResponse{User: user}
	return res
}
