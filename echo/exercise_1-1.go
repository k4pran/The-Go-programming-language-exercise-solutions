package main

/*
 * EXERCISE 1.2.1
 * Modify the echo program to also print os.Args[0], the name of the command that invoked it.
 */

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


/*
====================================================================
						EXAMPLE OUTPUT
====================================================================

/var/folders/18/36hnd7px0bj0sz14rvrhmbrc0000gn/T/go-build793967503/b001/exe/exercise_1-1 Hello World!
 */