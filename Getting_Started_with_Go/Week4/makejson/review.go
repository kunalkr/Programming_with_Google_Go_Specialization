package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := make(map[string]string, 2)

	var input string
	fmt.Print("Please enter your first name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	data["name"] = strings.Trim(input, "\r\n")

	fmt.Print("Please enter your address: ")
	input, _ = reader.ReadString('\n')
	data["address"] = strings.Trim(input, "\r\n")

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("There was a error while marshalling: %s", err.Error())
	}

	fmt.Printf("JSON data:\n%s", jsonData)
}
