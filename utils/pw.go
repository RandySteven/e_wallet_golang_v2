package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func IsPasswordValid(hashPass, requestPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(requestPass))
	if err == nil {
		return true
	}
	return false
}
