package service

import (
	"context"
	"message-service/genproto/messagepb"
	"message-service/internal/user/repository"
)

type Service struct {
	*messagepb.UnimplementedMessageServiceServer
	messageRepo repository.Repository
}

func NewService(messageRepo repository.Repository) *Service {
	return &Service{
		messageRepo: messageRepo,
	}
}

func (s *Service) CreateMessage(ctx context.Context, req *messagepb.CreateMessageRequest) (*messagepb.CreateMessageResponse, error) {
	return s.messageRepo.CreateMessage(ctx, req)
}
func (s *Service) UpdateMessage(ctx context.Context, req *messagepb.UpdateMessageRequest) (*messagepb.UpdateMessageResponse, error) {
	return s.messageRepo.UpdateMessage(ctx, req)
}
func (s *Service) DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageRequest) (*messagepb.DeleteMessageResponse, error) {
	return s.messageRepo.DeleteMessage(ctx, req)
}
func (s *Service) GetMessagesByTo(ctx context.Context, req *messagepb.GetMessagesByToRequest) (*messagepb.GetMessagesByToResponse, error) {
	return s.messageRepo.GetMessagesByTo(ctx, req)
}
