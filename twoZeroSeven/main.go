package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from 1st")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from 2nd")
		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	wg.Wait()
	fmt.Println("About to exit")
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
}
