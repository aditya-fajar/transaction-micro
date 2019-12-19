package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"transaction/config"
	"transaction/transaction"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	port := os.Getenv("PORT")

	r := gin.Default()
	db := config.DBInit()
	transaction := transaction.Transaction{DB: db}

	config.RegisterConsul()

	r.GET("/books", transaction.GetTransactions)
	r.POST("/books", transaction.CreateTransactions)

	r.GET("/healthcheck", config.Healthcheck)

	r.Run(":" + port)
}
