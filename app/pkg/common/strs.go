package common

import "unicode"

func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsAnyBlank(strList ...string) bool {
	for _, str := range strList {
		if IsBlank(str) {
			return true
		}
	}
	return false
}
