package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

// var wg = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel. use struct{} to save memory.

func main() {
	// Channels syncs data flow between goroutines.
	// You can send values into channels from one goroutine and receive those values into another goroutine.
	// Dedicate goroutines to either send or receive values from a channel!

	// ch := make(chan int) // create a channel of type int -> strongly typed.
	// wg.Add(2)
	// go func() { // receiving goroutine.
	// 	i := <-ch // receive value from channel.
	// 	fmt.Println("Received:", i)
	// 	wg.Done()
	// }()
	// go func() { // sending goroutine.
	// 	i := 42
	// 	ch <- i // send value into channel. Arrow depicts the direction of the data flow. Passes copy of i.
	// 	i = 27
	// 	wg.Done()
	// }()
	// wg.Wait()

	// ch := make(chan int)
	// for i := 0; i < 5; i++ { // spawn 10 goroutines. 5 will be receiving, 5 will be sending.
	// 	wg.Add(2)
	// 	go func() { // receiving goroutine.
	// 		i := <-ch
	// 		fmt.Println("Received:", i)
	// 		wg.Done()
	// 	}()
	// 	go func() { // sending goroutine.
	// 		ch <- 42 // blocks until receiving goroutine receives value.
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// ch := make(chan int) // unbuffered bi-directional channel.
	// wg.Add(2)
	// go func(ch <-chan int) { // receive only goroutine.
	// 	i := <-ch
	// 	fmt.Println("Received:", i)
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) { // send only goroutine.
	// 	ch <- 42
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// ch := make(chan int, 50) // buffered bi-directional channel. 50 is the buffer size i.e. 50 int values
	// wg.Add(2)
	// go func(ch <-chan int) {
	// 	for i := range ch {
	// 		fmt.Println("Received:", i)
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) {
	// 	ch <- 42
	// 	ch <- 27
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// ch := make(chan int, 50)
	// wg.Add(2)
	// go func(ch <-chan int) {
	// 	for {
	// 		if i, ok := <-ch; ok { // check if channel is closed.
	// 			fmt.Println("Received:", i)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) {
	// 	ch <- 42
	// 	ch <- 27
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// Logger channel.
	//go logger()
	// logCh <- logEntry{time.Now(), logInfo, "App is starting up..."}
	// logCh <- logEntry{time.Now(), logInfo, "App is shutting down..."}
	// time.Sleep(100 * time.Millisecond)

	// Closing the logger channel using select statements
	go logger2()
	logCh <- logEntry{time.Now(), logInfo, "App is starting up..."}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down..."}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}

}

// func logger() { // monitor channel for incoming log entries.
// 	for entry := range logCh {
// 		fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02 15:04:05"), entry.severity, entry.message)
// 	}
// }

func logger2() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02 15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
