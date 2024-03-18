package models

import (
	"github.com/devbersi/finance-go/db"
)

func Insert(expense Expense) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO expenses (title, description, value, paidOut) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, expense.Title, expense.Description, expense.Value, expense.PaidOut).Scan(&id)

	return
}