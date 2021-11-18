package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func asyncEcho(c chan<- int, n int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	c <- n
}

func TestChannels2(t *testing.T) {
	for i := 0; i < 10; i++ {
		c1 := make(chan int)
		c2 := make(chan int)
		c3 := make(chan int)
		c4 := make(chan int)

		go asyncEcho(c1, 1)
		go asyncEcho(c2, 2)
		go asyncEcho(c3, 3)
		go asyncEcho(c4, 4)

		r1 := <-c1
		r2 := <-c2
		r3 := 0
		r4 := 0

		select {
		case r3 = <-c3:
			fmt.Println("received on c3")
		case r4 = <-c4:
			fmt.Println("received on c4")
		}

		if r1 != 1 || r2 != 2 {
			t.Error("expected to receive on bot channels 1 and 2")
		}

		if r3 == 0 && r4 == 0 {
			t.Error("expected to receive at least on one of channels 3 and 4")
		}

		if r3 == 3 && r4 == 4 {
			t.Error("expected to receive at most on one of channels 3 and 4")
		}
	}
}
