package auth

import "golang.org/x/crypto/bcrypt"

func Validate(hash, origin []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, origin)
	if err != nil {
		println("Invalid reason: ", err.Error())
		return false
	}
	return true
}
