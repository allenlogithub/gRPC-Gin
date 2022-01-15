package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	bytePwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return "", errors.New("BcryptHashError")
	}

	return string(hash), nil
}

func ComparePassword(hashedPwd, pwd string) bool {
	byteHashedPwd, bytePwd := []byte(hashedPwd), []byte(pwd)
	if err := bcrypt.CompareHashAndPassword(byteHashedPwd, bytePwd); err != nil {
		return false
	}

	return true
}
