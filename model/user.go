package model

import (
	"log"
)

type UserModel struct {
	Model
}

var userModel *UserModel

// ArticleModelInstance creates an ArticleModel instance
func UserModelInstance() *UserModel {
	once.Do(func() {
		userModel = &UserModel{}
	})
	return userModel
}

// GetPwd gets password by account
func (u *UserModel) GetPwd(account string) string {
	u.InitSlave()
	stmt, err := u.Slave.Prepare("SELECT password FROM user WHERE email=? OR name=?")
	if err != nil {
		log.Fatal("fail to get user pwd: " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(account, account)
	if err != nil {
		log.Fatal("fail to get user pwd: " + err.Error())
	}
	defer rows.Close()

	var pwd string
	for rows.Next() {
		rows.Scan(&pwd)
		if len(pwd) > 0 {
			break
		}
	}

	return pwd
}
