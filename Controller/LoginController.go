package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登录验证
func Login(c *gin.Context) {
	// 这里应该跳转到登录页面然后登录完成后返回openid到GetCode路由
	c.Redirect(http.StatusMovedPermanently, "/GetCode?openid=122")
}
