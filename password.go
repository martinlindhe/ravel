package ravel

import "golang.org/x/crypto/bcrypt"

// HashPassword uses bcrypt, returns a 60 character hex string
func HashPassword(pwd string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
