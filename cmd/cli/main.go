package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"text/tabwriter"

	"github.com/jubobs/namecheck"
	_ "github.com/jubobs/namecheck/twitter"
)

type result struct {
	socnet string
	valid  bool
	avail  bool
	err    error
}

func main() {
	run(os.Args, os.Stdout, os.Stderr)
}

func run(args []string, stdout, stderr io.Writer) {
	if len(args) != 2 {
		fmt.Fprintln(stderr, "Usage:", args[0], "<username>")
		return
	}
	var username = os.Args[1]

	var wg sync.WaitGroup
	checkers := namecheck.Checkers()
	n := len(checkers)
	results := make(chan result, n)

	wg.Add(n)
	for _, checker := range checkers {
		go func(c namecheck.Checker) {
			defer wg.Done()
			res := result{socnet: c.String()}
			res.valid = c.Validate(username)
			if res.valid {
				dispo, err := c.Available(username)
				res.avail = err == nil && dispo
				res.err = err
			}
			results <- res
		}(checker)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// store the results drained from the channel in a slice
	rs := make([]result, 0, n)
	for res := range results {
		rs = append(rs, res)
	}
	// sort the slice in lexicographic order of social-socnetork names
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].socnet < rs[j].socnet
	})

	prettyPrint(rs, stdout)
}

func prettyPrint(rs []result, stdout io.Writer) {
	const padding = 3
	const tmpl = "%s\t%s\t%s\t\n"
	w := tabwriter.NewWriter(stdout, 0, 0, padding, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "social_socnetork\tvalid\tavailable\t")
	for _, r := range rs {
		fmt.Fprintf(w, tmpl, r.socnet, validString(r), availString(r))
	}
	w.Flush()
}

func availString(r result) string {
	switch {
	case r.err != nil:
		return "?"
	case r.avail:
		return "\u2714"
	default:
		return "\u2717"
	}
}

func validString(r result) string {
	if r.valid {
		return "\u2714"
	} else {
		return "\u2717"
	}
}
