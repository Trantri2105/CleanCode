package main

import (
	"fmt"
	"sync"
)

// Define a mutex to ensure thread safety during singleton creation
var lock = sync.Mutex{}

// Define the singleton struct
type single struct {
}

// Define a package-level variable to hold the singleton instance
var singleInstance *single

func getInstance() *single {
	// First check if instance already exists
	if singleInstance == nil {
		// Lock to prevent multiple goroutines from creating instances simultaneously
		lock.Lock()
		defer lock.Unlock()

		// Double-check if instance is still nil after acquiring the lock
		// (another goroutine might have created it while we were waiting for the lock)
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			getInstance()
			wg.Done()
		}()
	}
	wg.Wait()
}
