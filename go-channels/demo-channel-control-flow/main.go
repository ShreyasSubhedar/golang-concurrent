package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		// You can check the channel value.
		// If the channel is closed the ok will be `false`
		if val, ok := <-ch; ok {
			fmt.Println(val)
		}
		// You can iterate on channel
		// and get the value.
		// If channel is not closed the you'll get a panic
		for val := range ch {
			fmt.Println(val)
		}

	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 34
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch, wg)
	wg.Wait()
}
