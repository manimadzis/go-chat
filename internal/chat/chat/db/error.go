package db

import "fmt"

func UnknownErr(err error) error {
	return fmt.Errorf("ChatStorage: something went wrong: %v", err)
}

var (
	ErrUnknownChat = fmt.Errorf("unknown chat")
)
