package twitter_test

import (
	"testing"

	"github.com/jubobs/namecheck/twitter"
)

func TestIsLongEnoughFailsOnNamesShorterThan1Chars(t *testing.T) {
	username := ""
	want := false
	got := twitter.IsLongEnough(username)
	if got != want {
		t.Errorf("twitter.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestIsLongEnoughSucceedsOnNamesLongerThan0Chars(t *testing.T) {
	username := "a"
	want := true
	got := twitter.IsLongEnough(username)
	if got != want {
		t.Errorf("twitter.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestIsShortEnoughFailsOnNamesLongerThan15Chars(t *testing.T) {
	username := "obviously_longer_than_15"
	want := false
	got := twitter.IsShortEnough(username)
	if got != want {
		t.Errorf("twitter.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}

func TestIsShortEnoughSucceedsOnNamesShorterThan16Chars(t *testing.T) {
	username := "fifteen_exactly"
	want := true
	got := twitter.IsShortEnough(username)
	if got != want {
		t.Errorf("twitter.IsLongEnough(%s) = %t; want %t", username, got, want)
	}
}
