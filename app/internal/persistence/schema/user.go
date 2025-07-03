package schema

import "database/sql"

type GetUsers struct {
	UserID   sql.NullInt64  `gorm:"column:user_id"`
	Username sql.NullString `gorm:"column:username"`
	Password sql.NullString `gorm:"column:password"`
	CreBy    sql.NullString `gorm:"column:cre_by"`
	CreDate  sql.NullString `gorm:"column:cre_date"`
	UpdBy    sql.NullString `gorm:"column:upd_by"`
	UpdDate  sql.NullString `gorm:"column:upd_date"`
	ProgID   sql.NullString `gorm:"column:prog_id"`
	IsActive sql.NullInt64  `gorm:"column:is_active"`
	TotalRow sql.NullInt64  `gorm:"column:total_row"`
}

type FindUser struct {
	UserID   sql.NullInt64  `gorm:"column:user_id"`
	Username sql.NullString `gorm:"column:username"`
	Password sql.NullString `gorm:"column:password"`
	CreBy    sql.NullString `gorm:"column:cre_by"`
	CreDate  sql.NullString `gorm:"column:cre_date"`
	UpdBy    sql.NullString `gorm:"column:upd_by"`
	UpdDate  sql.NullString `gorm:"column:upd_date"`
	ProgID   sql.NullString `gorm:"column:prog_id"`
	IsActive sql.NullInt64  `gorm:"column:is_active"`
}

type CreUsers struct {
	UserID   *int64  `gorm:"column:user_id"`
	Username *string `gorm:"column:username"`
	Password *string `gorm:"column:password"`
	CreBy    *string `gorm:"column:cre_by"`
	CreDate  *string `gorm:"column:cre_date"`
	UpdBy    *string `gorm:"column:upd_by"`
	UpdDate  *string `gorm:"column:upd_date"`
	ProgID   *string `gorm:"column:prog_id"`
	IsActive *int32  `gorm:"column:is_active"`
}

type UpdUser struct {
	UserID   *int64  `gorm:"column:user_id"`
	Username *string `gorm:"column:username"`
	Password *string `gorm:"column:password"`
	CreBy    *string `gorm:"column:cre_by"`
	CreDate  *string `gorm:"column:cre_date"`
	UpdBy    *string `gorm:"column:upd_by"`
	UpdDate  *string `gorm:"column:upd_date"`
	ProgID   *string `gorm:"column:prog_id"`
	IsActive *int32  `gorm:"column:is_active"`
}
