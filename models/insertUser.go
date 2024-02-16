package models

import (
	"main.go/db"
)

func InsertUser(user User) (uuid string, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO users (uuid, name, email, deleted) VALUES ($1, $2, $3, $4) RETURNING uuid`

	err = conn.QueryRow(sql, user.Uuid, user.Name, user.Email, user.Deleted).Scan(&uuid)

	return
}
