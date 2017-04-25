package codemaru

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func FileMD5(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		panic(err.Error())
	}

	hashInBytes := hash.Sum(nil)[:16]

	return hex.EncodeToString(hashInBytes)
}
