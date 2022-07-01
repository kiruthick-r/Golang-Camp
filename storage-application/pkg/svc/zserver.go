package svc

import (
	"context"
	"errors"
	"storage-application/pkg/pb"
	"time"

	"github.com/jinzhu/gorm"
)

// Default implementation of the StorageApplication server interface
type server struct {
	pb.UnimplementedStorageServiceServer
	db *gorm.DB
}

var StartTime time.Time

func Init() {
	StartTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(StartTime)
}

func (s *server) CommandInfo(ctx context.Context, request *pb.CommandInfoUserRequest) (*pb.CommandInfoUserResponse, error) {
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	if request.Userrequest.GetValue() == "1" {
		user := &pb.StorageUserRequest{Command: request.Userrequest.GetCommand(), Value: request.Userrequest.GetValue(), Uptime: request.Userrequest.GetUptime()}
		_, err := pb.DefaultCreateStorageUserRequest(ctx, user, s.db)
		if err != nil {
			return nil, err
		}
		return &pb.CommandInfoUserResponse{Userresponse: &pb.StorageUserResponse{Value: "Portal Service"}}, nil
	} else if request.Userrequest.GetValue() == "2" {
		uptime := uptime().String()
		user := &pb.StorageUserRequest{Command: request.Userrequest.GetCommand(), Value: request.Userrequest.GetValue(), Uptime: uptime}
		_, err := pb.DefaultCreateStorageUserRequest(ctx, user, s.db)
		if err != nil {
			return nil, err
		}
		return &pb.CommandInfoUserResponse{Userresponse: &pb.StorageUserResponse{Value: "Storage Service"}}, nil
	}
	return nil, errors.New("wrong value")
}

func (s *server) CommandUptime(ctx context.Context, request *pb.CommandUptimeUserRequest) (*pb.CommandUptimeUserResponse, error) {
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	if request.Userrequest.GetValue() == "1" {
		user := &pb.StorageUserRequest{Command: request.Userrequest.GetCommand(), Value: request.Userrequest.GetValue(), Uptime: request.Userrequest.GetUptime()}
		_, err := pb.DefaultCreateStorageUserRequest(ctx, user, s.db)
		if err != nil {
			return nil, err
		}
		return &pb.CommandUptimeUserResponse{Userresponse: &pb.StorageUserResponse{Value: request.Userrequest.GetUptime()}}, nil
	} else if request.Userrequest.GetValue() == "2" {
		uptime := uptime().String()
		user := &pb.StorageUserRequest{Command: request.Userrequest.GetCommand(), Value: request.Userrequest.GetValue(), Uptime: uptime}
		_, err := pb.DefaultCreateStorageUserRequest(ctx, user, s.db)
		if err != nil {
			return nil, err
		}
		return &pb.CommandUptimeUserResponse{Userresponse: &pb.StorageUserResponse{Value: uptime}}, nil
	}
	return nil, errors.New("wrong value")
}

func (s *server) CommandReset(ctx context.Context, request *pb.CommandResetUserRequest) (*pb.CommandResetUserResponse, error) {
	if request == nil {
		return nil, errors.New("cannot be nil")
	}
	return &pb.CommandResetUserResponse{Userresponse: &pb.StorageUserResponse{Value: "Working"}}, nil
}

// NewBasicServer returns an instance of the default server interface
func NewBasicServer(database *gorm.DB) (pb.StorageServiceServer, error) {
	return &server{db: database}, nil
}
