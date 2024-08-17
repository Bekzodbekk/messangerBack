package userservice

import (
	"api-gateway/genproto/userpb"
	"api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithUserService(conf load.Config) (*userpb.UserServiceClient, error) {
	target := fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort)
	client, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userServiceClient := userpb.NewUserServiceClient(client)
	return &userServiceClient, nil
}
