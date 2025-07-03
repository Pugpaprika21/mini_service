package router

import (
	"github.com/labstack/echo/v4"
)

func (r *router) v1(route *echo.Group) {
	user := route.Group("/user", r.jwtx.Validate())
	{
		user.POST("/getUser", r.handler.UserHandler.GetUsers)
		user.POST("/findUser", r.handler.UserHandler.FindUser)
		user.POST("/creUsers", r.handler.UserHandler.CreUsers)
		user.POST("/updUser", r.handler.UserHandler.UpdUser)
		user.POST("/delUser", r.handler.UserHandler.DelUser)
	}
}
