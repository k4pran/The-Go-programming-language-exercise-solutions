package main

/*
 * EXERCISE 1.10
 * Find a web site that produces a large amount of data. Investigate caching by running
 * fetchall twice in succession to see whether the reported time changes much. Do you get
 * the same content each time. Modify fetchall to print its output to a file so it
 * can be examined.
 *
 * EXERCISE 1.11
 * Try fetchall with longer arguments lists, such as samples from the top million web sites available at alexa.com.
 * How does the program behave if a web site just doesn't respond?
 */

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"bufio"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Create("exercise_1-10_output.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	start := time.Now()
	ch := make(chan string)
	_, err = fmt.Fprintf(w, "Fetching urls...\n")
	check(err)

	for _, url := range os.Args[1:] {
		go fetchAll(url, ch)
	}

	var i int
	for range os.Args[1:] {
		fmt.Fprintf(w, "%d. %s\n", i, <-ch)
		i++
	}

	_, err = fmt.Fprintf(w, "%.2fs elapsed from program began\n", time.Since(start).Seconds())
	check(err)
	w.Flush()
}

func fetchAll(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*
====================================================================
						EXAMPLE OUTPUT
====================================================================

Fetching urls...
0. 0.85s  278922 http://www.ebay.com
1. 0.91s  272446 http://www.ebay.com
0.91s elapsed from program began

If a web site does not respond golang golang will not timeout. It defaults to 0
which means to wait indefinitely for a response.
 */