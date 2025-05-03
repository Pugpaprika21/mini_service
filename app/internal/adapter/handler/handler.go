package handler

import (
	"miniservice/app/internal/adapter/handler/user"
	"miniservice/app/internal/domain/service"
)

type Handler struct {
	UserHandler user.IUserhandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		UserHandler: user.NewUserhandler(service.UserService),
	}
}
