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
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			book, val := queryCache(id, m)
			if val {
				fmt.Println("--------------Getting from Cache-----------------------------")
				fmt.Println(book)

			}

		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			book, val := queryFromDB(id, m)
			if val {
				fmt.Println("--------------Getting from DB-----------------------------")
				fmt.Println(book)
			}
		}(id, wg, m)
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
