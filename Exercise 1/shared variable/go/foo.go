// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"sync"
	"time"
)

var i = 0
var mutex sync.Mutex
var chan1 = make(chan int)
var wg sync.WaitGroup

func incrementing() {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		mutex.Lock()
		i++
		mutex.Unlock()
	}
	wg.Done()

}

func decrementing() {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000000; j++ {
		mutex.Lock()
		i--
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	wg.Add(2)

	// TODO: Spawn both functions as goroutines
	go incrementing()

	go decrementing()

	
	wg.Wait()
	//close(chan1)

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
