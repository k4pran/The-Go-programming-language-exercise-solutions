package main

// Exercise 1.2.1 - Modify the echo program to also print os.Args[0], the name of the command that invoked it.

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
