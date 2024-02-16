package models

type User struct {
	Uuid    string
	Name    string
	Email   string
	Deleted bool
}
