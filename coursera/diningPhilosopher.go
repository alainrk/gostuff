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

func (p *Philosopher) run(wg *sync.WaitGroup) {
	fmt.Println("running philo:", p.id)
	for i := 0; i < int(maxDining); i++ {
		askPermission <- p.id
		fmt.Println("waiting for permission:", p.id)
		<-p.getPermission
		fmt.Println("starting to eat:", p.id)
		p.chopLeft.mu.Lock()
		p.chopRight.mu.Lock()
		fmt.Println("both chopstic locked:", p.id)
		p.dinners++
		if p.dinners == maxDining {
			done <- p.id
		}
		p.chopLeft.mu.Unlock()
		p.chopRight.mu.Unlock()
		fmt.Println("finishing to eat:", p.id)
		fmt.Println(p.id, "ate", p.dinners, "times")
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	chopsticks := make([]*Chopstick, amount)
	philosophers := make([]*Philosopher, amount)

	for i := 0; i < amount; i++ {
		chopsticks[i] = &Chopstick{i, sync.Mutex{}}
	}
	for i := 0; i < amount; i++ {
		philosophers[i] = &Philosopher{i, 0, chopsticks[i], chopsticks[(i+1)%amount], make(chan bool)}
	}

	wg.Add(1)
	go host(&wg, philosophers)

	for _, p := range philosophers {
		wg.Add(1)
		go p.run(&wg)
	}

	wg.Wait()
}
