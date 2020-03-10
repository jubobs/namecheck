package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/jubobs/namecheck"
)

func main() {
	http.HandleFunc("/check", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	var wg sync.WaitGroup

	checkers := namecheck.Checkers()
	n := len(checkers)
	ch := make(chan bool, 0)

	wg.Add(n)
	for _, checker := range checkers {
		go func(c namecheck.Checker) {
			defer wg.Done()
			dispo, _ := c.Available(username)
			ch <- dispo
		}(checker)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Fprintf(w, "%t\n", result)
	}
}
