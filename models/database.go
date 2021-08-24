package models

import (
	pop "github.com/gobuffalo/pop/v5"
)

var DB *pop.Connection

func EstablishConnection() {
	db, err := pop.Connect("development") // TODO: environment from... somewhere

	if err != nil {
		panic(err)
	}

	DB = db
}
