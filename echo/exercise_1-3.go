package main

/*
 * EXERCISE 1.3
 * Experiment to measure the difference in running time between our potentially inefficient versions and the one
 * that uses string.join.
*/

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// Time original method
	start_origin := time.Now()
	fmt.Println("Running original method without 'join'...")
	ex_1_original()
	original_performance := time.Since(start_origin)

	// Time new method with 'string.Join' func
	start_new := time.Now()
	fmt.Println("Running new method using 'join'...\n")
	ex_1_join()
	new_performance := time.Since(start_new)

	// Report results
	fmt.Println("------------------------------------------")
	fmt.Println("                 RESULTS")
	fmt.Println("------------------------------------------")
	fmt.Println("\t-Original timing: " + original_performance.String())
	fmt.Println("\t-New performance: " + new_performance.String() + "\n")

	// Summarise findings
	if original_performance < new_performance {
		fmt.Println("--The original method without using join is faster by " +
			(new_performance - original_performance).String() + "--")

	} else if new_performance < original_performance {
		fmt.Println("--The new method using string.Join is faster by " +
			(original_performance - new_performance).String() + "--")

	} else {
		fmt.Println("Both methods were equally fast / slow--")
	}
	fmt.Println()
}

// --ORIGINAL FUNCTION--

func ex_1_original() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// --POTENTIALLY IMPROVED FUNCTION USING 'JOIN'--

func ex_1_join() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s = strings.Join([]string{s, arg}, sep)
		sep = " "
	}
	fmt.Println(s)
}

/*
====================================================================
						EXAMPLE OUTPUT
====================================================================


------------------------------------------
                 RESULTS
------------------------------------------
        -Original timing: 37.706µs
        -New performance: 3.078µs

--The new method using string.Join is faster by 34.628µs--
 */
