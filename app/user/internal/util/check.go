package util

import (
	"fmt"
	"regexp"
)

func CheckEmailValid(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func CheckMobileValid(mobile string) (bool, error) {
	reg1 := regexp.MustCompile(`1[3-9]`)
	if reg1 == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	return false
}

func CheckPasswdValid(passwd string) bool {

	return false
}
