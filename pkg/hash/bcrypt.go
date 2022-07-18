package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type ConfigBcrypt struct {
	Cost int
}

type Bcrypt struct {
	cost int
}

func NewBcrypt(cfg ConfigBcrypt) *Bcrypt {
	return &Bcrypt{
		cost: cfg.Cost,
	}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	return string(hash), err
}

func (b *Bcrypt) IsPasswordEqualToHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
