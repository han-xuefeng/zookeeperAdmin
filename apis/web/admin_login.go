package web

import (
	"github.com/gin-gonic/gin"
	"han-xuefeng/zookeeperAdmin/infra"
	"han-xuefeng/zookeeperAdmin/infra/base"
)

func init()  {
	infra.RegisterApi(new(AdminLoginApi))
}

type AdminLoginApi struct {
}

func (a *AdminLoginApi)Init(){
	adminLogin := &AdminLoginApi{}
	group := base.Gin().Group("/admin_login")
	group.Use()
	{

	}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLoginOut)
}

func (a *AdminLoginApi) AdminLogin(ctx *gin.Context) {
	ctx.String(0,"**********")
}

func (a *AdminLoginApi) AdminLoginOut(ctx *gin.Context) {
	ctx.String(0,"**********")
}
