package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/trananh-it-hust/ChatApp/pkg/logger"
	"github.com/trananh-it-hust/ChatApp/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Log    *logger.Logger
	MDB    *gorm.DB
	Rdb    *redis.Client
)
