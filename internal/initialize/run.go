package initialize

func Initialize() {
	// Initialize the configuration
	LoadConfig()

	// Initialize the logger
	InitializeLogger()
	// Initialize the MySQL database connection
	InitializeMySQL()

	// Initialize the Redis connection
	InitializeRedis()

	// Initialize the router
	r := InitializeRouter()

	r.Run(":8080")
}
