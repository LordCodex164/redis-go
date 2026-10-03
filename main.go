package main

import "fmt"



//keys return every key currently in the store, sorted alphabetically
 

func main() {

	s := NewStore()

	newS := s.Set("social", "x")

	fmt.Println("store", newS)

	if val, err := s.Get("social"); err != nil {
		fmt.Println("val", val)
	}

	new_s := s.Delete("social")

	fmt.Println("new_s", new_s)

	m := map[string]string{
		"age": "40",
		"name": "isaac",
	}

	m["age"] = "50"

	m["role"] = "engineer"

	fmt.Println(m)
}