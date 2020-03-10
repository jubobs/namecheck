package twitter

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	minLen         = 1
	maxLen         = 15
	legalPattern   = "^[0-9A-Z_a-z]*$"
	illegalPattern = "twitter"
)

var legalRegexp = regexp.MustCompile(legalPattern)

type Twitter struct{}

func (*Twitter) Validate(username string) bool {
	return isLongEnough(username) &&
		isShortEnough(username) &&
		onlyContainsLegalChars(username) &&
		containsNoIllegalPattern(username)
}

func isLongEnough(username string) bool {
	return minLen <= utf8.RuneCountInString(username)
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func onlyContainsLegalChars(username string) bool {
	return legalRegexp.MatchString(username)
}

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(strings.ToLower(username), illegalPattern)
}

func (*Twitter) Available(username string) (bool, error) {
	address := "https://twitter.com/" + username
	resp, err := http.Get(address)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNotFound, nil
}

func (*Twitter) String() string {
	return "Twitter"
}
