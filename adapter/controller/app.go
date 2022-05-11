package controller

type AppController struct {
	UsersController interface{ UsersController }
}
