package service

import (
	"github.com/kongoole/minreuse-go/utils/log"

	"github.com/gorilla/sessions"
	"github.com/kongoole/minreuse-go/model"
	"golang.org/x/crypto/bcrypt"
)

type Login struct{}

var store = sessions.NewCookieStore([]byte("hello"))

// LoginService generates a singleton login service
func LoginService() *Login {
	return &Login{}
}

// CheckLogin checks account and password
// if passed, session will be set
func (l Login) CheckLogin(account string, pwd string) bool {
	// a, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	// fmt.Println(string(a))
	// get pwd by account
	userModel := model.UserModelInstance()
	hashedPwd := userModel.GetPwd(account)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		log.Fatal("fail to compare pwd: " + err.Error())

		return false
	}

	return true
}
