package model

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	BookID     uint32 `json:"book_id" gorm:"column:book_id"`
	BorrowCode string `json:"borrow_code" gorm:"column:borrow_code"`
	Amount     int32  `json:"amount" gorm:"column:amount"`
	Book
}
