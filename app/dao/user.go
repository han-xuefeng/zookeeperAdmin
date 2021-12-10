package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int
	Name string
}
