package models

import (
	"log"
)

func (r *Repo) LoginUser(user User) {

	q := `select * from person where email = $1`
	err := Db.QueryRow(q, user.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err != nil {
		log.Fatal("unable to execute search query", err)

	}

}
