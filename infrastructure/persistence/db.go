package persistence

import (
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func CreateDB() (*gorm.DB, error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "user=root password=secret dbname=postgres port=5432 sslmode=disable",
	}), &gorm.Config{})
	return db, err
}

func NewRepositories() (*Repositories, error) {
	db, err = CreateDB()
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, err
}

func (s *Repositories) Automigrate() error {
	return db.AutoMigrate(&entity.User{})
}
