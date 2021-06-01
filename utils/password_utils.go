package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Compare(userpass, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(userpass), []byte(password)); err != nil {
		return err
	}
	return nil
}
