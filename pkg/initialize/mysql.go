package initialize

import (
	"strconv"
	"time"

	"github.com/trananh-it-hust/ChatApp/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() {
	conf := global.Config.Mysql
	dsn := conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" + conf.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Log.Error("Failed to connect to MySQL", zap.Error(err))
		return
	}
	global.MDB = db
	global.Log.Info("Connected to MySQL", zap.String("host", conf.Host), zap.Int("port", conf.Port))
	SetPool()
}

func SetPool() {
	conf := global.Config.Mysql
	sqlDB, err := global.MDB.DB()
	if err != nil {
		global.Log.Error("Failed to get DB", zap.Error(err))
		return
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(conf.MaxIdle) * time.Second)
	sqlDB.SetMaxOpenConns(conf.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.MaxLife) * time.Second)
	global.Log.Info("MySQL connection pool set", zap.Int("maxIdle", conf.MaxIdle), zap.Int("maxOpen", conf.MaxOpen), zap.Int("maxLife", conf.MaxLife))
}
