package application

import (
	"github.com/vlpolak/swtgo/cache"
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/domain/repository"
)

type userApp struct {
	us repository.UserRepository
	lc cache.ActiveUsersCache
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, error)
	FindUser(string) (*entity.User, error)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, error) {
	return u.SaveUser(user)
}

func (u *userApp) FindUser(name string) (*entity.User, error) {
	return u.FindUser(name)
}
