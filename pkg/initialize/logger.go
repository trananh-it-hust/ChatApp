package initialize

import (
	"github.com/trananh-it-hust/ChatApp/global"
	"github.com/trananh-it-hust/ChatApp/pkg/logger"
)

func InitializeLogger() {
	global.Log = logger.NewLogger(global.Config.Logger)
}
