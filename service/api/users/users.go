package users

import (
	"context"

	"github.com/maul/sicepat_sample/domain/model"
	"github.com/maul/sicepat_sample/domain/repository"
	"github.com/maul/sicepat_sample/domain/service"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
}

type UsersService interface {
	AddUser(ctx context.Context, request service.AddUserRequest) (err error)
}

type Repository func(*UsersServiceImpl)

func UsersRepositories(Useres repository.UsersRepository) Repository {
	return func(usi *UsersServiceImpl) {
		usi.UsersRepository = Useres
	}
}

func NewUsersService(repositories ...Repository) UsersService {
	usi := &UsersServiceImpl{}

	for _, repo := range repositories {
		repo(usi)
	}

	return usi

}

func (us *UsersServiceImpl) AddUser(ctx context.Context, req service.AddUserRequest) (err error) {

	_, err = us.UsersRepository.InsertUser(ctx, model.User{
		UserID:    req.Input.UserID,
		Type:      req.Input.Type,
		CreatedBy: req.UserID,
		UpdatedBy: req.UserID,
	})

	return
}
