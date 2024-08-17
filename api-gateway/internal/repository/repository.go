package repository

import (
	"api-gateway/genproto/messagepb"
	"api-gateway/genproto/userpb"
	"context"
)

type ServiceRepository interface {
	// FIXME USER-SERVICE
	Register(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.ResponseInfo, error)
	Verify(ctx context.Context, req *userpb.VerifyRequest) (*userpb.CreateUserResponse, error)
	Login(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error)
	CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error)
	GetUserById(ctx context.Context, req *userpb.GetUserByIdRequest) (*userpb.GetUserByIdResponse, error)
	GetUserByFilter(ctx context.Context, req *userpb.GetUserByFilterRequest) (*userpb.GetUserByFilterResponse, error)
	GetUsers(ctx context.Context, req *userpb.Void) (*userpb.GetUsersResponse, error)

	// FIXME MESSAGE-SERVICE
	CreateMessage(ctx context.Context, req *messagepb.CreateMessageRequest) (*messagepb.CreateMessageResponse, error)
	UpdateMessage(ctx context.Context, req *messagepb.UpdateMessageRequest) (*messagepb.UpdateMessageResponse, error)
	DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageRequest) (*messagepb.DeleteMessageResponse, error)
	GetMessagesByTo(ctx context.Context, req *messagepb.GetMessagesByToRequest) (*messagepb.GetMessagesByToResponse, error)
}
