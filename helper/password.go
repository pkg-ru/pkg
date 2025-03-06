package helper

import (
	"crypto/md5"

	"golang.org/x/crypto/bcrypt"
)

// Хешируем пароль
func PasswordHash[T string | []byte](password T) string {
	pass := []byte(password)
	l := len(pass)
	if l > 72 {
		pass = pass[0:71]
	}
	bytes, err := bcrypt.GenerateFromPassword(pass, 14)
	if err == nil {
		return string(bytes)
	} else {
		return string(md5.New().Sum([]byte(password)))
	}
}

// Сравниваем пароль с хешем
func PasswordCheck[T string | []byte](hash_password, password T) bool {
	pass := []byte(password)
	l := len(pass)
	if l > 72 {
		pass = pass[0:71]
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), pass)
	if err != nil {
		return string(md5.New().Sum([]byte(password))) == string(hash_password)
	}
	return err == nil
}
