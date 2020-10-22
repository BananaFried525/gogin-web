package request

type User struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
}

func (u *User) IsEmpty() bool {
	if (User{}) == *u {
		return true
	}
	return false
}
