package namecheck

import (
	"fmt"
	"sync"
)

var (
	checkersMu sync.RWMutex
	checkers   = make([]Checker, 0)
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
	checkersMu.RLock()
	checkersMu.RUnlock()
	return checkers
}

func Register(c Checker) {
	checkersMu.Lock()
	defer checkersMu.Unlock()
	checkers = append(checkers, c)
}
