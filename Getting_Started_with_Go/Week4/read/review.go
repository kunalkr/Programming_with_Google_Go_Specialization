package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type names struct {
	first string
	last string
}

func main() {
	var fn string

	fmt.Println("NOTE: please input `.txt` file, which is in same directory with this code!")
	fmt.Print("Enter a file name to read -> ")
	fmt.Scan(&fn)

	fd, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fd.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	nList := make([]names, 0)

	r := bufio.NewScanner(fd)

	for r.Scan() {
		n := names{}
		n.first = strings.Split(r.Text(), " ")[0]
		n.last = strings.Split(r.Text(), " ")[1]
		nList = append(nList, n)
	}

	for _, n := range nList {
		fmt.Println(n.first, n.last)
	}
}