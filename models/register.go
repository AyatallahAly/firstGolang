package models

import (
	"log"
)

func (r *Repo) RegisterUser(user User) {
	q := `
	INSERT INTO person(email,password, role) 
	VALUES($1, $2,$3)
	`
	_, err := Db.Exec(q,
		//user.Username,
		user.Email,
		user.Password,
		user.Role,
	)
	if err != nil {
		log.Fatal(err)
	}
}