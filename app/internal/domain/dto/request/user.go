package request

type GetUsers struct {
	Lazyload *Lazyload `json:"lazyload"`
	UserID   *int64    `json:"user_id"`
	Username *string   `json:"username"`
	Password *string   `json:"password"`
	IsActive *int32    `json:"is_active"`
	CreBy    *string   `json:"cre_by"`
	CreDate  *string   `json:"cre_date"`
	UpdBy    *string   `json:"upd_by"`
	UpdDate  *string   `json:"upd_date"`
	ProgID   *string   `json:"prog_id"`
}

type FindUser struct {
	UserID   *int64  `json:"user_id" validate:"required"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	IsActive *int32  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
}

type CreUsersRows struct {
	UserID   *int64  `json:"user_id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	IsActive *int32  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
}

type CreUsers struct {
	CreUsersRows []*CreUsersRows `json:"cre_users_rows"`
}

type DelUser struct {
	UserID   *int64  `json:"user_id" validate:"required"`
	IsActive *string `json:"is_active"`
}

type DelUserIsActive struct {
	UserID   *int64  `json:"user_id" validate:"required"`
	IsActive *string `json:"is_active"`
}
