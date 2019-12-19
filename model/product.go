package model

type BookData struct {
	Data Book `json:"data"`
}

type Book struct {
	BookName string `json:"book_name"`
	Author   string `json:"author"`
	Qty      int32  `json:"qty"`
}
