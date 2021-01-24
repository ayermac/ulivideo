package redisClient

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 直接连接
func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", "127.0.0.1:6380")
	return pool
}

// 通过连接池
func PoolConnect() redis.Conn  {

	pool := &redis.Pool{
		MaxIdle: 1000,// 最大空闲连接数
		MaxActive: 1000,// 最大连接数
		IdleTimeout: 180 * time.Second,// 空闲连接超时时间
		Wait: false,// 超过最大连接数时，是等待还是报错
		Dial: func() (redis.Conn, error) {
			redisPassword, _ := beego.AppConfig.String("redisPassword")
			redisAddress, _ :=  beego.AppConfig.String("redisAddress")
			setPasswd := redis.DialPassword(redisPassword)
			c, err := redis.Dial("tcp", redisAddress, setPasswd)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return pool.Get()
}
