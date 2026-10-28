package Kv

import (
	"errors"
	"fmt"
	"sync"

	//"sync"
	"time"

	//using the base64 encoding library
	"encoding/base64"
	"strconv"
)

//not all variables are pointers but all pointers are variables
//pointers are variables that store the memory address of another variables

type Store struct {
	mu		sync.Mutex
	data    map[string]string
	maxSize int
}

func (s *Store) Keys() []string {
	//here we create an emoty slice with the capacity of the length of our data
	keys := make([]string, 0, len(s.data))
	//sort the keys auto
	for k, _ := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *Store) Rename(oldKey, newKey string) {
	fmt.Println("old_map", s.data)
	value := s.data[oldKey]
	s.data[newKey] = value
	delete(s.data, oldKey)
	fmt.Println("new_map", s.data)
}

func (s *Store) Pop(key string) (string, bool) {
	value, exists := s.data[key]
	delete(s.data, key)
	if exists {
		return value, true
	}
	return "", false
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return fmt.Errorf("key is missing")
	}

	//if the value does not exists, then it is a new value
	_, exists := s.data[key]

	//we are at capacity and is this a new key
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set Key: (%s) - %v, key, limit exceeded", key, value)
	}
	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("key is missing")
	}
	value, exists := s.data[key]
	if !exists {
		return "", errors.New("Something went wrong")
	}
	return value, nil
}

func (s *Store) Delete(key string) *Store {
	//key ->
	//s.data[key]
	delete(s.data, key)
	return s
}

func (s *Store) Len() int {
	return len(s.data)
}

func (s *Store) SetKeyWithEncryption(key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}
	return s.Get(key)
}

func NewStore(maxSize int) *Store {
	fmt.Println("GoKV - Go key value store project", maxSize)
	s := Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
	return &s
}

func (s *Store) Clone() *Store {
	cl := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}

	for k, v := range cl.data {
		fmt.Println("k", k, "v", v)
		cl.data[k] = v
	}
	return cl
}

//append the values of the store

func (s *Store) Values() []string {
	values := make([]string, len(s.data))

	for _, v := range s.data {
		values = append(values, v)
	}
	return values
}

// Incr increments the integer value stored at key by 1.
func (s *Store) Incr(key string) (int, error) {
	if key == "" {
		return 0, fmt.Errorf("key is missing")
	}

	//if we do not unlock, we are going to be prone to some deadlocks and
	//we're going to also waste resources

	s.mu.Lock()
	current := 0
	//if there is a old value, we read the value to current 
	if value, ok := s.data[key]; ok {
		v, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("error converting to string")
		}
		current = v //read from v to get the old value
	}
	current++
	time.Sleep(time.Millisecond) //simulate a task
	s.data[key] = strconv.Itoa(current)
	defer s.mu.Unlock()
	//fmt.Println("current", s.data)
	return current, nil
}

/*
Incr("pageViews")
*/
