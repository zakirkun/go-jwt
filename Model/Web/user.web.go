package web

type UserDto struct {
	// TODO validate in controller
	Name     string `json:"name" form:"name"  validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"min=8,required"`
}

type UserResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
