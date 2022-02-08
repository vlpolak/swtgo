package cache

import (
	"errors"
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/domain/repository"
	"sync"
)

type cachedUser struct {
	*entity.User
}

type LocalCache struct {
	mu             sync.RWMutex
	user           map[string]cachedUser
	userRepository repository.UserRepository
}

func NewLocalCache(ur repository.UserRepository) *LocalCache {
	lc := &LocalCache{
		user:           make(map[string]cachedUser),
		userRepository: ur,
	}
	return lc
}

func (lc *LocalCache) SaveActiveUser(u *entity.User) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.user[u.UserName] = cachedUser{
		User: u,
	}
}

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

func (lc *LocalCache) FindActiveUser() map[string]cachedUser {
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	user := lc.user
	return user
}
