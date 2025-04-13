package router

import (
	"miniservice/app/internal/adapter/appmiddleware"
	"miniservice/app/internal/adapter/handler"
	"miniservice/app/pkg/jwtx"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigVariantRouter struct {
	Handler *handler.Handler
	Jwtx    jwtx.IJwtx
}

type router struct {
	server  *echo.Echo
	handler *handler.Handler
	jwtx    jwtx.IJwtx
}

func New(conf *ConfigVariantRouter) *router {
	return &router{
		server:  echo.New(),
		handler: conf.Handler,
		jwtx:    conf.Jwtx,
	}
}

func (r *router) Register() {
	route := r.server

	route.Use(appmiddleware.ReqLogger())
	route.Use(middleware.Recover())
	route.Use(middleware.CORS())

	root := route.Group("api")

	r.v1(root.Group("/v1"))

	route.Logger.Fatal(route.Start(":" + os.Getenv("PORT_RUN")))
}
