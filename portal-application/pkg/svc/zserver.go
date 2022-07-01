package svc

import (
	"context"
	"errors"
	"portal-application/pkg/pb"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPortalServiceServer
}

var StartTime time.Time

func Init() {
	StartTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(StartTime)
}

func clientInit() pb.StorageServiceClient {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewStorageServiceClient(conn)
	return client
}

func (s *server) Command(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	client := clientInit()
	uptime := uptime().String()
	if request.GetCommand() == "info" {
		req := &pb.CommandInfoUserRequest{Userrequest: &pb.StorageUserRequest{Command: request.GetCommand(), Value: request.GetValue(), Uptime: uptime}}
		response, err := client.CommandInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return &pb.UserResponse{Value: response.Userresponse.GetValue()}, nil
	} else if request.GetCommand() == "uptime" {
		req := &pb.CommandUptimeUserRequest{Userrequest: &pb.StorageUserRequest{Command: request.GetCommand(), Value: request.GetValue(), Uptime: uptime}}
		response, err := client.CommandUptime(ctx, req)
		if err != nil {
			return nil, err
		}
		return &pb.UserResponse{Value: response.Userresponse.GetValue()}, nil
	} else if request.GetCommand() == "reset" {
		req := &pb.CommandResetUserRequest{Userrequest: &pb.StorageUserRequest{Command: request.GetCommand(), Value: request.GetValue(), Uptime: uptime}}
		response, err := client.CommandReset(ctx, req)
		if err != nil {
			return nil, err
		}
		return &pb.UserResponse{Value: response.Userresponse.GetValue()}, nil
	}
	return nil, errors.New("wrong command")
}

// NewBasicServer returns an instance of the default server interface
func NewBasicServer() (pb.PortalServiceServer, error) {
	return &server{}, nil
}
