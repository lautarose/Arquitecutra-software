package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	UserName string `gorm:"type:varchar(350);not null;unique"`
}

type Users []User
