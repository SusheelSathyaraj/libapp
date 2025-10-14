package models

type User struct {
	MemberID int    `json:"member_id"`
	Name     string `json:"name"`
	PhoneNo  int    `json:"phone_no"`
	Email    string `json:"email"`
}
