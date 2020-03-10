package twitter

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/jubobs/namecheck"
	"github.com/jubobs/namecheck/mock"
)

var (
	checker    Twitter
	dummyError = errors.New("Oh no!")
)

func TestValidateFailsOnNamesThatContainIllegalChars(t *testing.T) {
	username := "hyphen-"
	want := false
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.Validate(%s) = %t; want %t", username, got, want)
	}
}

func TestValidateFailsOnNamesThatContainIllegalPattern(t *testing.T) {
	username := "fooTwItterbar"
	want := false
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.Validate(%s) = %t; want %t", username, got, want)
	}
}

func TestValidateFailsOnNamesThatAreTooShort(t *testing.T) {
	username := ""
	want := false
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.Validate(%s) = %t; want %t", username, got, want)
	}
}

func TestValidateSucceedsOnNamesThatAreLongEnough(t *testing.T) {
	username := "a"
	want := true
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.Validate(%s) = %t; want %t", username, got, want)
	}
}

func TestValidateFailsOnNamesThatAreTooLong(t *testing.T) {
	username := strings.Repeat("a", maxLen+1)
	want := false
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.Validate(%s) = %t; want %t", username, got, want)
	}
}

func TestValidateSucceedsOnNamesThatAreShortEnough(t *testing.T) {
	username := strings.Repeat("a", maxLen)
	want := true
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.Validate(%s) = %t; want %t", username, got, want)
	}
}

func TestAvailable(t *testing.T) {
	username := "dummy"
	// table-driven tests for this one...
	cases := []struct {
		label  string
		client namecheck.HTTPClient
		want   bool
		err    error
	}{
		{
			label:  "notfound",
			client: mock.ClientWithStatusCode(http.StatusNotFound),
			want:   true,
			err:    nil,
		}, {
			label:  "ok",
			client: mock.ClientWithStatusCode(http.StatusOK),
			want:   false,
			err:    nil,
		}, {
			label:  "lowlevelerror",
			client: mock.ClientWithError(dummyError),
			want:   false,
			err: &namecheck.ErrUnknownAvailability{
				Username:      username,
				SocialNetwork: checker.String(),
				Cause:         dummyError,
			},
		},
	}

	const template = "t.Available(%q), got %t, want %t"
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			namecheck.Client = c.client // use mocked client!
			actual, err := checker.Available(username)
			if actual != c.want || (err == nil) != (c.err == nil) {
				t.Errorf(template, username, actual, c.want)
			}
		})
	}
}
