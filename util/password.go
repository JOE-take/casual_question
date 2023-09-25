package util

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {

	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidPassword(exising string, request string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(exising), []byte(request))
	if err != nil {
		return false, err
	}
	return true, nil
}
