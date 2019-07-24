package main

import (
	"fmt"

	"github.com/jubobs/namecheck/twitter"
)

func main() {
	const username = "jub0bs"

	fmt.Printf("%q is long enough: %t\n", username, twitter.IsLongEnough(username))
	fmt.Printf("%q is short enough: %t\n", username, twitter.IsShortEnough(username))
	fmt.Printf("%q only contains legal characters: %t\n", username, twitter.OnlyContainsLegalChars(username))
	fmt.Printf("%q contains no illegal pattern: %t\n", username, twitter.ContainsNoIllegalPattern(username))

}
