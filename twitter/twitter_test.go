package twitter

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jubobs/namecheck"
	"github.com/jubobs/namecheck/mock"
)

var (
	checker    Twitter
	dummyError = errors.New("Oh no!")
)

func TestIsLongEnoughFailsOnNamesShorterThan1Chars(t *testing.T) {
	username := ""
	want := false
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestIsLongEnoughSucceedsOnNamesLongerThan0Chars(t *testing.T) {
	username := "a"
	want := true
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestIsShortEnoughFailsOnNamesLongerThan15Chars(t *testing.T) {
	username := "obviously_longer_than_15"
	want := false
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestIsShortEnoughSucceedsOnNamesShorterThan16Chars(t *testing.T) {
	username := "fifteen_exactly"
	want := true
	got := checker.Validate(username)
	if got != want {
		t.Errorf("t.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestCheck(t *testing.T) {
	username := "dummy"
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

	const template = "Check(%q), got %t, want %t"
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			namecheck.Client = c.client // overwrite the client for this test case
			actual, err := checker.Available(username)
			if actual != c.want || (err == nil) != (c.err == nil) {
				t.Errorf(template, username, actual, c.want)
			}
		})
	}
}
