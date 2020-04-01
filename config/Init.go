package config

import (
	"os"

	"github.com/gin-contrib/sessions/redis"
	"github.com/joho/godotenv"
)

func Init() {

}

// InitRedis 初始化redis
func InitRedis() (Store, error) {
	// 从本地读取环境变量
	godotenv.Load("/go/src/shuiOauth/.env")
	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	r ,err:= redis.NewStore(
		10, 
		"tcp", 
		os.Getenv("REDIS_HOST"), 
		"",
		[]byte(os.Getenv("SESSION_SECRET")
		))
	return r,err
}
