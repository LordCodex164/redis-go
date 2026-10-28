package main

import (
	"fmt"
	store "goroutines/store/kv"
)

// var cmds = []Command{
// 	{Op: "SET", Key: "env", value: "production"},
// 	{Op: "SET", Key: "version", value: "2.1.4"},
// 	{Op: "SET", Key: "debug", value: "true"},
// 	{Op: "SET", Key: "PageViews", value: "0"},
// 	{Op: "DELETE", Key: "env", value: "master"},
// 	{Op: "GET", Key: "env", value: ""},
// 	{Op: "INCR", Key: "PageViews", value: "0"},
// }

func main() {
	s := store.NewStore(3)
	//PopulateCmdStoreWithDefault(*s, cmds)

	s.Set("PageViews", "0")

	const hits = 1000

	//create cmds slice
	cmds := make([]Command, hits)

	//cmds -> [0, 0, ....... 0]

	for i := range cmds {
		cmds[i] = Command{Op: "INCR", Key: "PageViews"}
	}

	//fmt.Println("cmds", cmds)

	RestoreCommands(cmds, s)

	got, _ := s.Get("PageViews")

	fmt.Printf("expected: %v\n", hits)
	fmt.Printf("got: %v\n", got)

	// fmt.Println("store size", s.Len())
	// fmt.Println("store keys", s.Keys())
	// fmt.Println("store values", s)
}

// func PopulateCmdStoreWithDefault(store store.Store, cmds []Command) error {

// 	defaults := map[string]string{
// 		"env": "dev",
// 		"version": "2.9.0",
// 	}
// 	for k, v := range defaults {
// 		//for each iteration, run the goroutine
// 		go func()  {
// 		if err := store.Set(k, v); err != nil {
// 			return fmt.Errorf("PopulateDefaults: %w", err)
// 		}
// 	}
// 	}
// 	return nil
// }