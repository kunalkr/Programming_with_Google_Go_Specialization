package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}
func (a *Animal) Move() string {
	return a.locomotion
}
func (a *Animal) Speak() string {
	return a.noise
}

func Cow() Animal {
	return Animal{"grass", "walk", "moo"}
}

func Bird() Animal {
	return Animal{"worms", "fly", "peep"}
}

func Snake() Animal {
	return Animal{"mice", "slither", "hsss"}
}

func printUsage(err error) {
	fmt.Println("Wrong input, enter something of the form '<cow|bird|snake> <eat|move|speak>' --", err)
}

func parseAnimal(s string) (*Animal, error) {
	switch s {
	case "cow":
		a := Cow()
		return &a, nil
	case "snake":
		a := Snake()
		return &a, nil
	case "bird":
		a := Bird()
		return &a, nil
	default:
		return nil, fmt.Errorf("invalid animal")
	}
}

func parseCommand(a *Animal, s string) (string, error) {
	switch s {
	case "eat":
		return a.Eat(), nil
	case "move":
		return a.Move(), nil
	case "speak":
		return a.Speak(), nil
	default:
		return "", fmt.Errorf("invalid command")
	}
}
func main() {
	for {
		var animalInput, commandInput string

		// Prompt the user for input
		fmt.Print(">")
		_, err := fmt.Scan(&animalInput, &commandInput)
		if err != nil {
			fmt.Println("Something went wrong, please try again --", err)
			continue
		}

		// Parse animal
		a, err := parseAnimal(animalInput)
		if err != nil {
			printUsage(err)
			continue
		}

		// Parse the command
		output, err := parseCommand(a, commandInput)
		if err != nil {
			printUsage(err)
			continue
		}

		// Print the requesed output
		fmt.Println(output)
	}

}
