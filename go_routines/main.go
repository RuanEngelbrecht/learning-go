package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // used to synchronize the execution of multiple go routines
var counter = 0
var m = sync.RWMutex{} // lock and unlock the mutex

func main() {
	// 'go' -> spin off green thread -> very cheap to create and destroy in go
	// go sayHello()
	// time.Sleep(100 * time.Millisecond) // Not best practice to sleep in main thread. Don't use this in production.

	// var msg = "Hello Anonymous!"
	// wg.Add(1) // sync the execution of the go routine
	// go func(msg string) {
	// 	fmt.Println(msg) // closure -> function has access to variables in the outer function.
	// 	// We've created a dependency on the msg variable in main i.e. race condition. Not a good idea.
	// 	// Rather use parameters to pass in variables
	// 	wg.Done() // decrement number of wait groups
	// }(msg)
	// wg.Wait() // continue execution

	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)      // add 2 wait groups
	// 	go sayHello()  // routines are racing against each other
	// 	go increment() // need to synchronize the execution of the go routines -> mutex
	// }
	// wg.Wait()

	// runtime.GOMAXPROCS(100) // set the number of go routines to run concurrently -> threads
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // read lock
		go sayHello()
		m.Lock() // write lock
		go increment()
	}
	wg.Wait()
}

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	wg.Done()
// }

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // unlock
	wg.Done()
}

func increment() {
	counter++
	m.Unlock() // unlock
	wg.Done()
}

/*
	Best Practices:
	- Don't create goroutines in libraries.
	- When you create a goroutine, know how it will end
		- Avoids subtle memory leaks
	- Check for race conditions at compile time
	 - go run -race main.go
*/
