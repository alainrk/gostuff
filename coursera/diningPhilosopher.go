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

func host(wg *sync.WaitGroup, philosophers []*Philosopher) {
	defer wg.Done()

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
			philosophers[id].getPermission <- true
		}
	}
}

func (p *Philosopher) run() {
	fmt.Println("Running philo: ", p.id)
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
	var wg sync.WaitGroup
	chopsticks := make([]*Chopstick, amount)
	philosophers := make([]*Philosopher, amount)

	for i := 0; i < amount; i++ {
		chopsticks[i] = &Chopstick{i, sync.Mutex{}}
		philosophers[i] = &Philosopher{i, 0, chopsticks[i], chopsticks[(i+1)%amount], make(chan bool)}
	}

	wg.Add(1)
	go host(&wg, philosophers)

	for _, p := range philosophers {
		go p.run()
	}

	wg.Wait()
}
