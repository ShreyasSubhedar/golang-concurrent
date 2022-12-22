package main

import "fmt"
import "sync"

// Here the 2 goroutines are runnning independently and channel is the mediator in between,
//
//	which communicating the data transfer
//	The quote says that :
//
// Dont communicate by sharing memory, share memory by communicating
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-ch)
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 45
	}(ch, wg)
	wg.Wait()
}
