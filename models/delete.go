package models

import "github.com/JohnathanBarb/brimos/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM todo WHERE id=$1`
	res, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
