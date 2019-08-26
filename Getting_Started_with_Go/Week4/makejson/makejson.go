package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := scanner.ReadString('\n')
	name = strings.Trim(name, " \n")
	fmt.Print("Enter address: ")
	address, _ := scanner.ReadString('\n')
	address = strings.Trim(address, " \n")

	// fmt.Println(name)
	// fmt.Println(address)

	userMap := make(map[string]string)
	userMap["name"] = name
	userMap["address"] = address

	// fmt.Println(userMap)

	jMap, _ := json.Marshal(userMap)
	os.Stdout.Write(jMap)
	fmt.Println()
}
