package models

type Zone struct {
	ID 		int64	`db:"id"`
	Name	string	`db:"name"`
	Description string `db:"name"`
}
