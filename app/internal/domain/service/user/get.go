package user

import (
	"context"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (u *userService) GetUsers(ctx context.Context, req *request.GetUsers, qry *qryparam.GetUsers) ([]response.GetUsers, error) {
	var sql sqlx.Sqlx

	sql.Stmt = `SELECT user_id, username, password, cre_by, cre_date, upd_by, upd_date, prog_id, is_active FROM users `

	if req.UserID != nil && *req.UserID != 0 {
		sql.WhereClause = append(sql.WhereClause, "user_id = ?")
		sql.Args = append(sql.Args, req.UserID)
	}

	if req.Username != nil && *req.Password != "" {
		sql.WhereClause = append(sql.WhereClause, "username = ?")
		sql.Args = append(sql.Args, req.Username)
	}

	if req.Password != nil && *req.Password != "" {
		sql.WhereClause = append(sql.WhereClause, "password = ?")
		sql.Args = append(sql.Args, req.Password)
	}

	if len(sql.WhereClause) > 0 {
		sql.Stmt += " WHERE " + strings.Join(sql.WhereClause, " and ")
	}

	sql.Stmt += " ORDER BY cre_date DESC"

	rows, err := u.repository.GetUsers(ctx, sql)
	if err != nil {
		return nil, err
	}

	resp := make([]response.GetUsers, len(rows))
	for i, rec := range rows {
		resp[i] = response.GetUsers{
			UserID:   &rec.UserID.Int64,
			Username: &rec.Username.String,
			Password: &rec.Password.String,
			IsActive: &rec.IsActive.Int64,
			CreBy:    &rec.CreBy.String,
			CreDate:  &rec.CreDate.String,
			UpdBy:    &rec.UpdBy.String,
			UpdDate:  &rec.UpdDate.String,
			ProgID:   &rec.ProgID.String,
		}
	}

	return resp, nil
}
