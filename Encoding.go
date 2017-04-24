package codemaru

import (
	"bytes"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

func CP949toUTF8(input []byte) []byte {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewDecoder())
	wr.Write(input)
	wr.Close()
	return bufs.Bytes()
}

func UTF8toCP949(input []byte) []byte {
	var bufs bytes.Buffer
	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	wr.Write(input)
	wr.Close()
	return bufs.Bytes()
}
