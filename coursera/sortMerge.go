package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const threadsNum = 4

func partSort(wg *sync.WaitGroup, l []int, c chan<- []int) {
	defer wg.Done()
	sort.Ints(l)
	fmt.Println("Partial: ", l)
	c <- l
}

func promptList() []int {
	var chunks []string
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a list of integers single-space separated:\n")
	input, _ := in.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	chunks = strings.Split(strings.TrimSpace(input), " ")

	l := make([]int, len(chunks))
	for i, v := range chunks {
		l[i], _ = strconv.Atoi(v)
	}
	return l
}

func monitor(wg *sync.WaitGroup, c chan []int) {
	wg.Wait()
	close(c)
}

func merge(chunks [][]int) []int {
	sorted := chunks[0]
	for _, chunk := range chunks[1:] {
		newSorted := []int{}
		i, j := 0, 0
		for i < len(sorted) && j < len(chunk) {
			if sorted[i] < chunk[j] {
				newSorted = append(newSorted, sorted[i])
				i++
			} else {
				newSorted = append(newSorted, chunk[j])
				j++
			}
		}

		for i < len(sorted) {
			newSorted = append(newSorted, sorted[i])
			i++
		}
		for j < len(chunk) {
			newSorted = append(newSorted, chunk[j])
			j++
		}
		sorted = newSorted
	}
	return sorted
}

func sortMerge() {
	l := promptList()
	chunkSize := len(l) / threadsNum

	// c := make(chan []int, threadsNum)
	var wg sync.WaitGroup
	c := make(chan []int)
	chunks := [][]int{}

	for i := 0; i < threadsNum; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == threadsNum-1 {
			end = len(l)
		}

		wg.Add(1)
		go partSort(&wg, l[start:end], c)
	}

	go monitor(&wg, c)

	for l := range c {
		chunks = append(chunks, l)
	}

	sorted := merge(chunks)
	fmt.Println(sorted)
}
