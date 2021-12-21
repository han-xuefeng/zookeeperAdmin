package dto

import "time"

type AdminLoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginOutput struct {
	Token string `form:"token" json:"token" comment:"token"`
}