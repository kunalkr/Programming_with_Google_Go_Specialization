package main

import (
	"fmt"
	"strings"
)

// Animal : Struct for animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat : Returns eating method of animal
func (a Animal) Eat() string {
	return a.food
}

// Move : Returns moving method of animal
func (a Animal) Move() string {
	return a.locomotion
}

// Speak : Returns speaking method of animal
func (a Animal) Speak() string {
	return a.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	var a string
	var info string
	var result string

	for {
		fmt.Print(">")
		fmt.Scan(&a)
		fmt.Scan(&info)

		switch strings.ToLower(a) {
		case "cow":
			switch strings.ToLower(info) {
			case "eat":
				result = cow.Eat()
			case "move":
				result = cow.Move()
			case "speak":
				result = cow.Speak()
			}
		case "bird":
			switch strings.ToLower(info) {
			case "eat":
				result = bird.Eat()
			case "move":
				result = bird.Move()
			case "speak":
				result = bird.Speak()
			}
		case "snake":
			switch strings.ToLower(info) {
			case "eat":
				result = snake.Eat()
			case "move":
				result = snake.Move()
			case "speak":
				result = snake.Speak()
			}
		}
		fmt.Println(result)
	}
}
