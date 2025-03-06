package helper

import (
	"regexp"
	"strings"
)

func PhoneRaw(phone string) string {
	phone = strings.Trim(phone, "\r\n\t ")
	reg, _ := regexp.Compile(`[^\+\d]+`)
	phone = reg.ReplaceAllString(phone, "")
	if len([]rune(phone)) >= 10 {
		if phone[0:1] == "+" {
		} else if len([]rune(phone)) == 10 {
			phone = "+7" + phone
		} else if phone[0:1] == "7" {
			phone = "+" + phone
		} else if phone[0:1] == "8" {
			phone = "+7" + phone[1:]
		}
		return phone
	}
	return ""
}

func PhoneFormat(phone string) string {
	reg, _ := regexp.Compile(`^(\+?)(\d{1,3})(\d{3})(\d{3})(\d{2})(\d{2})$`)
	return reg.ReplaceAllString(PhoneRaw(phone), "$1$2($3)$4-$5-$6")
}
