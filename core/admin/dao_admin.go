package admin

import (
	"gorm.io/gorm"
	"han-xuefeng/zookeeperAdmin/infra/base"
)

var adminDao *AdminDao

func GetAdminDao() *AdminDao {
	if adminDao == nil {
		adminDao = &AdminDao{
			runner: base.DbxDatabase(),
		}
	}
	return adminDao
}

type AdminDao struct {
	runner *gorm.DB
}

func (a *AdminDao)FindOne(admin *Admin, conds ...interface{}) error{
	return a.runner.First(admin, "user_name", conds).Error
}

func (a *AdminDao)FindOneById(admin *Admin) error{
	return a.runner.First(admin).Error
}

func (a *AdminDao)UpdateOne(admin *Admin) error {
	return a.runner.Updates(admin).Error
}