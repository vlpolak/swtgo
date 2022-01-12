package server

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
