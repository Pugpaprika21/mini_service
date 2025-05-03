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
	// Username *string `json:"username" validate:"required"`
}
