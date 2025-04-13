package repository

import (
	"miniservice/app/internal/adapter/repository/user"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepository user.IUserRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepository(db),
	}
}
