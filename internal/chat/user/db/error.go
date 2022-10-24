package db

import "fmt"

func UnknownErr(err error) error {
	return fmt.Errorf("something go wrong: %v", err)
}

var (
	ErrDuplicatedLogin = fmt.Errorf("duplicate login")
	ErrUnknownUser     = fmt.Errorf("unknown user")
)
