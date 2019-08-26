package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	fmt.Printf("Input string: ")

	stdinScanner := bufio.NewReader(os.Stdin)
	input, _ := stdinScanner.ReadString('\n')
	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n\n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
