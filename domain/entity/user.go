package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Uuid           uuid.UUID
	UserName       string
	HashedPassword string
}

type UserService interface {
	Create(user *User) (*User, error)
	Get(user *User) (*User, error)
}

type UserRepository interface {
	Save(user *User) (*User, error)
	Find(user *User) (*User, error)
	Migrate() error
}
