package db

import "fmt"

var (
	ErrDuplicatedLogin = fmt.Errorf("duplicate login")
	ErrUnknown         = fmt.Errorf("something go wrong")
)
