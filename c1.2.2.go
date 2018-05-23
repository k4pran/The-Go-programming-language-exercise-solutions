package main

// Exercise 1.2.1 - Modify the echo program to print the index and value of each of its arguments, one per line.

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for ind, arg := range os.Args[0:] {
		s = arg
		fmt.Printf("%d. %s", ind, s)
	}
}
