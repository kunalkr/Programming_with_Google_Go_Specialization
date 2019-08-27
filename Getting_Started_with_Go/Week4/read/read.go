package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Print("Enter text file name: ")
	fmt.Scanf("%s", &fileName)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Failed while opening file: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	namesSlice := make([]Name, 0)

	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")
		nameObj := Name{name[0], name[1]}
		namesSlice = append(namesSlice, nameObj)
	}

	for _, name := range namesSlice {
		fmt.Println("First: ", name.fname, "\tLast: ", name.lname)
	}
}
