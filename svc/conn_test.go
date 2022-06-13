package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestConn(t *testing.T) {
	client := NewRedisClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			t.Error(err)
		}
	}(client)
	ctx := context.Background()
	key := "test_key"
	value := "test_value"

	// Set 命令，第三个参数是超时时间， 0 表示不设置超时时间
	// 通过Err()方法获取错误信息
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		t.Error(err)
	}

	// Get 命令, 通过Result() 方法获取结果及错误信息
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Error(err)
	}
	if result != value {
		t.Errorf("expect %s, got %s", value, result)
	}

	// Del 命令
	err = client.Del(ctx, key).Err()
	if err != nil {
		t.Error(err)
	}

}
