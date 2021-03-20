package main

import (
	"fmt"
	"sync"
)

func main() {
	var a int
	m := new(sync.Mutex)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		m.Lock()
		a = 2
		m.Unlock()
		wg.Done()
	}()

	go func() {
		//m.Lock()
		a = 9
		//m.Unlock()
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(a)
}
