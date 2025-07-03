package user

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
	"miniservice/app/pkg/sqlx"
)

func (r *userRepository) FindUser(ctx context.Context, sql sqlx.Sqlx) (*schema.FindUser, error) {
	var row schema.FindUser
	result := r.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).First(&row)
	if result.Error != nil {
		return nil, errors.WrapDBError(result.Error)
	}

	return &row, nil
}
