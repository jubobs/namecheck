package twitter

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/jubobs/namecheck"
)

const (
	minLen         = 1
	maxLen         = 15
	illegalPattern = "twitter"
)

var legalRegexp = regexp.MustCompile("^[0-9A-Z_a-z]*$")

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

func (t *Twitter) Available(username string) (bool, error) {
	addr := fmt.Sprintf("https://twitter.com/%s", url.PathEscape(username))
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		// request could not be constructed: programming error!
		panic(err)
	}
	resp, err := namecheck.Client.Do(req)
	if err != nil {
		err1 := namecheck.ErrUnknownAvailability{
			Username:      username,
			SocialNetwork: t.String(),
			Cause:         err,
		}
		return false, &err1
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNotFound, nil
}

func (*Twitter) String() string {
	return "Twitter"
}
