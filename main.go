package main

import (
	"karate-backend/controllers"
	"karate-backend/initializers"
	"karate-backend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()
	r.POST("/v1/security/signup", controllers.Signup)
	r.POST("/v1/security/login", controllers.Login)
	r.GET("/v1/security/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
