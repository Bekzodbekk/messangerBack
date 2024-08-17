package repository

import (
	"context"
	"message-service/genproto/messagepb"
)

type Repository interface {
	CreateMessage(ctx context.Context, req *messagepb.CreateMessageRequest) (*messagepb.CreateMessageResponse, error)
	UpdateMessage(ctx context.Context, req *messagepb.UpdateMessageRequest) (*messagepb.UpdateMessageResponse, error)
	DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageRequest) (*messagepb.DeleteMessageResponse, error)
	GetMessagesByTo(ctx context.Context, req *messagepb.GetMessagesByToRequest) (*messagepb.GetMessagesByToResponse, error)
}
