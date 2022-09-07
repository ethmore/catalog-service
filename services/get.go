package services

import (
	"bytes"
	"catalog/dotEnv"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

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

type Products struct {
	Products []Product
}

func GetAllProducts() ([]Product, error) {
	jsonBody := []byte(`{"client_message": "to server!"}`)
	bodyReader := bytes.NewReader(jsonBody)
	requestURL := dotEnv.GoDotEnvVariable("GET_ALL_PRODUCTS")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}

	//Read response
	b, readErr := io.ReadAll(res.Body)
	if err != nil {
		return nil, readErr
	}
	defer res.Body.Close()

	var products Products
	if unmarshalErr := json.Unmarshal([]byte(b), &products); err != nil {
		return nil, unmarshalErr
	}

	return products.Products, nil
}

func GetProduct(requestBody GetProductResponse) (*Product, error) {
	body, _ := json.Marshal(requestBody)
	bodyReader := bytes.NewReader(body)

	requestURL := dotEnv.GoDotEnvVariable("GET_PRODUCT")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}

	//Read response
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	defer res.Body.Close()

	var product Product
	if unmarshalErr := json.Unmarshal([]byte(b), &product); err != nil {
		return nil, unmarshalErr
	}

	return &product, nil
}
