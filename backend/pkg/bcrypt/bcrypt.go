package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes), err
}

func IsPasswordEqualToHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
