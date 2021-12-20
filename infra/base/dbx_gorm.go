package base

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"han-xuefeng/zookeeperAdmin/infra"
)

//dbx 数据库实例
var database *gorm.DB

func DbxDatabase() *gorm.DB {
	Check(database)
	return database
}

type DbxGorm struct {
	infra.BaseStarter
}

func (db *DbxGorm)init(ctx infra.StarterContext)  {

	user,_ := ctx.Props().Get("mysql.user")
	password,_ := ctx.Props().Get("mysql.password")
	host,_ := ctx.Props().Get("mysql.host")
	dataName,_ := ctx.Props().Get("mysql.database")
	dsn :=  user + ":" + password + "@tcp(" + host + ")/" + dataName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	database = dbGorm
}
