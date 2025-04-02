package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"main.go/pkg/logger"
	"main.go/pkg/setting"
)

var (
	Config setting.Config
	Log    *logger.Logger
	MDB    *gorm.DB
	Rdb    *redis.Client
)
