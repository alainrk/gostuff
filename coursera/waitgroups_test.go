package main

import (
	"fmt"
	"sync"
	"testing"
)

var routinesAmount int = 10

var (
	mu   sync.Mutex
	decr int = routinesAmount
)

func rout(wg *sync.WaitGroup, whoami int) {
	wg.Add(1)
	defer wg.Done()

	mu.Lock()
	decr--
	fmt.Printf("GoR [%d], decr = %d\n", whoami, decr)
	mu.Unlock()
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < routinesAmount; i++ {
		go rout(&wg, i)
	}

	wg.Wait()

	if decr != 0 {
		t.Error("Expected 0, got ", decr)
	}
	fmt.Println("Main - Done")
}
