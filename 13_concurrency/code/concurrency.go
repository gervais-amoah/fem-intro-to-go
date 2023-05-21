package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handlePlanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC!")
	}
}

func printStuff() {
	defer wg.Done()
	defer handlePlanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}

}

func main() {
	wg.Add(1)
	go printStuff()
	wg.Wait()
}
