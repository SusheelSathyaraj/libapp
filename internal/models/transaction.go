package models

import "time"

type Transaction struct {
	TransactionID int       `json:"transaction_id"`
	MemberID      int       `json:"member_id"`
	BookID        int       `json:"book_id"`
	IssueDate     time.Time `json:"issue_date"`
	DueDate       time.Time `json:"due_date"`
	ReturnDate    time.Time `json:"return_date"`
	Fine          int       `json:"fine"`
	Status        string    `json:"status"` //"returned", "borrowed", "overdue"
}
