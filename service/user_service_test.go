package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/service/mocks"
	"testing"
)

func TestUserService_Validate(t *testing.T) {

	service := mock_entity.MockUserService{}

	err := service.Validate(entity.User{
		Uuid:           uuid.UUID{},
		UserName:       "Nickolson",
		HashedPassword: "",
	})

	assert.Nil(t, err)

}
