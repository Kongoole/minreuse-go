package service

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"github.com/kongoole/minreuse-go/model"
)

type Login struct {}

type Pwd struct {}

func (p Pwd) Encode(raw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
}

// CheckLogin checks account and password
// if passed, session will be set
func (l Login) CheckLogin(account string, pwd string) bool {
	// check account & pwd
	hashedPwd, err := Pwd{}.Encode(pwd)
	if err != nil {
		log.Fatal("fail to hash pwd: " + err.Error())
	}
	// get pwd by account
	userModel := model.UserModelInstance()
	if string(hashedPwd) != userModel.GetPwd(account) {
		return false
	}
	return true
}
