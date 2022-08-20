package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//convert plain string to hashed value
func HashedPassword(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password: ", err)
		return "", err
	}
	return string(hashed_password), nil
}

func MatchPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
