package gormmodels

import "github.com/bananafried525/gogin-web/models/request"

type User struct {
	Model
	UserName string `gorm:"column:user_name;unique" json:"userName"`
	Email    string `gorm:"column:user_email" json:"email"`
	Password string `gorm:"column:user_password" json:"password"`
	RoleID   uint
	Role     Role
}

func (u *User) NewGuest(userReq request.User, roleID int) {
	u.UserName = userReq.UserName
	u.Email = userReq.Email
	u.Password = userReq.Password
	if roleID != 0 {
		u.RoleID = uint(roleID)
	} else {
		u.RoleID = 2
	}
}

func NewUser() *User {
	return &User{}
}

type Role struct {
	ID   uint `gorm:"primaryKey"`
	Role string
}
