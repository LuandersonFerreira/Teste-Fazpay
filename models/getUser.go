package models

import "main.go/db"

func GetUser(uuid string) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM users WHERE uuid=$1 and deleted = false`, uuid)

	err = row.Scan(&user.Uuid, &user.Name, &user.Email, &user.Deleted)

	return
}
