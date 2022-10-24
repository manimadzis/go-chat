package repository

import "fmt"

func UnknownErr(err error) error {
	return fmt.Errorf("something went wrong: %v", err)
}

var (
	ErrUnknownChat     = fmt.Errorf("unknown chat")
	ErrUnknownMessage  = fmt.Errorf("unknown message")
	ErrDuplicatedLogin = fmt.Errorf("duplicate login")
	ErrUnknownUser     = fmt.Errorf("unknown user")
)
