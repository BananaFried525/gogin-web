package request

type User struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	a        string `json:"a" form:"a" binding:"required"`
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}

func (u *User) IsEmpty() bool {
	if (User{}) == *u {
		return true
	}
	return false
}
