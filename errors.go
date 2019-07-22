package namecheck

import (
	"fmt"
)

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
