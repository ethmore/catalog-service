Catalog Service
===

This GO project serves as a microservice for [eCommerce](https://github.com/users/ethmore/projects/4) project.


## Service tasks:

- Retrive product info from [auth-and-db-service](https://github.com/ethmore/auth-and-db-service)



# Installation

Ensure GO is installed on your system
```
go mod download
````

```
go run .
```

## Test
```
curl http://localhost:3005/test
```
### It should return:
```
StatusCode        : 200
StatusDescription : OK
Content           : {"message":"OK"}
```

## Example .env file
This file should be placed inside `dotEnv` folder
```
# Cors URLs
BFF_URL = http://localhost:3001

# Request URLs
GET_ALL_PRODUCTS = http://localhost:3002/getAllProducts
GET_PRODUCT = http://localhost:3002/getProduct
```