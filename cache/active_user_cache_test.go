package cache

import (
	"github.com/google/uuid"
	"github.com/nbio/st"
	"github.com/stretchr/testify/assert"
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/logger"
	"testing"
)

func TestActiveUsersCache_Set(t *testing.T) {
	auc, err := NewActiveUsersCache()
	if err != nil {
		logger.ErrorLogger("Active user cache was not created", err).Log()
		return
	}
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
			auc.Set(tt.UserName, &tt)
		})
	}
	st.Expect(t, len(auc.users), 3)
	aus := auc.Get()
	st.Expect(t, len(aus), 3)
}
