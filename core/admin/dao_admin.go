package admin

import (
	"gorm.io/gorm"
)

type AdminDao struct {
	runner *gorm.DB
}

func (a *AdminDao)FindOne(admin *Admin, conds ...interface{}) error{
	return a.runner.First(admin, "user_name", conds).Error
}