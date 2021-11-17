package main

import (
	"fmt"
	"sync"
	"testing"
)

var routinesAmount, decr int = 10, 10

func rout(wg *sync.WaitGroup, whoami int, decr *int) {
	wg.Add(1)
	defer wg.Done()
	fmt.Println("GoR - ", whoami)
	*decr--
}

func TestWaitGroup(t *testing.T) {
	// func main() {
	var wg sync.WaitGroup
	for i := 0; i < routinesAmount; i++ {
		go rout(&wg, i, &decr)
	}
	wg.Wait()
	if decr != 0 {
		t.Error("Expected 0, got ", decr)
	}
	fmt.Println("Main - Done")
}
