package namecheck

import (
	"fmt"

	"github.com/jubobs/namecheck/twitter"
)

type Validator interface {
	Validate(username string) bool
}

type Availabler interface {
	Available(username string) (bool, error)
}

type Checker interface {
	fmt.Stringer
	Validator
	Availabler
}

func Checkers() []Checker {
	// for this exercise, let's pretend we support 20 different social networks
	n := 20
	s := make([]Checker, 0, n)
	for i := 0; i < n; i++ {
		var c twitter.Twitter
		s = append(s, &c)
	}
	return s
}

type ErrUnknownAvailability struct {
	Username      string
	SocialNetwork string
	Cause         error
}

func (e *ErrUnknownAvailability) Error() string {
	const msg = "couldn't check availability of %q on %s"
	return fmt.Sprintf(msg, e.Username, e.SocialNetwork)
}

func (e *ErrUnknownAvailability) Unwrap() error {
	return e.Cause
}
