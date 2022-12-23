package main

import (
	"fmt"
	"sync"
)

// 1. Closed via the built-in function close
// 2. Cannot check for close channels
// 3. sending new message triggers a panic
// 4. Recieving message okay
//   - if buffered all buffered messages are available
//   - if unbuffered, or buffer empty, recieve zero value

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 23
		close(ch)
		// panic: send on closed channel
		// ch <- 23

	}(ch, wg)
	wg.Wait()
}
