package main

type CreateUserRequest struct {
	UserName string `json:"name"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	UserName string `json:"name"`
	Password string `json:"password"`
}
