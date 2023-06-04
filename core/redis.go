package core

import (
	"context"
	"github.com/go-redis/redis"
	"survey_backend/global"
	"time"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.AddrPort(),
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.Db,
		PoolSize: global.Config.Redis.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Errorf("redis连接失败: %s", err)
		return nil
	}
	return rdb
}
