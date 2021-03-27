package main

import (
	"fmt"
)

type node struct {
	value int
}

func createNode(value int) *node {
	n := node{value}
	return &n
}

func main() {
	n := createNode(4)
	fmt.Println(n, n.value)
}

