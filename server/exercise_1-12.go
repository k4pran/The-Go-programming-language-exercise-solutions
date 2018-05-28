package main

/*
 * EXERCISE_1.12
 * Modify the Lissajous server to read parameter values from the URL. For example,
 * you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets
 * the number of cycles to 20 instead of the default 5.
*/


import (
	"log"
	"net/http"
	"fmt"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", handlerLissajous)
	http.HandleFunc("/lissajous_rainbow", handlerLissajousRainbow)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerLissajous(w http.ResponseWriter, r *http.Request) {
	result := r.URL.Query().Get("cycles")
	cyclesFloat, err := strconv.ParseFloat(result, 64)
	if err == nil {
		lissajous(w, cyclesFloat)
	} else {
		lissajous(w, 5)
	}
}

func handlerLissajousRainbow(w http.ResponseWriter, r *http.Request) {
	result := r.URL.Query().Get("cycles")
	cyclesFloat, err := strconv.ParseFloat(result, 64)
	if err == nil {
		lissajousRainbow(w, cyclesFloat)
	} else {
		lissajousRainbow(w, 5)
	}
}

	func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err !=nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}
}

/*
====================================================================
						EXAMPLE OUTPUT
====================================================================

1. Run both go files in server directory - go run exercise_1-12.go lissajous.go
2. Go to localhost:8000 on your browser to view headers.
3. Go to http://localhost:8000/lissajous to view the basic gif
4. Go to http://localhost:8000/lissajous_rainbow to view the rainbow gif
5. Complete step 3 or 4 and add ?cycles=n - replacing 'n' with a number of your choice.
E.g http://localhost:8000/lissajous?cycles=30 to alter the number of cycles and see the results!
 */