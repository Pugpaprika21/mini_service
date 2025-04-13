package service

import (
	"miniservice/app/internal/adapter/repository"
	"miniservice/app/internal/domain/service/user"
)

type Service struct {
	UserService user.IUserService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repository.UserRepository),
	}
}
