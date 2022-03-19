package persistence

import (
	"github.com/vlpolak/swtgo/domain/entity"
	"gorm.io/gorm"
)

type UserPersistence interface {
	SaveUser(user *entity.User) (*entity.User, error)
	FindUser(name string) (*entity.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	userRepo := &UserRepo{db}
	return userRepo
}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepo) FindUser(name string) (*entity.User, error) {
	var user entity.User
	tx := r.db.Where(entity.User{UserName: name}).FirstOrCreate(&user)
	return &user, tx.Error
}
