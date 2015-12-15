package password

import "golang.org/x/crypto/bcrypt"

// Make hashes input using bcrypt, returns a hex string
func Make(pwd string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
