package initialize

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trananh-it-hust/ChatApp/global"
)

func Initialize() *gin.Engine {
	// Initialize the configuration
	LoadConfig()

	// Initialize the logger
	InitializeLogger()
	global.Log.Info(fmt.Sprintf("%s: Run server on port %s", time.Now().Format("2006-01-02 15:04:05"), global.Config.Server.Port))

	// Initialize the MySQL database connection
	InitializeMySQL()

	// Initialize the Redis connection
	InitializeRedis()

	// Initialize the router
	r := InitializeRouter()

	return r
}
