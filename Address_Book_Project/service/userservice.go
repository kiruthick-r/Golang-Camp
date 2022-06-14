package service

import (
	"Address_Book_Project/proto"
	"sync"
)

var syncMap sync.Map

func InitAddress() {
	address1 := &proto.User{Username: "kiruthick", Address: "india", Phone: "1111111111"}
	address2 := &proto.User{Username: "kiruthick1", Address: "usa", Phone: "2222222222"}
	address3 := &proto.User{Username: "kiruthick2", Address: "canada", Phone: "3333333333"}
	syncMap.Store(address1.Username, address1)
	syncMap.Store(address2.Username, address2)
	syncMap.Store(address3.Username, address3)
}

func Adduser(request *proto.AddUserRequest) *proto.AddUserResponse {
	syncMap.Store(request.User.GetUsername(), request.User)
	res := &proto.AddUserResponse{User: request.User}
	return res
}

func Getuser(request *proto.GetUserRequest) *proto.GetUserResponse {
	data, _ := syncMap.Load(request.GetUsername())
	res := &proto.GetUserResponse{User: data.(*proto.User)}
	return res
}

func Userlist() *proto.UserListResponse {
	var list []*proto.User
	syncMap.Range(func(key, value interface{}) bool {
		list = append(list, value.(*proto.User))
		return true
	})
	res := &proto.UserListResponse{}
	res.User = list
	return res
}

func Updateuser(request *proto.UpdateUserRequest) *proto.UpdateUserResponse {
	syncMap.Store(request.User.GetUsername(), request.User)
	data, _ := syncMap.Load(request.User.GetUsername())
	res := &proto.UpdateUserResponse{User: data.(*proto.User)}
	return res
}

func Deleteuser(request *proto.DeleteUserRequest) *proto.DeleteUserResponse {
	data, _ := syncMap.LoadAndDelete(request.GetUsername())
	res := &proto.DeleteUserResponse{User: data.(*proto.User)}
	return res
}
