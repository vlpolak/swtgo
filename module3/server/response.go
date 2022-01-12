package server

type CreateUserResponse struct {
	Id       string `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type LoginUserResponse struct {
	Url string `json:"url"`
}
