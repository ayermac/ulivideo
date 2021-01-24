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
		MaxIdle: 10,// 最大空闲连接数
		MaxActive: 100,// 最大连接数
		IdleTimeout: 180 * time.Second,// 空闲连接超时时间
		Wait: false,// 超过最大连接数时，是等待还是报错
		Dial: func() (redis.Conn, error) {
			redisAddress, _ :=  beego.AppConfig.String("redisAddress")
			c, err := redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			redisPassword, _ := beego.AppConfig.String("redisPassword")
			if _, err := c.Do("AUTH", redisPassword); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	return pool.Get()
}
