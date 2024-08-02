package main

import (
	"example/configs"

	"example/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create the LoginController
	// loginController := controllers.NewLoginController()

	// Create the router
	r := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(r)

	// Define the routes
	// r.POST("/login", loginController.LoginAdmin)
	r.Run("localhost:3000")
}
