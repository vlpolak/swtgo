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
	users          map[string]cachedUser
	userRepository repository.UserRepository
}

func NewLocalCache(ur repository.UserRepository) *LocalCache {
	lc := &LocalCache{
		users:          make(map[string]cachedUser),
		userRepository: ur,
	}
	return lc
}

func (lc *LocalCache) SaveActiveUser(u *entity.User) (*entity.User, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.users[u.UserName] = cachedUser{
		User: u,
	}
	return lc.userRepository.SaveUser(u)
}

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

func (lc *LocalCache) FindActiveUser(name string) (*entity.User, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	cu, ok := lc.users[name]
	if !ok {
		return lc.userRepository.FindUser(name)
	}
	return cu.User, nil
}
