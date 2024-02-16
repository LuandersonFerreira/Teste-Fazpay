package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"main.go/configs"
)

func OpenConnection() (*sql.DB, error) {
	configs := configs.GetDB()

	strConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configs.Host, configs.Port, configs.User, configs.Pass, configs.Database)

	conn, err := sql.Open("postgres", strConnection)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
