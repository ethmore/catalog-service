package routes

import (
	"catalog/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/getAllProducts", controllers.GetAllProducts())
	g.POST("/getProduct", controllers.GetProduct())
}
