package initialize

import "main.go/global"

func Initialize() {
	// Initialize the configuration
	LoadConfig()

	// Initialize the logger
	InitializeLogger()

	global.Log.Info("Logger initialized")

	// Initialize the MySQL database connection
	InitializeMySQL()

	// Initialize the Redis connection
	InitializeRedis()

	// Initialize the router
	r := InitializeRouter()

	r.Run(":8080")
}
