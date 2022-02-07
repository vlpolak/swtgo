package repository

import (
	"github.com/vlpolak/swtgo/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	FindUser(string) (*entity.User, error)
}
