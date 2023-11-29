# BASIC TRADE
**Golang Class Final Project - Simple API using Gin and GORM**

## Project Description

This project is the result of the final project for the Golang class. A simple API is built using the Gin web framework and GORM ORM to interact with a MySQL database.

## Installation and Usage

1. Make sure you have Golang and MySQL installed on your system.

2. Clone this repository:

   ```bash
   git clone https://github.com/NaofalMufid/basic-trade.git
   cd repo
   go mod tidy
   go run main.go
   ``````

3. Available Endpoints
   Railway url : https://fp-basic-trade.up.railway.app/
    | Method | Endpoint                    | Description                   |
    | ------ | --------------------------- | ----------------------------- |
    | POST   | /api/auth/register          | Register a new user           |
    | POST   | /api/auth/login             | Login with existing credentials |
    | GET    | /api/products/              | Get all products              |
    | POST   | /api/products/              | Create a new product          |
    | GET    | /api/products/:uuid         | Get product by UUID           |
    | PUT    | /api/products/:uuid         | Edit a product by UUID        |
    | DELETE | /api/products/:uuid         | Delete a product by UUID      |
    | GET    | /api/variants/              | Get all variants              |
    | POST   | /api/variants/              | Create a new variant          |
    | GET    | /api/variants/:uuid         | Get variant by UUID           |
    | PUT    | /api/variants/:uuid         | Edit a variant by UUID        |
    | DELETE | /api/variants/:uuid         | Delete a variant by UUID      |

