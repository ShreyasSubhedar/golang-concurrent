package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	rnd   = rand.New(rand.NewSource(time.Now().UnixNano()))
	cache = map[int]Book{}
)

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	dbChan := make(chan Book)
	cacheChan := make(chan Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			defer wg.Done()
			book, val := queryCache(id, m)
			if val {
				ch <- book
			}

		}(id, wg, m, cacheChan)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			defer wg.Done()
			book, val := queryFromDB(id, m)
			if val {
				ch <- book
			}
		}(id, wg, m, dbChan)
		go func(dbChan, cacheChan <-chan Book) {
			select {
			case b := <-cacheChan:
				fmt.Println("--------------------From cache--------------------")
				fmt.Println(b)
				<-dbChan
			case b := <-dbChan:
				fmt.Println("--------------------From DB--------------------")
				fmt.Println(b)
			}
		}(dbChan, cacheChan)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}
func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	if ok {
		return b, true
	}

	return Book{}, false
}

func queryFromDB(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
