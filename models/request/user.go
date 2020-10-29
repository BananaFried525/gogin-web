package request

type User struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (u *User) IsEmpty() bool {
	if (User{}) == *u {
		return true
	}
	return false
}
