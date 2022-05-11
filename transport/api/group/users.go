package group

import (
	"github.com/maul/sicepat_sample/transport/api/handler"
	"github.com/maul/sicepat_sample/transport/api/middleware"
	"gitlab.sicepat.tech/platform/golib/router"
)

func NewUsers(r *router.MyRouter, h handler.HandlerImpl, m middleware.MiddlewareService) {
	r.POST("/users", h.AddUser)
}
