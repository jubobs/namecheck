package main

import (
	"fmt"
	"log"
	"regexp"
	"unicode/utf8"
)

const minLen = 1
const maxLen = 15
const legalPattern = "^[0-9A-Z_a-z]*$"
const illegalPattern = "(?i)twitter"

func main() {
	const username = "jub0bs"

	fmt.Printf("%q is long enough: %t\n", username, isLongEnough(username))

	fmt.Printf("%q is short enough: %t\n", username, isShortEnough(username))

	ok, err := onlyContainsLegalChars(username)
	if err != nil {
		log.Fatal("something wrong happened", err)
	}
	fmt.Printf("%q only contains legal characters: %t\n", username, ok)

	ok, err = containsNoIllegalPattern(username)
	if err != nil {
		log.Fatal("something wrong happened", err)
	}
	fmt.Printf("%q contains no illegal pattern: %t\n", username, ok)

}

func isLongEnough(username string) bool {
	return minLen <= utf8.RuneCountInString(username)
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func onlyContainsLegalChars(username string) (bool, error) {
	matched, err := regexp.MatchString(legalPattern, username)
	return matched, err
}

func containsNoIllegalPattern(username string) (bool, error) {
	matched, err := regexp.MatchString(illegalPattern, username)
	return !matched, err
}
