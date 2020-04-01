package router

import (
	controller "shuiOauth/Controller"
	"shuiOauth/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由
func NewRouter() *gin.Engine {
	router := gin.Default()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	store, _ := config.InitRedis()
	router.Use(sessions.Sessions("mysession", store))
	// 测试
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	// session保存
	router.GET("/SetSession", controller.SetSession)
	// 获得code
	router.GET("/GetCode", controller.GetCode)
	// 获得token
	router.POST("/GetToken", controller.GetToken)
	// 获得资源
	router.POST("/GetResource", controller.GetResource)
	// 登录
	router.GET("/login", controller.Login)
	return router
}
