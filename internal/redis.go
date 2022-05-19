package internal

import (
	"context"
	"fmt"
)
import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

var RedisClient *redis.Client

func InitRedis() {
	h := AppConf.RedisConfig.Host
	p := AppConf.RedisConfig.Port
	passwd := AppConf.RedisConfig.Password
	addr := fmt.Sprintf("%s:%d", h, p)
	fmt.Println(addr)
	RedisClient = redis.NewClient(&redis.Options{Addr: addr, Password: passwd})
	ping := RedisClient.Ping(context.Background())
	fmt.Println(ping.String())
	fmt.Println("Redis初始化完成。。。")
}
