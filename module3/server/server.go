package server

import (
	"github.com/gorilla/mux"
	"github.com/vlpolak/swtgo/module3/repository"
	"net/http"
)

type Server struct {
	Router     *mux.Router
	Repository *repository.Repository
}

func CreateServer() *Server {
	s := &Server{
		Router:     &mux.Router{},
		Repository: repository.CreateRepository(),
	}

	s.Routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":18080", s.Router)
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/user", s.HandleRegisterUser()).Methods("POST")
	s.Router.HandleFunc("/user/login", s.HandleLogin()).Methods("POST")
}
