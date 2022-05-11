package middleware

import (
	"github.com/maul/sicepat_sample/adapter/controller"
)

type MiddlewareService interface {
}

type Middleware struct {
	UserController controller.UsersController
}

func NewMiddleware(user controller.UsersController) MiddlewareService {
	return &Middleware{
		UserController: user,
	}
}
