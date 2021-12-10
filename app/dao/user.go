package dao

type User struct {
	Id int `gorm:"primaryKey"`
	Name string
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "user"
}