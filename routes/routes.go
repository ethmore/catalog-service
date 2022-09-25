package routes

import (
	"catalog-service/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/test", controllers.Test())

	g.POST("/getAllProducts", controllers.GetAllProducts())
	g.POST("/getProduct", controllers.GetProduct())
}
