package server

import (
	"iKarate-GO/controllers"
	"iKarate-GO/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	router.POST("/v1/security/signup", controllers.Signup)
	router.POST("/v1/security/login", controllers.Login)
	router.GET("/v1/security/validate", middleware.RequireAuth, controllers.Validate)

	router.GET("/v1/catalogs/users", middleware.RequireAuth, controllers.GetUsers)
	router.POST("/v1/catalogs/users", middleware.RequireAuth, controllers.PostUser)
	router.PUT("/v1/catalogs/users", middleware.RequireAuth, controllers.PutUser)

	router.GET("/v1/catalogs/permissions", middleware.RequireAuth, controllers.GetPermissions)
	router.POST("/v1/catalogs/permissions", middleware.RequireAuth, controllers.PostPermissions)

	router.GET("/v1/catalogs/dojos", middleware.RequireAuth, controllers.GetDojos)
	router.POST("/v1/catalogs/dojos", middleware.RequireAuth, controllers.PostDojos)

	router.POST("/v1/catalogs/category", middleware.RequireAuth, controllers.PostCategory)
	router.GET("/v1/catalogs/category", middleware.RequireAuth, controllers.GetCatgories)

	router.POST("/v1/catalogs/schedule", middleware.RequireAuth, controllers.PostSchedule)
	router.GET("/v1/catalogs/schedule", middleware.RequireAuth, controllers.GetSchedules)

	router.POST("/v1/catalogs/product", middleware.RequireAuth, controllers.PostProduct)
	router.GET("/v1/catalogs/product", middleware.RequireAuth, controllers.GetProducts)

	router.POST("/v1/utility/loadusers", middleware.RequireAuth, controllers.PostLoadUsers)

	router.POST("/v1/utility/imageupload", middleware.RequireAuth, controllers.UploadImage)
	router.GET("/v1/utility/image", middleware.RequireAuth, controllers.GetImage)

	return router
}
