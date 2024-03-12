package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex for synchronization

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			mu.Unlock()

			fmt.Println(incrementer)
		}(i)
	}

	wg.Wait()
	fmt.Println("end value", incrementer)
}
