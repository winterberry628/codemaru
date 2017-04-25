package codemaru

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
)

func gzipStringPack(text []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(text)
	w.Close()
	return b.Bytes()
}

func gzipStringUnpack(text []byte) []byte {
	var b bytes.Buffer
	r, err := gzip.NewReader(&b)
	if err != nil {
		panic(err.Error())
	}
	r.Read(text)
	r.Close()
	return b.Bytes()
}

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

	return gzipStringPack([]byte(a))
}

func Decrypt(text []byte) []byte {
	text = gzipStringUnpack(text)

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
