package hasher

import (
	"crypto/sha1"
	"fmt"
)

//go:generate mockgen -source=hasher.go -destination=./mocks/hasher_mocks.go -package=mocks
type PasswordHash interface {
	Hash(password string) string
}

type SHA1Hash struct {
	salt string
}

func NewSHA1Hash(salt string) *SHA1Hash {
	return &SHA1Hash{salt: salt}
}

func (h *SHA1Hash) Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
