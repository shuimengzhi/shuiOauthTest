package config

import (
	"os"
	"time"

	"github.com/gin-contrib/sessions/redis"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DB 数据库链接单例
var DB *gorm.DB

// InitMysql 连接mysql
func InitMysql() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_HOST"))
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	DB = db
}

// CloseMysql 断开mysql连接
func CloseMysql() {
	DB.Close()
}

// InitRedis 初始化redis
func InitRedis() (redis.Store, error) {
	// 从本地读取环境变量
	godotenv.Load("/go/src/shuiOauth/.env")
	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	r, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_HOST"), "", []byte(os.Getenv("SESSION_SECRET")))
	return r, err
}
