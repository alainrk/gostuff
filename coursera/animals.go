package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food  string
	move  string
	sound string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.move)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	animalsMap := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	var animalInput, command string
	var chunks []string
	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, _ := in.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		chunks = strings.Split(strings.ToLower(strings.TrimSpace(input)), " ")

		animalInput = chunks[0]
		command = chunks[1]

		animal, ok := animalsMap[animalInput]
		if !ok {
			fmt.Println("This animal does not exist, yet.")
			continue
		}

		switch command {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid command")
		}
	}
}
