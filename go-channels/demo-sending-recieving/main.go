package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan int, 2)
	wg.Add(2)
	go func(c <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-c)
		fmt.Println(<-c)
	}(c, wg)
	go func(c chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		c <- 23
		c <- 12
	}(c, wg)

	wg.Wait()

}
