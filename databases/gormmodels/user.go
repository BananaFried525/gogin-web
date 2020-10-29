package gormmodels

type User struct {
	Model
	UserName string `gorm:"column:user_name;unique" json:"userName"`
	Email    string `gorm:"column:user_email" json:"email"`
	Password string `gorm:"column:user_password" json:"password"`
	RoleID   uint
	Role     Role
}

type Role struct {
	ID   uint `gorm:"primaryKey"`
	Role string
}
