package models

import "github.com/devbersi/finance-go/db"

func GetAll() (expenses []Expense, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM expense`)
	if err != nil {
		return 
	}
	
	for rows.Next() {
		var expense Expense 
		err = rows.Scan(&expense.ID, &expense.Title, &expense.Description, &expense.Value, &expense.PaidOut)
		if err != nil {
			continue 
		}

		expenses = append(expenses, expense)
	}

	return
}