package github

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
	minLen           = 1
	maxLen           = 39
	illegalPrefix    = "-"
	illegalSuffix    = "-"
	illegalSubstring = "--"
)

var legalRegexp = regexp.MustCompile("^[-0-9A-Za-z]*$")

type GitHub struct{}

func init() {
	namecheck.Register(&GitHub{})
}

func (*GitHub) Validate(username string) bool {
	return isLongEnough(username) &&
		isShortEnough(username) &&
		onlyContainsLegalChars(username) &&
		containsNoIllegalPrefix(username) &&
		containsNoIllegalSuffix(username) &&
		containsNoIllegalSubstring(username)
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

func containsNoIllegalSubstring(username string) bool {
	return !strings.Contains(username, illegalSubstring)
}

func containsNoIllegalPrefix(username string) bool {
	return !strings.HasPrefix(username, illegalPrefix)
}

func containsNoIllegalSuffix(username string) bool {
	return !strings.HasSuffix(username, illegalSuffix)
}

func (t *GitHub) Available(username string) (bool, error) {
	addr := fmt.Sprintf("https://github.com/%s", url.PathEscape(username))
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

func (*GitHub) String() string {
	return "GitHub"
}
