package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Response struct {
	Products []Product
}

type ProductResponse struct {
	Products Product
}

type Product struct {
	Id          string
	Title       string
	Price       string
	Description string
	Image       string
	Stock       string
}

type GetProductResponse struct {
	Token string
	Id    string
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Create new request to :3002
		jsonBody := []byte(`{"client_message": "to server!"}`)
		bodyReader := bytes.NewReader(jsonBody)
		requestURL := "http://localhost:3002/getAllProducts"

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		//Read response
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		var resp Response
		json.Unmarshal([]byte(b), &resp)
		ctx.JSON(200, gin.H{"message": "OK", "products": resp})
	}
}

func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Read client's request
		var requestBody GetProductResponse
		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			fmt.Println("body: ", err)
			return
		}

		//Create new request to :3002
		body, _ := json.Marshal(requestBody)
		bodyReader := bytes.NewReader(body)

		requestURL := "http://localhost:3002/getProduct"

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		//Read response
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		var resp ProductResponse
		json.Unmarshal([]byte(b), &resp)
		ctx.JSON(200, gin.H{"message": "OK", "productInfo": resp})
	}
}
