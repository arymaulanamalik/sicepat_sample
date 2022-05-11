package registry

import (
	"github.com/arymaulanamalik/sicepat_sample/adapter/controller"
	"github.com/arymaulanamalik/sicepat_sample/transport/api/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type module struct {
	mdb *mongo.Database
}

type Registry interface {
	NewAppController() controller.AppController
	Middleware() middleware.MiddlewareService
}

type Option func(*module)

func NewRegistry(options ...Option) Registry {
	m := &module{}
	for _, o := range options {
		o(m)
	}
	return m
}

func (m *module) Middleware() middleware.MiddlewareService {
	return middleware.NewMiddleware(
		m.NewUsersController(),
	)
}

func (m *module) NewAppController() controller.AppController {
	return controller.AppController{
		UsersController: m.NewUsersController(),
	}
}

func NewMongoConn(mdb *mongo.Database) Option {
	return func(m *module) {
		m.mdb = mdb
	}
}
