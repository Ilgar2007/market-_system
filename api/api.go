package api

import (
	"market/config"
	"market/storage"

	"github.com/gin-gonic/gin"

	_ "market/api/docs"
	"market/api/handler"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger JWT API
// @version 1.0
// @description Create  Go REST API with JWT Authentication in Gin Framework
// @contact.name API Support
// @termsOfService demo.com
// @contact.url http://demo.com/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath
// @Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Branch ...
	r.POST("/branch", handler.CreateBranch)
	r.GET("/branch/:id", handler.BranchGetById)
	r.GET("/branch", handler.GetListBranch)
	r.PUT("/branch/:id", handler.BranchUpdate)
	r.DELETE("/branch/:id", handler.BranchDelete)

	r.POST("/center", handler.CreateCenter)
	r.GET("/center/:id", handler.CenterGetById)
	r.GET("/center", handler.GetListCenter)
	r.PUT("/center/:id", handler.CenterUpdate)
	r.DELETE("/center/:id", handler.CenterDelete)

	r.POST("/provider", handler.CreateProvider)
	r.GET("/provider/:id", handler.ProviderGetById)
	r.GET("/provider", handler.GetListProvider)
	r.PUT("/provider/:id", handler.ProviderUpdate)
	r.DELETE("/provider/:id", handler.ProviderDelete)

	r.POST("/employee", handler.CreateEmployee)
	r.GET("/employee/:id", handler.EmployeeGetById)
	r.GET("/employee", handler.GetListEmployee)
	r.PUT("/employee/:id", handler.EmployeeUpdate)
	r.DELETE("/employee/:id", handler.EmployeeDelete)
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
