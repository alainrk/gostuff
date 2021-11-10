package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func makeJson() {
	people := make(map[string]string)

	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := in.ReadString('\n')
	fmt.Print("Enter address: ")
	address, _ := in.ReadString('\n')

	people["name"] = strings.TrimSuffix(name, "\n")
	people["address"] = strings.TrimSuffix(address, "\n")

	jsonb, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonb))
}
