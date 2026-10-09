package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	Store "redis-app/store"
	kv "redis-app/store/kv"
)

//keys return every key currently in the store, sorted alphabetically
var ErrKeyDoesNotExist = errors.New("key does not exist")
var ErrEmptyKey = errors.New("No Keys are found")
var ErrLimitExceeded = errors.New("Max limit exceeded")


func main() {

	s, m := CreateStore()

	s.Set("social", "x")
	m.Set("social", "x")
	s.Set("meal", "rice")
	m.Set("meal", "rice")
	err := s.Set("sport", "football")

	if err != nil {
		if errors.Is(err, ErrKeyDoesNotExist) {
			return
		}
		//500
		fmt.Println(err)
	}
	m.Get("meal")
	//m.Get("meal")

	encrypted, err := SetKeyWithEncryption(s, "a", "classist_conservative")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)

	m.Report()
}

func CreateStore() (*LoggingMiddleware, *MetricsMiddleware) {
	plain := kv.NewStore(4)
	logger := NewLoggingMiddleware(plain)
	metrics := NewMetricsMiddleware(plain)
	return logger, metrics
}

func SetKeyWithEncryption(store Store.IStore, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))

	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)
}

