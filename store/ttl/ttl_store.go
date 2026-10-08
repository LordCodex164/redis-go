package ttl

import (
	"errors"
	"fmt"
	"time"
)

type TTLStore struct {
	data 		map[string]TTLEntry
	ttlDuration time.Duration
}

type TTLEntry struct {
	value 		string
	expiresAt	time.Duration
}

var ErrKeyDoesNotExist = errors.New("key does not exist")
var ErrEmptyKey = errors.New("No Keys are found")
var ErrLimitExceeded = errors.New("Max limit exceeded")

func NewTTlStore(duration time.Duration) *TTLStore {
	return &TTLStore{
		data: make(map[string]TTLEntry),
		ttlDuration: duration,
	}
}

func (t *TTLStore) Set(key, val string) (error) {

	if key != "" {
		return ErrEmptyKey
	}

	ttlEntry := TTLEntry{
		value: val,
		expiresAt: t.ttlDuration,
	}
	
	t.data[key] = ttlEntry

	return nil

}

func (t *TTLStore) Get(key string) (string, error) {
	if key != "" {
		return "", ErrEmptyKey
	}

	entry, exists := t.data[key]

	if !exists {
		return "", fmt.Errorf("key (%s) does not exist on the map", key)
	}

	return entry.value, nil
}

func (t *TTLStore) Len() (int) {
	return len(t.data)
}