package user

import (
	"miniservice/app/internal/domain/service/user"

	"github.com/labstack/echo/v4"
)

type IUserhandler interface {
	GetUsers(c echo.Context) error
}

type userhandler struct {
	service user.IUserService
}

func NewUserhandler(service user.IUserService) IUserhandler {
	return &userhandler{service: service}
}
