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
