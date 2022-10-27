package passwordhash

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
)

type PasswordHash interface {
	Hash(password string) string
	Valid(password, passwordHash string) error
}

type sha1SaltHash struct {
	salt     string
	sha1Hash hash.Hash
}

func (s sha1SaltHash) Hash(password string) string {
	return s.hash(password)
}

func (s sha1SaltHash) hash(password string) string {
	saltedPassword := password + s.salt
	return base64.StdEncoding.EncodeToString(s.sha1Hash.Sum([]byte(saltedPassword)))
}

func (s sha1SaltHash) Valid(password, passwordHash string) error {
	if s.hash(password) != passwordHash {
		return fmt.Errorf("password hash and given hash not equal")
	}
	return nil
}

func NewSHA1SaltHash(salt string) PasswordHash {
	return &sha1SaltHash{
		salt:     salt,
		sha1Hash: sha1.New(),
	}
}
