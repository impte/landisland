package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func StringIsEmpty(text string) bool {
	if text == "" {
		return true
	}
	return false
}

func Md5(plaintext string) (string, error) {
	if StringIsEmpty(plaintext) {
		return "", errors.New("散列字符串不能为空！")
	}
	bytePwd := []byte(plaintext)
	md5Ctx := md5.New()
	md5Ctx.Write(bytePwd)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr), nil
}
