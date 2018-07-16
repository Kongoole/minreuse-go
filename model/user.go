package model

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

func (u *UserModel) GetPwd(account string) string {
	return ""
}
