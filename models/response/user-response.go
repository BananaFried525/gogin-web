package response

import "github.com/bananafried525/gogin-web/databases/gormmodels"

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (u *UserResponse) New(user gormmodels.User) {
	u.ID = user.ID
	u.UserName = user.UserName
	u.Email = user.Email
	u.Role = user.Role.Role
}
