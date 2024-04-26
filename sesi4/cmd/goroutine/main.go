package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("main execution started")

	var wg sync.WaitGroup

	wg.Add(1)
	defer secondProcess(10, &wg)

	wg.Add(1)
	go firstProcess(10, &wg)

	fmt.Println("Number of goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("main execution ended")
}

func firstProcess(index int, wg *sync.WaitGroup) {
	fmt.Println("First process")

	for i := 0; i < index; i++ {
		fmt.Println("First process", i)
	}

	fmt.Println("First process done")
	wg.Done()
}

func secondProcess(index int, wg *sync.WaitGroup) {

	fmt.Println("Second process")

	for i := 0; i < index; i++ {
		fmt.Println("Second process", i)
	}

	fmt.Println("Second process done")
	wg.Done()
}

func example1() {
	go goroutine()
	fmt.Println("Hello from main")

	time.Sleep(1 * time.Second)
}

func goroutine() {
	fmt.Println("Hello from goroutine")
}
