package user

import (
	"context"
	"miniservice/app/internal/domain/dto/schema"
	"miniservice/app/pkg/errors"
)

func (u *userRepository) CreUsers(ctx context.Context, params []schema.CreUsers) error {
	if err := u.db.WithContext(ctx).Table("users").Create(&params).Error; err != nil {
		return errors.WrapDBError(err)
	}
	return nil
}
