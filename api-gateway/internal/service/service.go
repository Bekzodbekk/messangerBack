package service

import (
	"api-gateway/genproto/messagepb"
	"api-gateway/genproto/userpb"
	"context"
)

type ServiceRepositoryClient struct {
	UserClient    userpb.UserServiceClient
	MessageClient messagepb.MessageServiceClient
}

func NewServiceRepositoryClinet(userCli *userpb.UserServiceClient, messageCli *messagepb.MessageServiceClient) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		UserClient:    *userCli,
		MessageClient: *messageCli,
	}
}

// FIXME USER-SERVICE
func (c *ServiceRepositoryClient) Register(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.ResponseInfo, error) {
	return c.UserClient.Register(ctx, req)
}
func (c *ServiceRepositoryClient) Verify(ctx context.Context, req *userpb.VerifyRequest) (*userpb.CreateUserResponse, error) {
	return c.UserClient.Verify(ctx, req)
}
func (c *ServiceRepositoryClient) Login(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error) {
	return c.UserClient.Login(ctx, req)
}
func (c *ServiceRepositoryClient) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return c.UserClient.CreateUser(ctx, req)
}
func (c *ServiceRepositoryClient) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return c.UserClient.UpdateUser(ctx, req)
}
func (c *ServiceRepositoryClient) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return c.UserClient.DeleteUser(ctx, req)
}
func (c *ServiceRepositoryClient) GetUserById(ctx context.Context, req *userpb.GetUserByIdRequest) (*userpb.GetUserByIdResponse, error) {
	return c.UserClient.GetUserById(ctx, req)
}
func (c *ServiceRepositoryClient) GetUserByFilter(ctx context.Context, req *userpb.GetUserByFilterRequest) (*userpb.GetUserByFilterResponse, error) {
	return c.UserClient.GetUserByFilter(ctx, req)
}
func (c *ServiceRepositoryClient) GetUsers(ctx context.Context, req *userpb.Void) (*userpb.GetUsersResponse, error) {
	return c.UserClient.GetUsers(ctx, req)
}

// FIXME MESSAGE-SERVICE
func (c *ServiceRepositoryClient) CreateMessage(ctx context.Context, req *messagepb.CreateMessageRequest) (*messagepb.CreateMessageResponse, error) {
	return c.MessageClient.CreateMessage(ctx, req)
}
func (c *ServiceRepositoryClient) UpdateMessage(ctx context.Context, req *messagepb.UpdateMessageRequest) (*messagepb.UpdateMessageResponse, error) {
	return c.MessageClient.UpdateMessage(ctx, req)
}
func (c *ServiceRepositoryClient) DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageRequest) (*messagepb.DeleteMessageResponse, error) {
	return c.MessageClient.DeleteMessage(ctx, req)
}
func (c *ServiceRepositoryClient) GetMessagesByTo(ctx context.Context, req *messagepb.GetMessagesByToRequest) (*messagepb.GetMessagesByToResponse, error) {
	return c.MessageClient.GetMessagesByTo(ctx, req)
}
