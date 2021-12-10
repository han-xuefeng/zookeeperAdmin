package starter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var dbg 全局唯一变量
var dbGorm *gorm.DB

func NewDbGorm() *gorm.DB {
	return dbGorm
}

type GormStarter struct {
	BaseStarter
}

func (s *GormStarter) Init(ctx StarterContext)  {
	initDbGorm(ctx)
}

// 初始化gorm
func initDbGorm(ctx StarterContext) {
	dsn := "root:root@tcp(127.0.0.1:3306)/zookeeper_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbGorm = db
}
