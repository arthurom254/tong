package admin

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("item not found in registry")

var ErrAlreadyExists = errors.New("item already registered")

type Registry[T any] struct {
	mu    sync.RWMutex
	items map[string]T
}

func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{
		items: make(map[string]T),
	}
}

func (r *Registry[T]) Register(name string, item T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.items[name]; exists {
		return ErrAlreadyExists
	}

	r.items[name] = item
	return nil
}

func (r *Registry[T]) Get(name string) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.items[name]
	if !exists {
		var zero T
		return zero, ErrNotFound
	}

	return item, nil
}

func (r *Registry[T]) Remove(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.items, name)
}

func (r *Registry[T]) List() map[string]T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	copiedItems := make(map[string]T, len(r.items))
	for k, v := range r.items {
		copiedItems[k] = v
	}
	return copiedItems
}

/*

registry := NewRegistry[string]()

_ = registry.Register("apiKey", "err_pot_90")
val, _ := registry.Get("apiKey") // out err_pot_90


*/
