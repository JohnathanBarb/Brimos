package models

import "github.com/JohnathanBarb/brimos/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `INSERT INTO todos (title, description, done) values ($1, $2, %3) RETURNING id`
	err = conn.QueryRow(stmt, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
