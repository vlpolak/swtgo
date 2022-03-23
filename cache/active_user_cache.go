package cache

import (
	"errors"
	"github.com/vlpolak/swtgo/domain/entity"
	"sync"
)

type ActiveUsers interface {
	Set(userName string, user *entity.User)
	Get() (map[string]*entity.User, bool)
}

type ActiveUsersCache struct {
	sync.RWMutex
	users map[string]*entity.User
}

func NewActiveUsersCache() (*ActiveUsersCache, error) {
	activeUsers := make(map[string]*entity.User)
	activeUsersCache := ActiveUsersCache{
		users: activeUsers,
	}
	return &activeUsersCache, nil
}

func (c *ActiveUsersCache) Set(userName string, user *entity.User) {
	c.Lock()
	defer c.Unlock()
	c.users[userName] = &entity.User{
		UserName:       user.UserName,
		HashedPassword: user.HashedPassword,
		Uuid:           user.Uuid,
	}
	c.users[userName] = user
}

func (c *ActiveUsersCache) Get() (au map[string]*entity.User) {
	c.RLock()
	defer c.RUnlock()
	return c.users
}

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)
