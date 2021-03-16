package main

import (
	"fmt"
	"sync"
)

const count = 10000

func main() {
	var (
		counter int
		mutex   sync.Mutex
		wg      = sync.WaitGroup{}
	)

	wg.Add(count)

	for i := 0; i < count; i += 1 {
		go func() {
			mutex.Lock()
			counter += 1
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
