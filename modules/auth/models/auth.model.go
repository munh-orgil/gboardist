package auth_models

type ReqLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqRegister struct {
	Otp      uint   `json:"otp" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqForgot struct {
	Otp      uint   `json:"otp" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqChange struct {
	Old string `json:"old" validate:"required"`
	New string `json:"new" validate:"required"`
}
