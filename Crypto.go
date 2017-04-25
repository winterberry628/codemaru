package codemaru

import (
	"encoding/base64"
	"strings"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Encrypt(text []byte) []byte {
	a := base64.StdEncoding.EncodeToString(text)
	a = base64.StdEncoding.EncodeToString([]byte(ReverseString(a)))

	a = strings.Replace(a, "a", "^", len(a))
	a = strings.Replace(a, "b", "$", len(a))
	a = strings.Replace(a, "c", "!", len(a))
	a = strings.Replace(a, "d", "%", len(a))

	return []byte(a)
}

func Decrypt(text []byte) []byte {
	textString := string(text)
	textString = strings.Replace(textString, "^", "a", len(textString))
	textString = strings.Replace(textString, "$", "b", len(textString))
	textString = strings.Replace(textString, "!", "c", len(textString))
	textString = strings.Replace(textString, "%", "d", len(textString))

	a, err := base64.StdEncoding.DecodeString(textString)
	if err != nil {
		panic(err.Error())
	}
	a, err = base64.StdEncoding.DecodeString(ReverseString(string(a)))
	if err != nil {
		panic(err.Error())
	}
	return []byte(a)
}
