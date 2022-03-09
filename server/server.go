package main

import (
	"github.com/vlpolak/swtgo/application"
	"github.com/vlpolak/swtgo/cache"
	"github.com/vlpolak/swtgo/infrastructure/persistence"
	"net/http"
)

type Users struct {
	us application.UserAppInterface
	lc cache.ActiveUsersCache
}

func NewUsers(us application.UserAppInterface) *Users {
	return &Users{
		us: us,
	}
}

type Server struct {
	Users **Users
}

func CreateServer() *Server {
	services, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}
	users := NewUsers(services.User)
	services.Automigrate()
	s := &Server{
		Users: &users,
	}
	return s
}

func (s *Server) Serve() error {
	http.HandleFunc("/", s.HandleHome)
	http.HandleFunc("/login/", s.LoginHandlerFunc)
	http.HandleFunc("/2fa/", s.Setup2FAHandlerFunc)
	http.HandleFunc("/qr.png", s.GenQRCodeHandlerFunc)
	http.HandleFunc("/verify2fa/", s.Verifi2faHandlerFunc)
	return http.ListenAndServe(":8080", nil)
}
