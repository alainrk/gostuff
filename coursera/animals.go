package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TAnimal struct {
	food  string
	move  string
	sound string
}

func (a TAnimal) Eat() {
	fmt.Println(a.food)
}

func (a TAnimal) Move() {
	fmt.Println(a.move)
}

func (a TAnimal) Speak() {
	fmt.Println(a.sound)
}

func animals() {
	cow := TAnimal{"grass", "walk", "moo"}
	bird := TAnimal{"worms", "fly", "peep"}
	snake := TAnimal{"mice", "slither", "hsss"}

	animalsMap := map[string]TAnimal{
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
