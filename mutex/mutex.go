package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var GFG = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	GFG = GFG + 1
	m.Unlock()
	wg.Done()
}
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var w sync.WaitGroup

	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w, &m)
	}
	w.Wait()
	fmt.Println("Value of x", GFG)
}
