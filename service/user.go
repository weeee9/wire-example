package service

import (
	"context"
	"weeee9/wire-example/model"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

type DefaultUserService struct {
	repo model.UserRepository
}

func (service DefaultUserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return service.repo.GetAllUsers(ctx)
}

func (service DefaultUserService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return service.repo.GetUserByID(ctx, id)
}

func NewUserService(repo model.UserRepository) UserService {
	return DefaultUserService{
		repo: repo,
	}
}
