package controllers

import (
	"catalog/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := services.GetAllProducts()
		if err != nil {
			fmt.Println("GetAllProducts", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		ctx.JSON(200, gin.H{"message": "OK", "products": products})
	}
}

func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Read client's request
		var requestBody services.GetProductResponse
		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			fmt.Println("body: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		//Create new request to :3002
		product, getErr := services.GetProduct(requestBody)
		if getErr != nil {
			fmt.Println("GetProduct", getErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(200, gin.H{"message": "OK", "productInfo": product})
	}
}
