package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func VerifyPassword(hashedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	if err != nil {
		return false
	}
	return true
}
