package web

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"han-xuefeng/zookeeperAdmin/core/admin"
	"han-xuefeng/zookeeperAdmin/dto"
	"han-xuefeng/zookeeperAdmin/infra"
	"han-xuefeng/zookeeperAdmin/infra/base"
	"han-xuefeng/zookeeperAdmin/lib"
	"han-xuefeng/zookeeperAdmin/middleware"
	"time"
)

func init()  {
	infra.RegisterApi(new(AdminLoginApi))
}

type AdminLoginApi struct {
	service *admin.AdminLoginService
}

func (a *AdminLoginApi)Init(){
	a.service = admin.GetAdminLoginService()
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
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	sess := sessions.Default(ctx)
	sess.Set(lib.AdminSessionInfoKey, string(sessBts))
	//sess.Set("a", "b")
	sess.Save()
	adminLoginOutput := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(ctx,adminLoginOutput)
}

func (a *AdminLoginApi) AdminLoginOut(ctx *gin.Context) {
	ctx.String(0,"**********")
}
