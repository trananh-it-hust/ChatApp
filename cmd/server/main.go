package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/trananh-it-hust/ChatApp/cmd/swag/docs"
	"github.com/trananh-it-hust/ChatApp/pkg/initialize"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Swagger documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := initialize.Initialize()
	// Start the server on port 8080

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
