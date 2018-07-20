package service

import "golang.org/x/crypto/bcrypt"

type Pwd struct{}

func (p Pwd) Encode(raw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
}
