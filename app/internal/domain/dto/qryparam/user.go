package qryparam

type GetUsers struct {
	UserID   *int64  `query:"user_id"`
	Username *string `query:"username"`
	Password *string `query:"password"`
	IsActive *int32  `query:"is_active"`
	CreBy    *string `query:"cre_by"`
	CreDate  *string `query:"cre_date"`
	UpdBy    *string `query:"upd_by"`
	UpdDate  *string `query:"upd_date"`
	ProgID   *string `query:"prog_id"`
}
