package persistence

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vlpolak/swtgo/domain/entity"
	"testing"
)

func TestUserRepository_SaveUser(t *testing.T) {
	db, err = CreateDB()
	ur := NewUserRepository(db)
	tests := []entity.User{{
		Uuid:           uuid.UUID{},
		UserName:       "werrerwe8989sdf8dfs",
		HashedPassword: "342342343553452332",
	}, {
		Uuid:           uuid.UUID{},
		UserName:       "Iiiufsdjsfdjjksdfkj",
		HashedPassword: "76345767645645674",
	}, {
		Uuid:           uuid.UUID{},
		UserName:       "Fsdjkkasdjjkasdjk",
		HashedPassword: "34823487823478y788y7234",
	}}

	for _, tt := range tests {
		t.Run(tt.UserName, func(t *testing.T) {
			assert.Nil(t, err)
			ur.SaveUser(&tt)
		})
	}
}

func TestUserRepository_FindUser(t *testing.T) {
	db, err = CreateDB()
	ur := NewUserRepository(db)
	tests := []string{
		"werrerwe8989sdf8dfs",
		"Iiiufsdjsfdjjksdfkj",
		"Fsdjkkasdjjkasdjk"}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			assert.Nil(t, err)
			ur.FindUser(tt)
		})
	}
}
