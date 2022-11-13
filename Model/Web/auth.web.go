package web

type LoginDto struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
