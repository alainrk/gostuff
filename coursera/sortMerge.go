package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partSort() {

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

func main() {
	l := promptList()
	fmt.Println(l)
}
