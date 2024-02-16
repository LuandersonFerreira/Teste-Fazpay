package models

import "main.go/db"

func UpdateUser(user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET name=$2, email=$3 WHERE uuid=$1`, user.Uuid, user.Name, user.Email)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
