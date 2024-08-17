package service

import (
	"context"
	"user-service/genproto/userpb"
	"user-service/internal/user/repository"
)

type Service struct {
	*userpb.UnimplementedUserServiceServer
	userRepo repository.Repository
}

func NewService(userRepo repository.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) Register(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.ResponseInfo, error) {
	return s.userRepo.Register(ctx, req)
}
func (s *Service) Verify(ctx context.Context, req *userpb.VerifyRequest) (*userpb.CreateUserResponse, error) {
	return s.userRepo.Verify(ctx, req)
}
func (s *Service) Login(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error) {
	return s.userRepo.Login(ctx, req)
}
func (s *Service) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return s.userRepo.CreateUser(ctx, req)
}
func (s *Service) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return s.userRepo.UpdateUser(ctx, req)
}
func (s *Service) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return s.userRepo.DeleteUser(ctx, req)
}
func (s *Service) GetUserById(ctx context.Context, req *userpb.GetUserByIdRequest) (*userpb.GetUserByIdResponse, error) {
	return s.userRepo.GetUserById(ctx, req)
}
func (s *Service) GetUserByFilter(ctx context.Context, req *userpb.GetUserByFilterRequest) (*userpb.GetUserByFilterResponse, error) {
	return s.userRepo.GetUserByFilter(ctx, req)
}
func (s *Service) GetUsers(ctx context.Context, req *userpb.Void) (*userpb.GetUsersResponse, error) {
	return s.userRepo.GetUsers(ctx, req)
}
