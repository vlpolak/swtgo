package service

import (
	"errors"
	"github.com/vlpolak/swtgo/domain/entity"
	"sync"
)

var once sync.Once

type userService struct {
	userRepository entity.UserRepository
}

var instance *userService

func NewUserService(r entity.UserRepository) entity.UserService {
	once.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})
	return instance
}
func (*userService) Validate(user *entity.User) error {
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

func (u *userService) Create(user *entity.User) (*entity.User, error) {
	return u.userRepository.Save(user)
}

func (u *userService) Get(user *entity.User) (*entity.User, error) {
	return u.userRepository.Find(user)
}
