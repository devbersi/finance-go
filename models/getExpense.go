package models

import "github.com/devbersi/finance-go/db"

func Get(id int64) (expense Expense, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM expense WHERE id= $1`, id)
	
	err = row.Scan(&expense.ID, &expense.Title, &expense.Description, &expense.Value, &expense.PaidOut)

	return
}