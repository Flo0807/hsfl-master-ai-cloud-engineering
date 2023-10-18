package crypto

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (hasher *BcryptHasher) Hash(data []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
}

func (hasher *BcryptHasher) Validate(data []byte, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, data) == nil
}
