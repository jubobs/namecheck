package twitter

import (
	"regexp"
	"unicode/utf8"
)

const (
	minLen         = 1
	maxLen         = 15
	legalPattern   = "^[0-9A-Z_a-z]*$"
	illegalPattern = "(?i)twitter"
)

var (
	legalRegexp   = regexp.MustCompile(legalPattern)
	illegalRegexp = regexp.MustCompile(illegalPattern)
)

func IsLongEnough(username string) bool {
	return minLen <= utf8.RuneCountInString(username)
}

func IsShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func OnlyContainsLegalChars(username string) bool {
	return legalRegexp.MatchString(username)
}

func ContainsNoIllegalPattern(username string) bool {
	return !illegalRegexp.MatchString(username)
}
