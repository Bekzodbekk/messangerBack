package userservice

import (
	"fmt"
	"net"
	"user-service/genproto/userpb"
	config "user-service/internal/user/pkg/load"
	"user-service/internal/user/service"

	"google.golang.org/grpc"
)

type RunService struct {
	service service.Service
}

func NewRunSerivce(service service.Service) *RunService {
	return &RunService{
		service: service,
	}
}

func (r RunService) RUN(conf config.Config) error {
	target := fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &r.service)

	return server.Serve(listener)
}
