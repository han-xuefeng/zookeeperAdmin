package web

import (
	"github.com/gin-gonic/gin"
	"han-xuefeng/zookeeperAdmin/dto"
	"han-xuefeng/zookeeperAdmin/infra"
	"han-xuefeng/zookeeperAdmin/infra/base"
	"han-xuefeng/zookeeperAdmin/middleware"
	"han-xuefeng/zookeeperAdmin/service"
)

func init()  {
	infra.RegisterApi(new(AdminLoginApi))
}

type AdminLoginApi struct {
	service service.AdminLoginService
}

func (a *AdminLoginApi)Init(){
	a.service = service.GetAdminLoginService()
	adminLogin := &AdminLoginApi{}
	group := base.Gin().Group("/admin_login")
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLoginOut)
}

func (a *AdminLoginApi) AdminLogin(ctx *gin.Context) {
	input := &dto.AdminLoginInput{}
	err := ctx.ShouldBind(input)
	if err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	admin,err := a.service.Login(input)
	if err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	middleware.ResponseSuccess(ctx,admin )
}

func (a *AdminLoginApi) AdminLoginOut(ctx *gin.Context) {
	ctx.String(0,"**********")
}
