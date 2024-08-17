package messageservice

import (
	"fmt"
	"message-service/genproto/messagepb"
	config "message-service/internal/user/pkg/load"
	"message-service/internal/user/service"
	"net"

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
	target := fmt.Sprintf("%s:%d", conf.MessageServiceHost, conf.MessageServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	messagepb.RegisterMessageServiceServer(server, &r.service)
	return server.Serve(listener)
}
