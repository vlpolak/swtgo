package main

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/pkg/hasher"
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

		(*s.Users).us.SaveUser(&userEntity)

		log.Printf("User was created", userEntity)

		response := &CreateUserResponse{Id: entity.User{}.Uuid.String(), UserName: entity.User{}.UserName}
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

		r, err := createLoginUserRequest(*request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		err = validateLoginUserRequest(r)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		foundUser, _ := (*s.Users).us.FindUser(r.UserName)

		log.Printf("User was found", foundUser)

		er := hasher.CheckPasswordHash(r.Password, foundUser.HashedPassword)
		log.Printf("User was found password", er)

		response := &LoginUserResponse{"https://www.google.com/"}
		err = json.NewEncoder(writer).Encode(response)

		if err != nil {
			log.Printf("Couldn't convert user to json %v", r)
			http.Error(writer, "Registration failed", http.StatusInternalServerError)
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

func createLoginUserRequest(request http.Request) (LoginUserRequest, error) {
	var r LoginUserRequest
	err := json.NewDecoder(request.Body).Decode(&r)

	if err != nil {
		log.Printf("Couldn't convert json to object %v. Request was %v", err, request)
		return LoginUserRequest{}, errors.New("invalid json")
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

func validateLoginUserRequest(r LoginUserRequest) error {
	if len(r.UserName) < 4 {
		return errors.New("request was incorrect, user name should contain at least 4 characters")
	}
	if len(r.Password) < 8 {
		return errors.New("request was incorrect, user password should contain at least 8 characters")
	}
	return nil
}

func createUser(userName, password string) (entity.User, error) {
	hp, err := hasher.HashPassword(password)
	if err != nil {
		log.Printf("Couldn't create user")
		return entity.User{}, errors.New("couldn't create user")
	}
	return entity.User{Uuid: uuid.New(), UserName: userName, HashedPassword: hp}, nil
}
