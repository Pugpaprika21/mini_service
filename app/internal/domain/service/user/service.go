package user

import (
	"context"
	"miniservice/app/internal/adapter/repository/user"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
)

type IUserService interface {
	GetUsers(ctx context.Context, req *request.GetUsers, qry *qryparam.GetUsers) ([]response.GetUsers, int64, error)
}

type userService struct {
	repository user.IUserRepository
}

func NewUserService(repository user.IUserRepository) IUserService {
	return &userService{repository: repository}
}
