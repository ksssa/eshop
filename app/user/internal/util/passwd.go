package util

import (
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func CryptoPasswd(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswd(passwd1, passwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwd1), []byte(passwd2))
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
