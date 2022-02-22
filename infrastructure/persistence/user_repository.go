package persistence

import (
	"github.com/vlpolak/swtgo/cache"
	"github.com/vlpolak/swtgo/domain/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	userRepo := &UserRepo{db}
	cache.NewLocalCache(userRepo)
	return userRepo
}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepo) FindUser(name string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("user_name = ?", name).Take(&user).Error
	db.Where(entity.User{UserName: name}).FirstOrInit(&user)
	return &user, err
}
