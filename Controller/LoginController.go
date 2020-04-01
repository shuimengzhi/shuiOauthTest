package controller

import "github.com/gin-gonic/gin"

// Login 登录验证
func Login(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is login"})
}
