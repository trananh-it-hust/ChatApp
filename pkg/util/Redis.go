package util

import (
	"context"
	"fmt"
	"time"

	"main.go/global"
)

func SetValue(key string, value interface{}) error {
	ctx := context.Background()
	_, err := global.Rdb.Set(ctx, key, value, 0).Result()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %v", err)
	}
	return nil
}

func GetValue(key string) (string, error) {
	ctx := context.Background()
	value, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get value from Redis: %v", err)
	}
	return value, nil
}

func DeleteValue(key string) error {
	ctx := context.Background()
	_, err := global.Rdb.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to delete value from Redis: %v", err)
	}
	return nil
}

func SetValueWithExpiration(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	_, err := global.Rdb.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %v", err)
	}
	return nil
}

func KeyExists(key string) (bool, error) {
	ctx := context.Background()
	exists, err := global.Rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if key exists in Redis: %v", err)
	}
	return exists > 0, nil
}
