package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name  string
	food  string
	move  string
	sound string
}

func (t Cow) Eat() {
	fmt.Println(t.food)
}
func (t Cow) Move() {
	fmt.Println(t.move)
}
func (t Cow) Speak() {
	fmt.Println(t.sound)
}
func NewCow(name string) *Cow {
	return &Cow{name, "grass", "walk", "moo"}
}

type Bird struct {
	name  string
	food  string
	move  string
	sound string
}

func (t Bird) Eat() {
	fmt.Println(t.food)
}
func (t Bird) Move() {
	fmt.Println(t.move)
}
func (t Bird) Speak() {
	fmt.Println(t.sound)
}
func NewBird(name string) *Bird {
	return &Bird{name, "worms", "fly", "peep"}
}

type Snake struct {
	name  string
	food  string
	move  string
	sound string
}

func (t Snake) Eat() {
	fmt.Println(t.food)
}
func (t Snake) Move() {
	fmt.Println(t.move)
}
func (t Snake) Speak() {
	fmt.Println(t.sound)
}
func NewSnake(name string) *Snake {
	return &Snake{name, "mice", "slither", "hsss"}
}

func animalsInterface() {
	animals := make(map[string]Animal)
	var chunks []string
	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, _ := in.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		chunks = strings.Split(strings.ToLower(strings.TrimSpace(input)), " ")

		command := chunks[0]

		if command == "newanimal" {
			name := chunks[1]
			animal := chunks[2]

			switch animal {
			case "cow":
				animals[name] = NewCow(name)
			case "bird":
				animals[name] = NewBird(name)
			case "snake":
				animals[name] = NewSnake(name)
			default:
				fmt.Println("Unknown animal")
				continue
			}
			fmt.Println("Created it!")
			continue
		}

		if command == "query" {
			name := chunks[1]
			action := chunks[2]
			animal, ok := animals[name]
			if !ok {
				fmt.Println("Unknown animal")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown action")
			}
		}
	}
}
