package main

import (
	"catalog/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PublicRoutes(public)

	router.Run(":3005")
}
