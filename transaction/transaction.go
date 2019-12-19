package transaction

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"transaction/model"
)

type Transaction struct {
	DB *gorm.DB
}

type BookData struct {
	Data []Book `json:"data"`
}

type Book struct {
	BookName string `json:"book_name"`
	Author   string `json:"author"`
	Qty      int32  `json:"qty"`
}

func (t Transaction) GetTransactions(c *gin.Context) {
	var transaction []model.Transaction
	var book BookData
	db := t.DB

	db.Find(&transaction)

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	productUrl := os.Getenv("PRODUCT_URL")
	productPort := ":" + os.Getenv("PRODUCT_PORT")
	r := req.New()
	req.Debug = true

	resp, err := r.Get(productUrl + productPort + "/books")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "error get book data " + err.Error(),
		})
	}

	resp.ToJSON(&book)

	c.JSON(200, gin.H{
		"data":    transaction,
		"product": book,
	})
}

func (t Transaction) CreateTransactions(c *gin.Context) {
	var request model.Transaction
	db := t.DB

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	db.Create(&request)

	c.JSON(200, gin.H{
		"message": "create transaction success",
		"data":    request,
	})
}
