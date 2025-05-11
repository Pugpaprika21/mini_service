package user

import (
	"miniservice/app/internal/domain/service/user"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetUsers(c echo.Context) error
	CreUsers(c echo.Context) error
}

type userHandler struct {
	service user.IUserService
}

func NewUserhandler(service user.IUserService) IUserHandler {
	return &userHandler{service: service}
}
