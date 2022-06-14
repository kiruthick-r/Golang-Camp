package service

import "Address_Book_Project/proto"

var address = make(map[string]*proto.User)

func InitAddress() {
	address1 := &proto.User{Username: "kiruthick", Address: "india", Phone: "1111111111"}
	address2 := &proto.User{Username: "kiruthick1", Address: "usa", Phone: "2222222222"}
	address3 := &proto.User{Username: "kiruthick2", Address: "canada", Phone: "3333333333"}
	address[address1.Username] = address1
	address[address2.Username] = address2
	address[address3.Username] = address3
}

func Adduser(request *proto.AddRequest) *proto.AddResponse {
	a, b, c := request.User.GetUsername(), request.User.GetAddress(), request.User.GetPhone()
	addressServer := &proto.User{Username: a, Address: b, Phone: c}
	address[addressServer.Username] = addressServer
	res := &proto.AddResponse{User: addressServer}
	return res
}

func Getuser(request *proto.GetRequest) *proto.GetResponse {
	var addressServer *proto.User = address[request.GetUsername()]
	res := &proto.GetResponse{User: addressServer}
	return res
}

func Userlist() *proto.GetListResponse {
	res := &proto.GetListResponse{}
	var list []*proto.User
	for _, address1 := range address {
		list = append(list, address1)
	}
	res.User = list
	return res
}

func Updateuser(request *proto.UpdateRequest) *proto.UpdateResponse {
	a, b, c := request.User.GetUsername(), request.User.GetAddress(), request.User.GetPhone()
	addressServer := &proto.User{Username: a, Address: b, Phone: c}
	address[a] = addressServer
	res := &proto.UpdateResponse{User: addressServer}
	return res
}

func Deleteuser(request *proto.DeleteRequest) *proto.DeleteResponse {
	addressServer := &proto.User{Username: address[request.GetUsername()].Username, Address: address[request.GetUsername()].Address, Phone: address[request.GetUsername()].Phone}
	delete(address, request.GetUsername())
	res := &proto.DeleteResponse{User: addressServer}
	return res
}
