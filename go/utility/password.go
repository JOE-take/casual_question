package utility

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {

	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
