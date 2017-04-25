package codemaru

import (
	"encoding/base64"
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

	return []byte(a)
}

func Decrypt(text []byte) []byte {
	a, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		panic(err.Error())
	}
	a, err = base64.StdEncoding.DecodeString(ReverseString(string(a)))
	if err != nil {
		panic(err.Error())
	}
	return []byte(a)
}
