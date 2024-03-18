package models

type Expense struct {
	ID int64 `json:"id"`
	Title string `json: "title"`
	Description string `json: "description"`
	Value int `json: "value"`
	PaidOut bool `json: "paidOut"`
}