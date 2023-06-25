package main

import (
	"log"

	"github.com/Zaidannzzz/test/backend/config"
	"github.com/Zaidannzzz/test/backend/httpserver/controllers"
	"github.com/Zaidannzzz/test/backend/httpserver/repositories"
	"github.com/Zaidannzzz/test/backend/httpserver/routers"
	"github.com/Zaidannzzz/test/backend/httpserver/services"
	"github.com/Zaidannzzz/test/backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:3030
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()
	appRoute := app.Group("/api")
	db, err := config.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	routers.UserRouter(appRoute, userController, authService)

	err = app.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
