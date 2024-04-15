package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // Redis 密码，如果没有设置密码则为空字符串
		DB:       0,                // Redis 数据库索引，默认为 0
	})

	// 创建上下文对象
	ctx := context.Background()

	// 测试连接是否成功
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	// 设置键值对
	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取值
	value, err := client.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key:", value)
	}

	// 删除键值对
	result, err := client.Del(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted keys:", result)

	// 关闭连接
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
