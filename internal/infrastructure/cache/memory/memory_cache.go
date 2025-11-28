package memorycache

import (
	"context"
	"sync"
	"time"

	cacheports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/cache"
)

type MemoryCache struct {
	store map[string][]byte
	mu    sync.RWMutex
}

var _ cacheports.Cache = (*MemoryCache)(nil)

func New() *MemoryCache {
	return &MemoryCache{
		store: make(map[string][]byte),
	}
}

func (m *MemoryCache) Set(ctx context.Context, key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[key] = value
	return nil
}

func (m *MemoryCache) SetTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	m.Set(ctx, key, value)

	go func() {
		select {
		case <-time.After(ttl):
			m.Delete(context.Background(), key)
		}
	}()

	return nil
}

func (m *MemoryCache) Get(ctx context.Context, key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.store[key]
	if !ok {
		return nil, nil
	}
	return v, nil
}

func (m *MemoryCache) Delete(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
	return nil
}
