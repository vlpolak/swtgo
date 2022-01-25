package application

import (
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	FindUser(string) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) FindUser(name string) (*entity.User, map[string]string) {
	return u.us.FindUser(name)
}
