package Common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"log"
)

var RedisClient *redis.Pool

type RedisCfg struct {
	Host, Port string
	Size       int64
}

func init() {
	var cfg RedisCfg
	if _, err := toml.DecodeFile("./Config/redis.toml", &cfg); err != nil {
		log.Fatal(err)
	}
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 10,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", cfg.Host+":"+cfg.Port)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

func Interface2String(inter interface{}) {
	switch inter.(type) {
	case string:
		fmt.Println("string", inter.(string))
		break
	case int:
		fmt.Println("int", inter.(int))
		break
	case float64:
		fmt.Println("float64", inter.(float64))
		break
	default:
		fmt.Println("都不是")
	}

}

func GetRedisConn() redis.Conn {
	return RedisClient.Get()
}
