package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxNameLength = 20

type name struct {
	fname string
	lname string
}

func truncate(s string, m int) string {
	if len(s) > m {
		return s[:m]
	}
	return s
}

func newName(fname string, lname string) name {
	fname = truncate(fname, maxNameLength)
	lname = truncate(lname, maxNameLength)
	n := name{fname, lname}
	return n
}

func read() {
	var path string
	fmt.Print("Insert file path: ")
	fmt.Scanf("%s", &path)

	people := make([]name, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		uname := strings.Split(scanner.Text(), " ")
		name := newName(uname[0], uname[1])
		people = append(people, name)
	}

	for _, p := range people {
		fmt.Println(p.fname, p.lname)
	}
}
