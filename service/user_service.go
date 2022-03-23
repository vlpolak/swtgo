package service

import (
	"errors"
	"github.com/vlpolak/swtgo/domain/entity"
	"sync"
)

var once sync.Once

type Authorization interface {
	Create(user *entity.User) (*entity.User, error)
	Get(user *entity.User) (*entity.User, error)
	Validate(user *entity.User) error
}

type UserService struct {
	userRepository entity.UserRepository
}

var instance *UserService

func NewUserService(r entity.UserRepository) entity.UserService {
	once.Do(func() {
		instance = &UserService{
			userRepository: r,
		}
	})
	return instance
}
func (u *UserService) Validate(user *entity.User) error {
	if user == nil {
		err := errors.New("The user is empty")
		return err
	}
	if user.UserName == "" {
		err := errors.New("The name of user is empty")
		return err
	}
	return nil
}

func (u *UserService) Create(user *entity.User) (*entity.User, error) {
	return u.userRepository.Save(user)
}

func (u *UserService) Get(user *entity.User) (*entity.User, error) {
	return u.userRepository.Find(user)
}
