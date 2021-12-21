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
)

// admin 模块控制器
type AdminApi struct {
}

func init() {
	infra.RegisterApi(&AdminApi{})
}

func (a *AdminApi) Init(){
	admin := &AdminApi{}
	group := base.Gin().Group("/admin")
	group.GET("/admin_info", admin.AdminInfo)
	group.POST("/change_pwd", admin.ChangePwd)
}

func (a *AdminApi)AdminInfo(c *gin.Context){
	sess := sessions.Default(c)
	sessInfo := sess.Get(lib.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfo.(string)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

func (a *AdminApi)ChangePwd(c *gin.Context){
	params := &dto.ChangePwdInput{}
	err := c.ShouldBind(params)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//获取session
	sess := sessions.Default(c)
	sessInfo := sess.Get(lib.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfo.(string)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	adminService := &admin.AdminService{}
	admin,err  := adminService.GetOneById(adminSessionInfo.ID)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	//用户存在更改密码
	saltPassword := lib.GenSaltPassword(admin.Salt, params.Password)
	admin.Password = saltPassword
	err = adminService.UpdateOne(admin)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

