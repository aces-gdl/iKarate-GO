package main

import (
	"iKarate-GO/controllers"
	"iKarate-GO/initializers"
	"iKarate-GO/middleware"

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

	r.GET("/v1/catalogs/user", middleware.RequireAuth, controllers.GetUsers)

	r.POST("/v1/catalogs/club", middleware.RequireAuth, controllers.PostClub)
	r.GET("/v1/catalogs/club", middleware.RequireAuth, controllers.GetClubs)

	r.POST("/v1/catalogs/court", middleware.RequireAuth, controllers.PostCourts)
	r.GET("/v1/catalogs/court", middleware.RequireAuth, controllers.GetCourts)
	r.GET("/v1/catalogs/court/byclub", middleware.RequireAuth, controllers.GetCourtsByClub)

	r.Run()
}
