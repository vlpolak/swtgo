package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	mt   sync.RWMutex
	data map[string]User
}

type User struct {
	Uuid           uuid.UUID
	UserName       string
	HashedPassword string
}

func CreateRepository() *Repository {
	repository := &Repository{
		data: make(map[string]User),
	}
	return repository
}

func (repository *Repository) View(fn func(tr *Transaction) error) error {
	return repository.managed(false, fn)
}

func (repository *Repository) Update(fn func(tr *Transaction) error) error {
	return repository.managed(true, fn)
}

func (repository *Repository) managed(writable bool, fn func(tr *Transaction) error) (err error) {
	var tr *Transaction
	tr, err = repository.Begin(writable)
	if err != nil {
		return
	}

	defer repository.End(writable, tr)

	err = fn(tr)
	return
}

func (repository *Repository) Begin(writable bool) (*Transaction, error) {
	tr := &Transaction{
		repository: repository,
		writable:   writable,
	}
	tr.lock()

	return tr, nil
}

func (repository *Repository) End(writable bool, tr *Transaction) {
	if writable {
		tr.unlock()
	} else {
		tr.unlock()
	}
}

type Transaction struct {
	repository *Repository
	writable   bool
}

func (tr *Transaction) Set(key string, value User) {
	tr.repository.data[key] = value
}

func (tr *Transaction) Get(key string) (User, error) {
	v, ok := tr.repository.data[key]
	if ok {
		return v, nil
	}
	return v, errors.New(fmt.Sprintf("User %s is not found", key))
}

func (tr *Transaction) lock() {
	if tr.writable {
		tr.repository.mt.Lock()
	} else {
		tr.repository.mt.RLock()
	}
}

func (tr *Transaction) unlock() {
	if tr.writable {
		tr.repository.mt.Unlock()
	} else {
		tr.repository.mt.RUnlock()
	}
}
