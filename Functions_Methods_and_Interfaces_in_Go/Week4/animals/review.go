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

/*
Cow
*/

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

/*
Bird
*/

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

/*
Snake
*/

type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	db := make(map[string]Animal, 0)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		op, err := reader.ReadString('\n')
		checkError(err)
		op = strings.Trim(op, "\n")
		opSplited := strings.Split(op, " ")

		if len(opSplited) == 3 {
			switch opSplited[0] {
			case "newanimal":
				switch opSplited[2] {
				case "cow":
					db[opSplited[1]] = Cow{}
				case "bird":
					db[opSplited[1]] = Bird{}
				case "snake":
					db[opSplited[1]] = Snake{}
				}
			case "query":
				if selectedAnimal, ok := db[opSplited[1]]; ok {
					switch opSplited[2] {
					case "eat":
						selectedAnimal.Eat()
					case "move":
						selectedAnimal.Move()
					case "speak":
						selectedAnimal.Speak()
					default:
						fmt.Printf("Invalid information requested %s\n", opSplited[2])
					}
				} else {
					fmt.Printf("Animal %s not found\n", opSplited[1])
				}
			}
		} else {
			fmt.Printf("Invalid number of arguments %d\n", len(opSplited))
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
