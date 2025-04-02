package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"main.go/global"
)

var ctx = context.Background()

func InitializeRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port), // 55000
		Password: r.Password,                           // no password set
		DB:       r.Database,                           // use default DB
		PoolSize: 10,                                   //
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Log.Error("Redis initialization Error:", zap.Error(err))
	}

	global.Log.Info("Initializing Redis Successfully")
	global.Rdb = rdb
}
