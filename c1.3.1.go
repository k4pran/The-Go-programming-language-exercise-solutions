package main

// Modify dup2 to print the names of all files in which each duplicated line occurs.

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)


// ====================================================================
// EXAMPLE OUTPUT
// ====================================================================
// 9       We are legion - .../exercises/duplicates.txt
// 4       The fantastic four - .../exercises/more_duplicates.txt
// 6       We are legion - .../exercises/more_duplicates.txt
// 2       This is a duplicate line - .../exercises/duplicates.txt
// 3       This line is a triplet - .../exercises/duplicates.txt
// ____________________________________________________________________


func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	sep := " - "
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		fileName := f.Name()
		counts[strings.Join([]string{input.Text(), fileName}, sep)]++
	}
}
