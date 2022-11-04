package domain

type DTO interface {
	Valid() error
}
