package main

/*
 * EXERCISE 1.2
 * Modify the echo program to print the index and value of each of its arguments, one per line.
 */

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

/*
====================================================================
						EXAMPLE OUTPUT
====================================================================

0. /var/folders/18/36hnd7px0bj0sz14rvrhmbrc0000gn/T/go-build445586410/b001/exe/exercise_1-21. Hello2. World!
 */