package service

import (
	"context"

	"eCommerce/models"
	"eCommerce/repository"
)

type IUserService interface {
	Create(ctx context.Context, u models.User) (int, error)
	Login(ctx context.Context, userName string, passWordHashed string) (bool, int, error)
}

type UserService struct {
	IUserRepo repository.IUserRepo
}

func (_self UserService) Create(ctx context.Context, u models.User) (int, error) {
	return _self.IUserRepo.Create(ctx, u)
}

func (_self UserService) Login(ctx context.Context, userName string, passWordHashed string) (bool, int, error) {
	return _self.IUserRepo.Login(ctx, userName, passWordHashed)
}
