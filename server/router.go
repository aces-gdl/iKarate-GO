package server

import (
	"iKarate-GO/controllers"
	"iKarate-GO/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/v1/security/signup", controllers.Signup)
	router.POST("/v1/security/login", controllers.Login)
	router.GET("/v1/security/validate", middleware.RequireAuth, controllers.Validate)

	router.GET("/v1/catalogs/user", middleware.RequireAuth, controllers.GetUsers)

	router.POST("/v1/catalogs/club", middleware.RequireAuth, controllers.PostClub)
	router.GET("/v1/catalogs/clubs", middleware.RequireAuth, controllers.GetClubs)

	router.POST("/v1/catalogs/court", middleware.RequireAuth, controllers.PostCourts)
	router.GET("/v1/catalogs/court", middleware.RequireAuth, controllers.GetCourts)
	router.GET("/v1/catalogs/court/byclub", middleware.RequireAuth, controllers.GetCourtsByClub)

	router.POST("/v1/catalgs/simulateenrollment", middleware.RequireAuth, controllers.PostSimulateEnrollment)
	router.POST("/v1/catalgs/creategroups", middleware.RequireAuth, controllers.PostCreateGroups)
	router.GET("/v1/catalogs/getteams", middleware.RequireAuth, controllers.GetEnrolledTeams)
	router.GET("/v1/catalogs/getteamsbygroup", middleware.RequireAuth, controllers.GetGroups)

	return router
}
