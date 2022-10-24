package db

import "fmt"

func UnknownErr(err error) error {
	return fmt.Errorf("MessageStorage: something go wrong: %v", err)
}

var (
	ErrUnknownMessage = fmt.Errorf("unknown message")
)
