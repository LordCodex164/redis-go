package main

import (
	store "goroutines/store/kv"
	"log"
	"sync"
)

var (

)

type Command struct {
	Op 		string
	Key		string
	value	string
}

func dispactch(cmd Command, storer *store.Store) {
	switch cmd.Op {
	case "SET":
		log.Printf("SET key %s", cmd.Key)
		storer.Set(cmd.Key, cmd.value)
	case "GET":
		log.Printf("GET key %s", cmd.Key)
		storer.Get(cmd.Key)
	case "DELETE":
		storer.Pop(cmd.Key)
	case "INCR":
		log.Printf("INCR KEY %s", cmd.Key)
		storer.Incr(cmd.Key)
	}
}

func RestoreCommands(commands []Command, store *store.Store) {
	var wg sync.WaitGroup

	//without anycontrol, it is going to wait forever it does not communicate with the main execution
	//how do we fix this
	//this is where goroutine sychronization comes in
	for _, cmd := range commands {
		wg.Go(func (){
			dispactch(cmd, store)
		})
	}
	wg.Wait()
}