package controller

import (
	"context"

	"github.com/arymaulanamalik/sicepat_sample/domain/service"
	"github.com/arymaulanamalik/sicepat_sample/service/api/users"
)

type UsersControllerImpl struct {
	UsersService users.UsersService
}

type UsersController interface {
	AddUser(context.Context, service.AddUserRequest) error
}

func NewUsersController(users users.UsersService) UsersController {
	return &UsersControllerImpl{
		UsersService: users,
	}
}

func (uc *UsersControllerImpl) AddUser(ctx context.Context, req service.AddUserRequest) error {
	return uc.UsersService.AddUser(ctx, req)
}
