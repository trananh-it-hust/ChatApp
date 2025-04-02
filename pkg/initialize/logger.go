package initialize

import (
	"main.go/global"
	"main.go/pkg/logger"
)

func InitializeLogger() {
	global.Log = logger.NewLogger(global.Config.Logger)
}
