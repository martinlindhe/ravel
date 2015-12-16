package ravel

import "golang.org/x/crypto/bcrypt"

// HashPassword used bcrypt, returns a hex string
func HashPassword(pwd string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
