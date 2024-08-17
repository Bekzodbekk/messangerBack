package messageservice

import (
	"api-gateway/genproto/messagepb"
	"api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithMessageService(conf load.Config) (*messagepb.MessageServiceClient, error) {
	target := fmt.Sprintf("%s:%d", conf.MessageServiceHost, conf.MessageServicePort)
	client, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	messageServiceClient := messagepb.NewMessageServiceClient(client)
	return &messageServiceClient, nil
}
