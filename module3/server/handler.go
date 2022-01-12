package server

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/vlpolak/swtgo/module2/pkg/hasher"
	"github.com/vlpolak/swtgo/module3/repository"
	"log"
	"net/http"
)

const webSocketUrl = "ws://fancy-chat.io/ws&token=one-time-token"

func (s *Server) HandleRegisterUser() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")

		r, err := createCreateUserRequest(*request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		err = validateCreateUserRequest(r)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		userEntity, err := createUser(r.UserName, r.Password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = s.Repository.Update(func(tx *repository.Transaction) error {
			tx.Set(r.UserName, userEntity)
			return nil
		})
		log.Printf("User created: %v", userEntity)

		//create and return response
		response := &CreateUserResponse{Id: userEntity.Uuid.String(), UserName: userEntity.UserName}
		err = json.NewEncoder(writer).Encode(response)

		if err != nil {
			log.Printf("Couldn't convert user to json %v", r)
			http.Error(writer, "Registration failed", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) HandleLogin() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		r, err := createLoginRequest(*request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.Repository.View(func(tx *repository.Transaction) error {
			userEntity, e := tx.Get(r.UserName)
			if e != nil {
				return errors.New("no user found")
			}
			valid := hasher.CheckPasswordHash(r.Password, userEntity.HashedPassword)
			if !valid {
				return errors.New("invalid password")
			}
			return e
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		response := &LoginUserResponse{Url: webSocketUrl}
		err = json.NewEncoder(writer).Encode(response)

		if err != nil {
			log.Printf("Couldn't convert user to json %v", r)
			http.Error(writer, "Login failed", http.StatusInternalServerError)
			return
		}
	}
}

func createCreateUserRequest(request http.Request) (CreateUserRequest, error) {
	var r CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&r)

	if err != nil {
		log.Printf("Couldn't convert json to object %v. Request was %v", err, request)
		return CreateUserRequest{}, errors.New("invalid json")
	}
	return r, nil
}

func validateCreateUserRequest(r CreateUserRequest) error {
	if len(r.UserName) < 4 {
		return errors.New("request was incorrect, user name should contain at least 4 characters")
	}
	if len(r.Password) < 8 {
		return errors.New("request was incorrect, user password should contain at least 8 characters")
	}
	return nil
}

func createLoginRequest(request http.Request) (LoginUserRequest, error) {
	var r LoginUserRequest
	err := json.NewDecoder(request.Body).Decode(&r)

	if err != nil {
		log.Printf("Couldn't convert json %v, request %v", err, request)
		return LoginUserRequest{}, errors.New("Couldn't convert json")
	}
	return r, nil
}

func createUser(userName, password string) (repository.User, error) {

	hp, err := hasher.HashPassword(password)
	if err != nil {
		log.Printf("Couldn't create user")
		return repository.User{}, errors.New("couldn't create user")
	}
	return repository.User{Uuid: uuid.New(), UserName: userName, HashedPassword: hp}, nil
}
