package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	encryptedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encryptedBytes), err
}

//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
