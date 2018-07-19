package service

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/kongoole/minreuse-go/model"
	"golang.org/x/crypto/bcrypt"
)

type Login struct{}

type Pwd struct{}

var store = sessions.NewCookieStore([]byte("hello"))

func (p Pwd) Encode(raw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
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
		log.Println("fail to compare pwd: " + err.Error())
		return false
	}

	return true
}
