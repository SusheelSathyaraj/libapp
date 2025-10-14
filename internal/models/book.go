package models

type Book struct {
	BookID    int    `json:"book_id"`
	BookName  string `json:"book_name"`
	Author    string `json:"author"`
	ISBN      int    `json:"isbn"`
	Genre     string `json:"genre"`
	Available bool   `json:"available"`
}
