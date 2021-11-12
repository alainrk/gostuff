package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(l []int, i int) {
	j := i + 1
	if j >= len(l) {
		return
	}
	l[i], l[j] = l[j], l[i]
}

func BubbleSort(l []int) {
	for {
		swapped := false

		for i := 0; i < len(l)-1; i++ {
			if l[i] > l[i+1] {
				Swap(l, i)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func printResult(l []int) {
	for _, x := range l {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}

func getCleanSliceNumStringsFromString(s string) []string {
	l := strings.TrimSuffix(s, "\n")
	l = strings.Trim(l, " ")
	var slice []string
	for _, v := range strings.Split(l, " ") {
		if v != "" {
			slice = append(slice, v)
		}
	}
	return slice
}

func getSliceFromUser(maxn uint8) []int {
	in := bufio.NewReader(os.Stdin)

	var numSliceStr []string
	for {
		fmt.Print("Enter numbers divided by one space (up to 10 numbers): ")
		numStr, _ := in.ReadString('\n')
		numSliceStr = getCleanSliceNumStringsFromString(numStr)
		if len(numSliceStr) > int(maxn) {
			fmt.Println("Too many numbers.")
			continue
		}
		break
	}

	if len(numSliceStr) < int(maxn) {
		maxn = uint8(len(numSliceStr))
	}

	l := make([]int, maxn)
	i := uint8(0)
	for _, s := range numSliceStr {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		l[i] = x
		i++

		if i >= maxn {
			break
		}
	}
	return l
}

func main() {
	l := getSliceFromUser(10)
	BubbleSort(l)
	printResult(l)
}
