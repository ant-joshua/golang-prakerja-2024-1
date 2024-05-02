package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) (bool, error) {

	result := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if result != nil {
		return false, result
	}

	return true, nil
}
