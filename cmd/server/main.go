package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/jubobs/namecheck"
	_ "github.com/jubobs/namecheck/twitter"
)

type (
	result struct {
		SocialNetwork string `json:"social_network"`
		Valid         string `json:"valid"`
		Available     string `json:"available"`
	}
	entity struct {
		Username string   `json:"username"`
		Results  []result `json:"results"`
	}
)

func main() {
	http.HandleFunc("/check", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if len(username) == 0 {
		http.Error(w, "missing 'username' query parameter", http.StatusBadRequest)
		return
	}
	checkers := namecheck.Checkers()
	n := len(checkers)
	ch := make(chan result, n)

	var wg sync.WaitGroup
	wg.Add(n)
	for _, checker := range checkers {
		go func(c namecheck.Checker) {
			defer wg.Done()
			res := result{SocialNetwork: c.String()}
			valid := c.Validate(username)
			res.Valid = strconv.FormatBool(valid)
			if valid {
				available, err := c.Available(username)
				res.Available = availString(available, err)
			}
			ch <- res
		}(checker)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	writeResp(w, username, ch)
}

func writeResp(w http.ResponseWriter, username string, ch <-chan result) error {
	w.Header().Add("Content-Type", "application/json")
	v := entity{Username: username}
	results := make([]result, 0)
	for res := range ch {
		results = append(results, res)
	}
	v.Results = results
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(&v)
}

func availString(available bool, err error) string {
	if err != nil {
		return "unknown"
	}
	return strconv.FormatBool(available)
}
