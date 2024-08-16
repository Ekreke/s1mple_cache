package cache

import (
	"context"
	"sync"
	"time"
)

var (
// _ ICache[int, any] = (*simple.Cache[int, any])(nil)
// _ ICache[int, any] = (*lru.Cache[int, any])(nil)
)

type ICache[K comparable, V any] interface {

	// Set stores the given key-value pair in the cache.
	Set(ctx context.Context, key K, value V) error

	// Get retrieves the value associated with the given key from the cache.
	Get(ctx context.Context, key K) (V, error)

	// Delete removes the value associated with the given key from the cache.
	Delete(ctx context.Context, key K) error

	Keys() []K
}

// cache struct
type Cache[K comparable, V any] struct {
	// data store
	cache ICache[K, *Item[V]]
	// cache mutex
	mutext sync.RWMutex
	// cache cleaner
	cleaner *cleaner
}

// actually cache item
type Item[V any] struct {
	value      V
	expiration time.Time
}

func WithExpiration(exp time.Duration) ItemOption {

}

type ItemOption func(*itemOptions)

type itemOptions struct {
	expiration time.Time
}

func newItem[V any](value V, opts ...ItemOption) *Item[V] {

}
