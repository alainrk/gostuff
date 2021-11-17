package main

import (
	"fmt"
	"sync"
	"testing"
)

func prod(wg *sync.WaitGroup, whoami int, s []int, c chan<- int) {
	wg.Add(1)
	defer wg.Done()

	res := 1
	for _, v := range s {
		res *= v
	}
	fmt.Printf("Partial from gor[%d] = %d\n", whoami, res)
	c <- res
}

func TestChannels1(t *testing.T) {
	c := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		go prod(&wg, i, []int{1, 2, 3}, c)
	}

	res := 1
	for i := 0; i < 3; i++ {
		x := <-c
		if x != 6 {
			t.Errorf("Expected 6 partial result, got %d", x)
		}
		res *= x
	}
	close(c)
	wg.Wait()

	if res != 216 {
		t.Errorf("Expected 216 total result, got %d", res)
	}
}
