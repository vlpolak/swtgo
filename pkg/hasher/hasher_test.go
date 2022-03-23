package hasher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {

	tests := []struct {
		password string
		checked  bool
	}{
		{
			password: "122234234",
			checked:  true,
		},
		{
			password: "FGghjhjh223123231",
			checked:  true,
		},
		{
			password: "jljkl12123hhihihi",
			checked:  true,
		},
		{
			password: "qeqwe23423434nmjjkjk",
			checked:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			hp, err := HashPassword(tt.password)
			assert.Nil(t, err)
			assert.NotNil(t, hp)
			ch := CheckPasswordHash(tt.password, hp)
			assert.True(t, ch)
		})
	}
}
