package controller

import (
	"os"
"shuiOauth/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liudng/godump"
)

// Gcr 转换
// type Gcr

// GetCode 获得code
func GetCode(c *gin.Context) {

	session := sessions.Default(c)
	// 接受请求的信息
	gcr := GetCodeRequest{
		ClientID:     session.Get("client_id").(string),
		RedirectURL:  session.Get("redirect_uri").(string),
		ResponseType: session.Get("response_type").(string),
		State:        session.Get("state").(string),
	}
	gcr.service.CodeVerificationService()
	godump.Dump(gcr)
	os.Exit(1)
	// 验证请求的信息

	os.Exit(1)
}

// GetToken 获得token
func GetToken(c *gin.Context) {

}

// GetResource 获得资源
func GetResource(c *gin.Context) {

}
