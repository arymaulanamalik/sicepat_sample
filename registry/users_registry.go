package registry

import (
	"github.com/arymaulanamalik/sicepat_sample/adapter/controller"
	"github.com/arymaulanamalik/sicepat_sample/domain/repository"
	"github.com/arymaulanamalik/sicepat_sample/service/api/users"
)

func (m *module) NewUsersRepository() repository.UsersRepository {
	return repository.NewUsersRepository(m.mdb)
}

func (m *module) NewUsersRegistry() users.UsersService {
	return users.NewUsersService(
		users.UsersRepositories(m.NewUsersRepository()),
	)
}

func (m *module) NewUsersController() controller.UsersController {
	return controller.NewUsersController(m.NewUsersRegistry())
}
