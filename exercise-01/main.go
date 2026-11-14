package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)

	defer cancel()

	var wg sync.WaitGroup

	//create 5 concurrent requests at the same time that take 1 second

	for i:=0; i<5; i++ {
		wg.Add(1)
		go func (id int)  {
			makeRequest(ctx, id)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func makeRequest(ctx context.Context, id int) {
	fmt.Printf("req (%d) initiated\n", id)

	//create a done channel

	done := make(chan struct{})

	go func ()  {
		time.Sleep(time.Second * 3) //simulate a network delay for 1 sec
		//close the done channel after the delay
		//closing a channel makes it possible for receivers to proceed immmediatelu
		close(done)
	}()

	//ctx.Done is literally saying when the context(in this case, timeout) is already done.


	select {
	case <-done:
		fmt.Printf("req (%d) done!\n", id)
	case <- ctx.Done():
		fmt.Printf("req (%v) failed timeout error: %d\n", id, ctx.Err())
	}
}