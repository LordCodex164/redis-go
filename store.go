package main

import (
	"errors"
	"fmt"
)

//not all variables are pointers but all pointers are variables
//pointers are variables that store the memory address of another variables

type Store struct {
	data map[string]string
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

func (s *Store) Set(key, value string) *Store {
	s.data[key] = value
	return s
}

func (s *Store) Get(key string) (string, error) {
	fmt.Println(s.data)
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

func (s *Store) Len() (int) {
	return len(s.data)	
}


func NewStore() *Store {

	fmt.Println("GoKV - Go key value store project")

	s := Store{
		data: map[string]string{},
	}
	
	return &s
}