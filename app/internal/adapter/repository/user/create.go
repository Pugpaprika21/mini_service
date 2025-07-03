package user

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
)

func (r *userRepository) CreUsers(ctx context.Context, params []schema.CreUsers) error {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Error; err != nil {
		return errors.WrapDBError(err)
	}

	if err := tx.Table("users").Create(&params).Error; err != nil {
		tx.Rollback()
		return errors.WrapDBError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WrapDBError(err)
	}

	return nil
}
