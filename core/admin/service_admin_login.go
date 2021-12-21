package admin

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"han-xuefeng/zookeeperAdmin/dto"
	"han-xuefeng/zookeeperAdmin/infra/base"
	"han-xuefeng/zookeeperAdmin/lib"
)

//AdminLoginService
type AdminLoginService struct {

}

func GetAdminLoginService() *AdminLoginService {
	return new(AdminLoginService)
}

func (a *AdminLoginService)Login(input *dto.AdminLoginInput) (*Admin, error) {


	adminDao := &AdminDao{
		runner: base.DbxDatabase(),
	}
	admin := &Admin{}
	err := adminDao.FindOne(admin, "user_name", input.Username)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("用户名密码错误")
	}
	if err != nil {
		return nil, err
	}

	altPassword := lib.GenSaltPassword(admin.Salt, input.Password)
	if admin.Password != altPassword {
		return nil, errors.New("用户名密码错误")
	}
	return admin,nil
}

