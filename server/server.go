package main

import (
	"github.com/gorilla/mux"
	"github.com/vlpolak/swtgo/application"
	"github.com/vlpolak/swtgo/cache"
	"github.com/vlpolak/swtgo/infrastructure/persistence"
	"github.com/vlpolak/swtgo/logger"
	"log"
	"net/http"
)

type Users struct {
	us application.UserAppInterface
	lc cache.LocalCache
}

func NewUsers(us application.UserAppInterface) *Users {
	return &Users{
		us: us,
	}
}

type Server struct {
	Router *mux.Router
	Users  **Users
}

func CreateServer() *Server {
	services, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}
	users := NewUsers(services.User)
	services.Automigrate()
	s := &Server{
		Router: &mux.Router{},
		Users:  &users,
	}
	s.Routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.HttpLogger(s.Router).ServeHTTP(w, r)
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/user", s.HandleRegisterUser()).Methods("POST")
	s.Router.HandleFunc("/user/login", s.HandleLogin()).Methods("POST")
	s.Router.HandleFunc("/user/active", s.HandleGetAvtiveUsers()).Methods("GET")
	wrappedMux := logger.HttpLogger(s.Router)
	log.Fatal(http.ListenAndServe("localhost:18080", wrappedMux))
}
