package tool

import (
	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client

func init() {
	redisCfg := GetCfg().Redis
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})
}

func GetRedisConn() *redis.Client {
	return RedisConn
}
