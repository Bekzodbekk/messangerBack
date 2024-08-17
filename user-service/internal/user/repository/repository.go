package repository

import (
	"context"
	"user-service/genproto/userpb"
)

type Repository interface {
	Register(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.ResponseInfo, error)
	Verify(ctx context.Context, req *userpb.VerifyRequest) (*userpb.CreateUserResponse, error)
	Login(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error)

	CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error)
	GetUserById(ctx context.Context, req *userpb.GetUserByIdRequest) (*userpb.GetUserByIdResponse, error)
	GetUserByFilter(ctx context.Context, req *userpb.GetUserByFilterRequest) (*userpb.GetUserByFilterResponse, error)
	GetUsers(ctx context.Context, req *userpb.Void) (*userpb.GetUsersResponse, error)
}
