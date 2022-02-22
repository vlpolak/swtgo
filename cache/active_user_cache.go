package cache

import (
	"errors"
	"github.com/vlpolak/swtgo/domain/entity"
	"sync"
)

type ActiveUsersCache struct {
	sync.RWMutex
	users map[string]entity.User
}

func NewActiveUsersCache() *ActiveUsersCache {
	activeUsers := make(map[string]entity.User)
	activeUsersCache := ActiveUsersCache{
		users: activeUsers,
	}
	return &activeUsersCache
}

func (c *ActiveUsersCache) Set(userName string, user entity.User) {
	c.Lock()
	defer c.Unlock()
	c.users[userName] = entity.User{
		UserName:       user.UserName,
		HashedPassword: user.HashedPassword,
		Uuid:           user.Uuid,
	}
	c.users[userName] = user
}

func (c *ActiveUsersCache) Get(userName string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	user, found := c.users[userName]
	if !found {
		return nil, false
	}
	return user, true
}

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)
