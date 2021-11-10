package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortedInsert2Slice(s []int, num int) []int {
	news := append(s, num)
	sort.Ints(news)
	return news
}

func slice() {
	s := make([]int, 0, 3)

	for {
		var input string
		print("Enter an integer to add to the list (X to quit): ")
		_, err := fmt.Scanf("%s", &input)

		if input == "X" {
			break
		}

		if err != nil {
			fmt.Println("Invalid input", err)
			continue
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		s = sortedInsert2Slice(s, num)
		fmt.Println(s)
	}
}
