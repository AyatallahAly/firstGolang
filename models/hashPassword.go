package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repo) BeforeSave(user User) (pass string) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("unable to hash password ", err)
	}
	pass = string(hashedPassword)
	return pass
}
