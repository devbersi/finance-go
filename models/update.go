package models

import "github.com/devbersi/finance-go/db"


func Update(id int64, expense Expense) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE expenses SET title=$2, description=$3, value=$4, paidOut=$5 WHERE id=$1`, expense.ID, expense.Title, expense.Description, expense.Value, expense.PaidOut)
	if err != nil {
		return 0, err
	}
	
	return res.RowsAffected()
}