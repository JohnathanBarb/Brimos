package models

import "github.com/JohnathanBarb/brimos/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM todo WHERE id = $1`
	row := conn.QueryRow(stmt, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
