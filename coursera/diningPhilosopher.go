package main

import (
	"fmt"
	"sync"
)

const amount int = 5

var (
	dining    int = 0
	maxDining int = 2
)

var (
	askPermission = make(chan int)
	done          = make(chan int)
)

type Chopstick struct {
	id int
	mu sync.Mutex
}

type Philosopher struct {
	id            int
	dinners       int
	chopLeft      *Chopstick
	chopRight     *Chopstick
	getPermission chan bool
}

func host(philosophers []*Philosopher) {
	release := make(chan bool)
	requests := []int{}
	stuffed := 0

	for {
		select {
		case id := <-askPermission:
			requests = append(requests, id)
		case <-release:
			dining--
		case <-done:
			stuffed++
		}
		if stuffed == amount {
			break
		}
		for dining < maxDining && len(requests) > 0 {
			dining++
			var id int
			id, requests = requests[0], requests[1:]
			philosophers[id].eat()
		}
	}
}

func (p *Philosopher) eat() {
	for i := 0; i < int(maxDining); i++ {
		askPermission <- p.id
		<-p.getPermission
		fmt.Println("starting to eat: ", p.id)
		p.chopLeft.mu.Lock()
		p.chopRight.mu.Lock()
		p.dinners++
		if p.dinners == maxDining {
			done <- p.id
		}
		p.chopLeft.mu.Unlock()
		p.chopRight.mu.Unlock()
		fmt.Println("finishing to eat: ", p.id)
	}
}

func main() {
	chopsticks := make([]*Chopstick, amount)
	philosophers := make([]*Philosopher, amount)

	for i := 0; i < amount; i++ {
		chopsticks[i] = &Chopstick{i, sync.Mutex{}}
		philosophers[i] = &Philosopher{i, 0, chopsticks[i], chopsticks[(i+1)%amount], make(chan bool)}
	}

	go host(philosophers)

	for _, p := range philosophers {
		go p.eat()
	}
}
