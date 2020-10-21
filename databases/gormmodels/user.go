package gormmodels

type User struct {
	Model
	UserName string `gorm:"column:username" json:"userName"`
}
