package repository

import (
	"github.com/vlpolak/swtgo/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	FindUser(string) (*entity.User, map[string]string)
}
