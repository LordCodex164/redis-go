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

	s := CreateStore()

	s.Set("social", "x")
	s.Set("meal", "rice")

	err := s.Set("sport", "football")

	if err != nil {
		if errors.Is(err, ErrKeyDoesNotExist) {
			return
		}
		//500
		fmt.Println(err)
	}

	// val, err := s.Get("social")

	// if err != nil {
	// 	fmt.Println("val", val)
	// 	if errors.Is(err, ErrKeyDoesNotExist) {
	// 		return
	// 	}
	// 	//500
	// 	fmt.Println(err)
	// }

	// s.Delete("social")

	// m := map[string]string{
	// 	"age": "40",
	// 	"name": "isaac",
	// }

	// m["age"] = "50"

	// m["role"] = "engineer"

	encrypted, err := SetKeyWithEncryption(s, "a", "classist_conservative")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)

	//ttlStore := NewTTlStore(3)

	// encryptedTtlValue, err := SetKeyWithEncryption(*ttlStore, "a", "classist_conservative")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(encryptedTtlValue)

}

func CreateStore() *LoggingMiddleware {
	plain := kv.NewStore(4)
	logger := NewLoggingMiddleware(plain)
	metrics := NewMetricsMiddleware(plain)
	return logger

}

func SetKeyWithEncryption(store Store.IStore, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))

	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	return store.Get(key)
}
