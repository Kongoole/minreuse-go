package service

import (
	"fmt"

	"github.com/kongoole/minreuse-go/model"
	"golang.org/x/crypto/bcrypt"
)

type Login struct{}

var loginService *Login

// LoginService generates a singleton login service
func LoginService() *Login {
	once.Do(func() {
		loginService = &Login{}
	})
	return loginService
}

// CheckLogin checks account and password
// if passed, session will be set
func (l Login) CheckLogin(account string, pwd string) bool {
	// get pwd by account
	userModel := model.UserModelInstance()
	accountPwd := userModel.GetPwd(account)
	a, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	fmt.Println(string(a), accountPwd)
	err := bcrypt.CompareHashAndPassword([]byte(accountPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
