package main

import (
	"catalog/dotEnv"
	"catalog/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{dotEnv.GoDotEnvVariable("BFF_URL")}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PublicRoutes(public)

	if err := router.Run(":3005"); err != nil {
		panic(err)
	}
}
