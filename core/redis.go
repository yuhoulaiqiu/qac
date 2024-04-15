package core

import (
	"QAComm/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// InitRedis Redis数据库初始化及连接
func InitRedis() *redis.Client {
	if global.Config.Redis.Addr == "" {
		global.Log.Warnln("未配置Redis，取消连接")
		return nil
	}
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,     // Redis 地址
		Password: global.Config.Redis.Password, // Redis 密码
		DB:       global.Config.Redis.Db,       // Redis 数据库索引
	})
	// 创建上下文对象
	ctx := context.Background()
	// 测试连接是否成功
	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.Log.Fatal(fmt.Sprintf("[%s] redis连接失败", global.Config.Redis.Addr))
	}
	return client
}
