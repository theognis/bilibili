package dao

import (
	"bilibili/tool"
	"github.com/gin-gonic/gin"
	"time"
)

type RedisDao struct {

}

func (r *RedisDao) RedisGetValue(ctx *gin.Context, key string) (string, error)  {
	redisConn := tool.GetRedisConn()
	cmd := redisConn.Get(ctx, key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}

	return cmd.Val(), nil
}

func (r *RedisDao) RedisSetValue(ctx *gin.Context, key string, value string) error {
	redisConn := tool.GetRedisConn()
	cmd := redisConn.Set(ctx, key, value, time.Minute * 10)
	return cmd.Err()
}