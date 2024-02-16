package models

import "main.go/db"

func DeleteUser(uuid string) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET deleted= true WHERE uuid=$1`, uuid)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
