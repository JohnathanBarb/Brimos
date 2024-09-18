package models

import "github.com/JohnathanBarb/brimos/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `UPDATE todo SET title=$2, description=$3, done=$4 WHERE id=$1`
	res, err := conn.Exec(stmt, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
