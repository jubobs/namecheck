package twitter

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/jubobs/namecheck"
)

const (
	minLen         = 1
	maxLen         = 15
	legalPattern   = "^[0-9A-Z_a-z]*$"
	illegalPattern = "twitter"
)

var (
	legalRegexp                  = regexp.MustCompile(legalPattern)
	client      namecheck.Client = http.DefaultClient
)

type Twitter struct{}

func init() {
	namecheck.Register(&Twitter{})
}

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
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNotFound, nil
}

func (*Twitter) String() string {
	return "Twitter"
}
