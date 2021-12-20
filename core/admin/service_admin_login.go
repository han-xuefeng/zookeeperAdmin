package admin

import (
	"github.com/pkg/errors"
	"han-xuefeng/zookeeperAdmin/dto"
	"han-xuefeng/zookeeperAdmin/infra/base"
	"han-xuefeng/zookeeperAdmin/lib"
	"han-xuefeng/zookeeperAdmin/service"
	"sync"
)

type adminLoginService struct {

}
var once sync.Once

func init() {
	once.Do(func() {
		service.IAdminLoginService = new(adminLoginService)
	})
}

func (a *adminLoginService)Login(input *dto.AdminLoginInput) (*Admin, error) {

	admin := &Admin{
		UserName: input.Username,
	}
	adminDao := &AdminDao{
		runner: base.DbxDatabase(),
	}
	err := adminDao.runner.Find(admin).Error
	if err != nil {
		return nil, err
	}

	altPassword := lib.GenSaltPassword(admin.Salt, input.Password)
	if admin.Password != altPassword {
		return nil, errors.New("用户名密码错误")
	}
	return admin,nil
}