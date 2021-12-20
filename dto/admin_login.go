package dto

type AdminLoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}