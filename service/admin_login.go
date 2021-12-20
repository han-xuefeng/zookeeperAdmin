package service

import (
	"han-xuefeng/zookeeperAdmin/core/admin"
	"han-xuefeng/zookeeperAdmin/dto"
)

var IAdminLoginService AdminLoginService

type AdminLoginService interface {
	Login(input *dto.AdminLoginInput) (*admin.Admin, error)
}

func GetAdminLoginService() AdminLoginService {
	return IAdminLoginService
}

