package web

type UserCreateRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}

type UserResponse struct {
	Id       string
	Username string
	Token    string
}
