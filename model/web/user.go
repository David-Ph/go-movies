package web

type UserCreateRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username" form:"username"`
	Password string `validate:"required,max=200,min=1" json:"password" form:"password"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
