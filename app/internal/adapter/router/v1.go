package router

import (
	"github.com/labstack/echo/v4"
)

func (r *router) v1(route *echo.Group) {
	user := route.Group("/user", r.jwtx.Validate())
	{
		user.POST("/getUser", r.handler.UserHandler.GetUsers)
	}
}
