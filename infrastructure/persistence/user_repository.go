package persistence

import (
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/domain/repository"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) FindUser(name string) (*entity.User, map[string]string) {
	var user entity.User
	err := r.db.Where("name = ?", name).Take(&user).Error
	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	err := r.db.Create(&user)
	if err != nil {
		return nil, nil
	}
	return user, nil
}
