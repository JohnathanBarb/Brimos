package db

import (
	"database/sql"
	"fmt"
	"github.com/JohnathanBarb/brimos/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDBConfig()

	sc := fmt.Sprintf("host=%s port=%s user=%s pass=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}