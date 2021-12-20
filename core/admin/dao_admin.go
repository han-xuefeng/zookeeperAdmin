package admin

import (
	"gorm.io/gorm"
)

type AdminDao struct {
	runner *gorm.DB
}

func (dao *AdminDao)getOneById(id int) (*Admin, error){
	admin := &Admin{Id: id}
	err := dao.runner.First(admin).Error
	if err != nil {
		return nil,err
	}
	return admin,nil
}