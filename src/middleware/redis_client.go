package middleware

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"lingye-gin/src/config"
	"log"
	"time"
)

var RedisPoolFactory *redis.Pool

type RedisPool struct{}

// 格式化成字符串
func (RedisPool) String(a interface{}, err error) (string, error) {
	return redis.String(a, err)
}

func (RedisPool) Init() {
	config.Logger.Info("RedisPool Init")

	RedisPoolFactory = &redis.Pool{
		MaxIdle:   config.AppProps.Redis.MaxIdle,
		MaxActive: config.AppProps.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", config.AppProps.Redis.Addr,
				redis.DialReadTimeout(time.Second*10),
				redis.DialConnectTimeout(time.Second*30),
				redis.DialPassword(config.AppProps.Redis.Passwd),
				redis.DialDatabase(config.AppProps.Redis.Database),
			)
			if err != nil {
				log.Println("ERROR: fail init redis pool:", err.Error())
				if conn != nil {
					_ = conn.Close()
				}
				panic(fmt.Sprintf("ERROR: fail init redis pool: %s", err.Error()))
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			pong, err := c.Do("PING")
			// Output: PONG <nil>
			config.Logger.Info("RedisPool Ping And ", pong, err)
			return err
		},
	}
	config.Logger.Info("RedisPool Ok")
	return
}

func (RedisPool) Set() {
	config.Logger.Println("ExampleClient_String")
	// 先从 pool 取出一个连接
	conn := RedisPoolFactory.Get()
	defer conn.Close()

	if _, err := conn.Do("Set", "key", "value", 1*time.Second); err != nil {
		config.Logger.Println(err)
		panic(err)
	}
}
