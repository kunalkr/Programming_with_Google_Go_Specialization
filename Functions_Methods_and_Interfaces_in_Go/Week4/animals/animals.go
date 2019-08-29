package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal : Interface for Animal type
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow : Struct for cow type animal
type Cow struct {
	food       string
	locomotion string
	sound      string
}

// Bird : Struct for bird type animal
type Bird struct {
	food       string
	locomotion string
	sound      string
}

// Snake : Struct for snake type animal
type Snake struct {
	food       string
	locomotion string
	sound      string
}

// Eat : Function to print what a cow type animal eats
func (c *Cow) Eat() {
	fmt.Println(c.food)
}

// Move : Function to print how a cow type animal moves
func (c *Cow) Move() {
	fmt.Println(c.locomotion)
}

// Speak : Function to print what a cow type animal speaks
func (c *Cow) Speak() {
	fmt.Println(c.sound)
}

// Eat : Function to print what a bird type animal eats
func (b *Bird) Eat() {
	fmt.Println(b.food)
}

// Move : Function to print how a bird type animal moves
func (b *Bird) Move() {
	fmt.Println(b.locomotion)
}

// Speak : Function to print what a bird type animal speaks
func (b *Bird) Speak() {
	fmt.Println(b.sound)
}

// Eat : Function to print what a cow snake animal eats
func (s *Snake) Eat() {
	fmt.Println(s.food)
}

// Move : Function to print how a cow snake animal moves
func (s *Snake) Move() {
	fmt.Println(s.locomotion)
}

// Speak : Function to print what a cow snake animal speaks
func (s *Snake) Speak() {
	fmt.Println(s.sound)
}

func main() {
	animalList := map[string]Animal{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		input := strings.Trim(scanner.Text(), " \r\n")
		splitInput := strings.Split(input, " ")
		command, param1, param2 := splitInput[0], splitInput[1], splitInput[2]

		switch command {
		case "newanimal":
			switch strings.ToLower(param2) {
			case "cow":
				animalList[param1] = &Cow{food: "grass", locomotion: "walk", sound: "moo"}
			case "bird":
				animalList[param1] = &Bird{food: "worms", locomotion: "fly", sound: "peep"}
			case "snake":
				animalList[param1] = &Snake{food: "mice", locomotion: "slither", sound: "hsss"}
			}
			fmt.Println("Created it!")
		case "query":
			switch strings.ToLower(param2) {
			case "eat":
				animalList[param1].Eat()
			case "move":
				animalList[param1].Move()
			case "speak":
				animalList[param1].Speak()
			}
		}
	}
}
