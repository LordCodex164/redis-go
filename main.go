package main

import (
	"errors"
	"fmt"
)

//keys return every key currently in the store, sorted alphabetically
 
var ErrKeyDoesNotExist = errors.New("key does not exist")

func main() {

	s := NewStore(2)

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

	val, err := s.Get("social")

	if err != nil {
		fmt.Println("val", val)
		if errors.Is(err, ErrKeyDoesNotExist) {
			return
		}
		//500
		fmt.Println(err)
	}

	s.Delete("social")

	m := map[string]string{
		"age": "40",
		"name": "isaac",
	}

	m["age"] = "50"

	m["role"] = "engineer"
}