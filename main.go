package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	GlobalLock sync.RWMutex
	Cache      map[int]Book
}

var (
	rnd  = rand.New(rand.NewSource(time.Now().UnixNano()))
	data = Data{}
)

func main() {
	data.Cache = make(map[int]Book, 0)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			book, val := queryCache(id)
			if val {
				fmt.Println("--------------Getting from Cache-----------------------------")
				fmt.Println(book)

			}

		}(id, wg)

		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			book, val := queryFromDB(id)
			if val {
				fmt.Println("--------------Getting from DB-----------------------------")
				fmt.Println(book)
			}
		}(id, wg)
	}
	wg.Wait()
}
func queryCache(id int) (Book, bool) {

	if b, ok := data.GetCacheLock(id); ok {
		return b, true
	}
	return Book{}, false
}

func queryFromDB(id int) (Book, bool) {
	// time.Sleep(2 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			data.AddCacheLock(b.ID, b)
			return b, true
		}
	}
	return Book{}, false
}

// ---------------Mutex Un/Locks----------------

func (d *Data) GetCacheLock(id int) (Book, bool) {
	d.GlobalLock.Lock()
	defer d.GlobalLock.Unlock()
	return d.getCachecNoLock(id)
}

func (d *Data) getCachecNoLock(id int) (Book, bool) {
	if book, ok := d.Cache[id]; ok {
		return book, ok
	}
	return Book{}, false
}

func (d *Data) AddCacheLock(id int, book Book) {
	d.GlobalLock.Lock()
	defer d.GlobalLock.Unlock()
	d.addCachecNoLock(id, book)
}

func (d *Data) addCachecNoLock(id int, book Book) {
	d.Cache[id] = book
}
