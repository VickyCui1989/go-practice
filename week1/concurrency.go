package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple function to run as a goroutine
func say(name string, wg *sync.WaitGroup) {
	defer wg.Done() // mark as done when function finishes
	for i := 0; i < 3; i++ {
		fmt.Printf("Hi from %s (%d)\n", name, i)
		time.Sleep(200 * time.Millisecond)
	}
}

// Demonstrates goroutines and WaitGroup
func GoroutineDemo() {
	var wg sync.WaitGroup
	wg.Add(2) // we will launch 2 goroutines

	go say("Goroutine A", &wg)
	go say("Goroutine B", &wg)

	wg.Wait() // wait for both to finish
	fmt.Println("Both goroutines finished.")
}

// Demonstrates channel communication
func ChannelDemo() {
	ch := make(chan string)

	// Sender goroutine
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "Hello from channel"
	}()

	// Receiver in main goroutine
	msg := <-ch
	fmt.Println("Received:", msg)
}

// Entry point for this file
func ConcurrencyDemo() {
	fmt.Println("=== Goroutine Demo ===")
	GoroutineDemo()

	fmt.Println("\n=== Channel Demo ===")
	ChannelDemo()
}
